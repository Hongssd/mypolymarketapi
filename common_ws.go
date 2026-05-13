package mypolymarketapi

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
)

const (
	PM_CLOB_API_WS_HOST   = "ws-subscriptions-clob.polymarket.com"
	PM_CLOB_API_WS_MARKET = "/ws/market"
	PM_CLOB_API_WS_USER   = "/ws/user"
	PM_SPORTS_API_WS_HOST = "sports-api.polymarket.com"
	PM_SPORTS_API_WS_PATH = "/ws"
)

const (
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

var (
	WebsocketTimeout   = 10 * time.Second
	WebsocketKeepalive = true
)

type WsChannelType int

const (
	WS_MARKET WsChannelType = iota
	WS_USER
	WS_SPORTS
)

func (t WsChannelType) String() string {
	switch t {
	case WS_MARKET:
		return "market"
	case WS_USER:
		return "user"
	case WS_SPORTS:
		return "sports"
	default:
		return ""
	}
}

type WsAuth struct {
	ApiKey     string `json:"apiKey"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

// WsSubscribeReq 统一 WS 订阅/动态订阅请求体：
// - market 首次订阅：type=market + assets_ids + 可选 initial_dump/level/custom_feature_enabled
// - market 动态订阅：assets_ids + operation + 可选 level/custom_feature_enabled（无需 type）
// - user 首次订阅：type=user + auth；markets 可选（省略则全市场，见官方 WSS user 文档）
// - user 动态订阅：markets + operation（无需 type/auth）
type WsSubscribeReq struct {
	// common
	Type      string `json:"type,omitempty"`      // market / user
	Operation string `json:"operation,omitempty"` // subscribe / unsubscribe（动态订阅）

	// market
	AssetsIDs            []string `json:"assets_ids,omitempty"`
	InitialDump          bool     `json:"initial_dump,omitempty"`
	Level                int      `json:"level,omitempty"`
	CustomFeatureEnabled bool     `json:"custom_feature_enabled,omitempty"`

	// user
	Auth    *WsAuth   `json:"auth,omitempty"`
	Markets *[]string `json:"markets,omitempty"`
}

type WsStreamClient struct {
	isInit  bool
	client  *Client
	channel WsChannelType
	conn    *websocket.Conn

	commonSubMap MySyncMap[string, *Subscription]

	reSubscribeMu      *sync.Mutex
	AutoReConnectTimes int // 自动重连次数

	resultChan      chan []byte
	errChan         chan error
	writeMu         *sync.Mutex
	isClose         bool
	waitSubResult   bool
	waitSubResultMu *sync.Mutex

	// dataHandler 由子类型（如 MarketWsStreamClient）注入，每条非 PONG 原始消息都会调用它。
	// 若为 nil，则消息被静默丢弃。
	dataHandler func(data []byte)
	// reconnectHook 在每次自动重连成功后由 handleResult goroutine 调用。
	// 若为 nil，则走 reSubscribeForReconnect 的遗留逻辑（User/Sports 兼容）。
	reconnectHook func()
}

// Subscription 表示通用 WS 订阅对象（由具体频道自行分发消息到 resultChan）。
type Subscription struct {
	SubId int64
	Ws    *WsStreamClient // 当前订阅绑定的 websocket 客户端
	Args  WsSubscribeReq

	resultChan chan any
	errChan    chan error
	closeChan  chan struct{}
}

// ResultChan 返回业务消息通道（由内部分发协程写入）。
func (sub *Subscription) ResultChan() chan any {
	return sub.resultChan
}

// ErrChan 返回订阅错误通道。
func (sub *Subscription) ErrChan() chan error {
	return sub.errChan
}

// CloseChan 返回订阅关闭信号通道。
func (sub *Subscription) CloseChan() chan struct{} {
	return sub.closeChan
}

// MarketWsStreamClient 定义在 ws_market.go（完整广播-订阅实现）。
// UserWsStreamClient 定义在 ws_user.go（完整广播-订阅实现）。

type SportsWsStreamClient struct {
	WsStreamClient
}

func (*MyPolymarket) NewSportsWsStreamClient() *SportsWsStreamClient {
	return &SportsWsStreamClient{
		WsStreamClient: WsStreamClient{
			client:          &Client{},
			channel:         WS_SPORTS,
			writeMu:         &sync.Mutex{},
			isClose:         true,
			waitSubResult:   false,
			waitSubResultMu: &sync.Mutex{},

			reSubscribeMu: &sync.Mutex{},

			// sports 频道暂不使用 market 事件流
			resultChan: make(chan []byte),
			errChan:    make(chan error),
		},
	}
}

func (ws *WsStreamClient) ResultChan() chan []byte { return ws.resultChan }
func (ws *WsStreamClient) ErrChan() chan error     { return ws.errChan }
func (ws *WsStreamClient) GetConn() *websocket.Conn {
	return ws.conn
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiURL := handlerWsStreamRequestAPI(ws.channel)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiURL, ws.channel, ws.resultChan, ws.errChan, ws.writeMu)
		ws.conn = conn
		ws.isClose = false
		ws.isInit = false
		ws.clearWaitSubResult()
		log.Info("OpenConn success to ", apiURL)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	}
	// 重连：先显式关闭旧连接释放底层 TCP 资源，再建新连接。
	// handleResult goroutine 仍然存活并继续监听同一组 channel，无需重新启动。
	_ = ws.conn.Close()
	ws.conn = nil
	conn, err := wsStreamServe(apiURL, ws.channel, ws.resultChan, ws.errChan, ws.writeMu)
	ws.conn = conn
	ws.isInit = false
	ws.isClose = false
	ws.clearWaitSubResult()
	log.Info("Auto ReOpenConn success to ", apiURL)
	return err
}

func safeCloseChanStruct(ch chan struct{}) {
	if ch == nil {
		return
	}
	defer func() { recover() }()
	close(ch)
}

func safeCloseChanAny(ch chan any) {
	if ch == nil {
		return
	}
	defer func() { recover() }()
	close(ch)
}

func safeCloseChanErr(ch chan error) {
	if ch == nil {
		return
	}
	defer func() { recover() }()
	close(ch)
}

func safeCloseChanBytes(ch chan []byte) {
	if ch == nil {
		return
	}
	defer func() { recover() }()
	close(ch)
}

// sendWsCloseToAllSub 手动关闭连接后，关闭各订阅的 closeChan / resultChan / errChan，并重置 commonSubMap。
// 须在全局 resultChan 已关闭且 handleResult 即将/已经退出之后调用（见 Close 内顺序）。
func (ws *WsStreamClient) sendWsCloseToAllSub() {
	if ws == nil {
		return
	}
	seen := make(map[*Subscription]struct{})
	ws.commonSubMap.Range(func(_ string, sub *Subscription) bool {
		if sub == nil {
			return true
		}
		if _, ok := seen[sub]; ok {
			return true
		}
		seen[sub] = struct{}{}
		safeCloseChanStruct(sub.closeChan)
		safeCloseChanAny(sub.resultChan)
		safeCloseChanErr(sub.errChan)
		return true
	})
	ws.commonSubMap = NewMySyncMap[string, *Subscription]()
}

func (ws *WsStreamClient) Close() error {
	if ws == nil {
		return nil
	}
	ws.isClose = true
	ws.AutoReConnectTimes = 0

	var closeErr error
	if ws.conn != nil {
		if ws.writeMu != nil {
			ws.writeMu.Lock()
		}
		closeErr = ws.conn.Close()
		ws.conn = nil
		if ws.writeMu != nil {
			ws.writeMu.Unlock()
		}
	}

	// 先关掉底层读协程与 handleResult 共用的通道，避免继续向已关闭的订阅 channel 写入
	safeCloseChanBytes(ws.resultChan)
	safeCloseChanErr(ws.errChan)
	time.Sleep(30 * time.Millisecond)

	ws.sendWsCloseToAllSub()

	ws.resultChan = nil
	ws.errChan = nil
	ws.clearWaitSubResult()
	return closeErr
}

func (ws *WsStreamClient) SendPing() error {
	if ws == nil || ws.conn == nil || ws.isClose {
		return errors.New("websocket is closed")
	}
	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()
	return ws.conn.WriteMessage(websocket.TextMessage, []byte("PING"))
}

func subscribe(ws *WsStreamClient, args WsSubscribeReq) (*Subscription, error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, errors.New("websocket is close")
	}

	// if !ws.isInit {
	// 	switch ws.channel {
	// 	case WS_MARKET:
	// 		err := ws.initMarketWs(args.CustomFeatureEnabled, args.InitialDump, args.Level)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		ws.isInit = true
	// 	}
	// }

	ws.waitSubResultMu.Lock()
	if ws.waitSubResult {
		ws.waitSubResultMu.Unlock()
		return nil, errors.New("websocket is busy")
	}
	ws.waitSubResult = true
	ws.waitSubResultMu.Unlock()

	if err := ws.sendMessage(args); err != nil {
		ws.clearWaitSubResult()
		return nil, err
	}

	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()

	sub := &Subscription{
		SubId:      id,
		Args:       args,
		Ws:         ws,
		resultChan: make(chan any, 1024),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
	}

	return sub, nil
}

func (ws *WsStreamClient) DeferSub() {
	ws.clearWaitSubResult()
}

func (ws *WsStreamClient) clearWaitSubResult() {
	if ws == nil || ws.waitSubResultMu == nil {
		return
	}
	ws.waitSubResultMu.Lock()
	ws.waitSubResult = false
	ws.waitSubResultMu.Unlock()
}

func (ws *WsStreamClient) sendMessage(req WsSubscribeReq) error {
	if ws == nil || ws.conn == nil || ws.isClose {
		return errors.New("websocket is closed")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()
	return ws.conn.WriteMessage(websocket.TextMessage, data)
}

func (ws *WsStreamClient) WsConnURL() string {
	return handlerWsStreamRequestAPI(ws.channel)
}

func wsStreamServe(api string, channel WsChannelType, resultChan chan []byte, errChan chan error, writeMu *sync.Mutex) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	if WsUseProxy {
		proxy, _ := getBestProxyAndWeight()
		if proxy == nil {
			return nil, errors.New("no proxy available")
		}
		proxyUrl, err := url.Parse(proxy.ProxyUrl)
		if err != nil {
			return nil, err
		}
		dialer = &websocket.Dialer{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}
	c, _, err := dialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	// market/user 要求客户端主动发 PING；sports 由服务端发 ping，客户端只需回 pong。
	if WebsocketKeepalive && (channel == WS_MARKET || channel == WS_USER) {
		go keepAlive(c, WebsocketTimeout, writeMu)
	}
	go func() {
		// recover 防止在 Close() 关闭 errChan/resultChan 的同时，
		// 本 goroutine 仍尝试向已关闭 channel 发送数据而 panic。
		defer func() { recover() }()
		for {
			_, msg, e := c.ReadMessage()
			if e != nil {
				errChan <- e
				return
			}
			// sports 频道是文本心跳：服务器发 ping，客户端需回 pong。
			if channel == WS_SPORTS && strings.EqualFold(string(msg), "ping") {
				writeMu.Lock()
				_ = c.WriteMessage(websocket.TextMessage, []byte("pong"))
				writeMu.Unlock()
				continue
			}
			resultChan <- msg
		}
	}()
	return c, nil
}

func keepAlive(c *websocket.Conn, interval time.Duration, writeMu *sync.Mutex) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		writeMu.Lock()
		err := c.WriteMessage(websocket.TextMessage, []byte("PING"))
		writeMu.Unlock()
		if err != nil {
			return
		}
	}
}

func handlerWsStreamRequestAPI(channelType WsChannelType) string {
	host := PM_CLOB_API_WS_HOST
	if channelType == WS_SPORTS {
		host = PM_SPORTS_API_WS_HOST
	}
	u := url.URL{Scheme: "wss", Host: host, Path: getWsAPI(channelType)}
	return u.String()
}

func getWsAPI(channelType WsChannelType) string {
	switch channelType {
	case WS_MARKET:
		return PM_CLOB_API_WS_MARKET
	case WS_USER:
		return PM_CLOB_API_WS_USER
	case WS_SPORTS:
		return PM_SPORTS_API_WS_PATH
	default:
		return ""
	}
}

// reSubscribeForReconnect 仅供 User/Sports 等仍使用 commonSubMap 的客户端兼容使用。
// MarketWsStreamClient 通过 reconnectHook 处理重连，不走此路径。
func (ws *WsStreamClient) reSubscribeForReconnect() error {
	ws.reSubscribeMu.Lock()
	defer ws.reSubscribeMu.Unlock()

	isDoReSubscribe := map[int64]bool{}
	var wErr error

	ws.commonSubMap.Range(func(_ string, sub *Subscription) bool {
		if _, ok := isDoReSubscribe[sub.SubId]; ok {
			return true
		}
		reSub, err := subscribe(ws, sub.Args)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		sub.SubId = reSub.SubId
		log.Infof("reSubscribe Success: args:%v", reSub.Args)
		isDoReSubscribe[sub.SubId] = true
		time.Sleep(500 * time.Millisecond)
		return true
	})

	return wErr
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				// ws 标记为非关闭，且错误包含 EOF/close/reset 时自动重连。
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					const maxAttempts = 10
					var connErr error
					for attempt := 0; attempt < maxAttempts; attempt++ {
						connErr = ws.OpenConn()
						if connErr == nil {
							break
						}
						// 指数退避：500ms、1s、2s、4s … 上限 30s
						backoff := time.Duration(1<<uint(attempt)) * 500 * time.Millisecond
						if backoff > 30*time.Second {
							backoff = 30 * time.Second
						}
						log.Warnf("reconnect attempt %d/%d failed, retry in %v: %v", attempt+1, maxAttempts, backoff, connErr)
						time.Sleep(backoff)
					}
					if connErr != nil {
						log.Errorf("reconnect failed after %d attempts, giving up: %v", maxAttempts, connErr)
						return
					}
					ws.AutoReConnectTimes++
					if ws.reconnectHook != nil {
						go ws.reconnectHook()
					} else {
						go func() {
							if reSubErr := ws.reSubscribeForReconnect(); reSubErr != nil {
								log.Error("reSubscribe error: ", reSubErr)
							}
						}()
					}
				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					log.Error("resultChan is closed")
					return
				}
				if strings.EqualFold(string(data), "PONG") {
					continue
				}
				// log.Debug("msg: ", string(data))
				if ws.dataHandler != nil {
					ws.dataHandler(data)
				}
			}
		}
	}()
}
