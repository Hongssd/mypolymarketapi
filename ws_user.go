package mypolymarketapi

import (
	"errors"
	"sync"
)

// ---------------------------------------------------------------------------
// User 频道事件类型
// ---------------------------------------------------------------------------

type UserEventType string

const (
	UserEventTypeTrade UserEventType = "trade"
	UserEventTypeOrder UserEventType = "order"
)

func (e UserEventType) String() string { return string(e) }

const (
	inboundUserTradeCap = 256
	inboundUserOrderCap = 256
)

// userSubscribeAllMarketsKey 为内部扇出 map 键：表示订阅「全市场」（协议 JSON 省略 markets 字段）。
// 与 condition_id 不会冲突（链上 id 为 0x… 非空）。
const userSubscribeAllMarketsKey = ""

// ---------------------------------------------------------------------------
// UserWsStreamClient：带广播-订阅模型的 User Channel WS 客户端
//
// 协议与 https://docs.polymarket.com/api-reference/wss/user 一致：
//   - 全市场：type=user + auth，JSON 省略 markets（omitempty）
//   - 按 condition 过滤：首次 type=user + auth + markets；增量 operation + markets
// ---------------------------------------------------------------------------

type UserWsStreamClient struct {
	WsStreamClient

	// 协议订阅引用计数：condition_id → 活跃订阅者总数；userSubscribeAllMarketsKey 表示「全市场」协议槽
	protocolSubsMu sync.Mutex
	protocolSubs   map[string]int
	initialSubSent bool

	tradeMu      sync.RWMutex
	tradeStreams map[string]*marketStream[WsUserTrade]

	orderMu      sync.RWMutex
	orderStreams map[string]*marketStream[WsUserOrder]

	ownersMu sync.Mutex
	owners   map[string]*ownerRecord
}

func (*MyPolymarket) NewUserWsStreamClient(client *Client) *UserWsStreamClient {
	ws := &UserWsStreamClient{
		WsStreamClient: WsStreamClient{
			client:          client,
			channel:         WS_USER,
			writeMu:         &sync.Mutex{},
			isClose:         true,
			waitSubResult:   false,
			waitSubResultMu: &sync.Mutex{},
			reSubscribeMu:   &sync.Mutex{},
			resultChan:      make(chan []byte),
			errChan:         make(chan error),
		},
		protocolSubs: make(map[string]int),
		tradeStreams: make(map[string]*marketStream[WsUserTrade]),
		orderStreams: make(map[string]*marketStream[WsUserOrder]),
		owners:       make(map[string]*ownerRecord),
	}
	ws.WsStreamClient.dataHandler = ws.parseAndRoute
	ws.WsStreamClient.reconnectHook = ws.onReconnect
	return ws
}

func (ws *UserWsStreamClient) ensureProtocolSubscribed(marketIDs []string) {
	ws.protocolSubsMu.Lock()
	defer ws.protocolSubsMu.Unlock()

	// 全市场：仅 auth + type，Markets 为 nil → JSON 省略 markets
	if len(marketIDs) == 0 {
		key := userSubscribeAllMarketsKey
		wasZero := ws.protocolSubs[key] == 0
		ws.protocolSubs[key]++

		if !wasZero {
			return
		}

		creds := ws.client.ApiKeyCreds
		if creds == nil {
			ws.protocolSubs[key]--
			delete(ws.protocolSubs, key)
			log.Error("ensureProtocolSubscribed user: ApiKeyCreds is nil")
			return
		}
		auth := &WsAuth{
			ApiKey:     creds.APIKey,
			Secret:     creds.Secret,
			Passphrase: creds.Passphrase,
		}
		req := WsSubscribeReq{
			Type: "user",
			Auth: auth,
		}
		if err := ws.sendMessage(req); err != nil {
			log.Errorf("ensureProtocolSubscribed user sendMessage: %v", err)
			ws.protocolSubs[key]--
			delete(ws.protocolSubs, key)
			return
		}
		ws.initialSubSent = true
		return
	}

	newMarkets := make([]string, 0, len(marketIDs))
	for _, id := range marketIDs {
		if ws.protocolSubs[id] == 0 {
			newMarkets = append(newMarkets, id)
		}
		ws.protocolSubs[id]++
	}
	if len(newMarkets) == 0 {
		return
	}

	creds := ws.client.ApiKeyCreds
	if creds == nil {
		for _, id := range newMarkets {
			ws.protocolSubs[id]--
			if ws.protocolSubs[id] <= 0 {
				delete(ws.protocolSubs, id)
			}
		}
		log.Error("ensureProtocolSubscribed user: ApiKeyCreds is nil")
		return
	}

	var req WsSubscribeReq
	if !ws.initialSubSent {
		auth := &WsAuth{
			ApiKey:     creds.APIKey,
			Secret:     creds.Secret,
			Passphrase: creds.Passphrase,
		}
		m := append([]string(nil), newMarkets...)
		req = WsSubscribeReq{
			Type:    "user",
			Auth:    auth,
			Markets: &m,
		}
	} else {
		m := append([]string(nil), newMarkets...)
		req = WsSubscribeReq{
			Operation: SUBSCRIBE,
			Markets:   &m,
		}
	}

	if err := ws.sendMessage(req); err != nil {
		log.Errorf("ensureProtocolSubscribed user sendMessage: %v", err)
		for _, id := range newMarkets {
			ws.protocolSubs[id]--
			if ws.protocolSubs[id] <= 0 {
				delete(ws.protocolSubs, id)
			}
		}
		return
	}
	if !ws.initialSubSent {
		ws.initialSubSent = true
	}
}

func (ws *UserWsStreamClient) decrementProtocolSubs(marketIDs []string) {
	ws.protocolSubsMu.Lock()
	defer ws.protocolSubsMu.Unlock()

	if len(marketIDs) == 0 {
		key := userSubscribeAllMarketsKey
		ws.protocolSubs[key]--
		if ws.protocolSubs[key] <= 0 {
			delete(ws.protocolSubs, key)
		}
		if len(ws.protocolSubs) == 0 {
			ws.initialSubSent = false
		}
		// 全市场无对应「取消全部」的单独报文约定，不向服务端发 unsubscribe
		return
	}

	toUnsub := make([]string, 0, len(marketIDs))
	for _, id := range marketIDs {
		ws.protocolSubs[id]--
		if ws.protocolSubs[id] <= 0 {
			delete(ws.protocolSubs, id)
			toUnsub = append(toUnsub, id)
		}
	}
	if len(toUnsub) == 0 {
		if len(ws.protocolSubs) == 0 {
			ws.initialSubSent = false
		}
		return
	}
	m := append([]string(nil), toUnsub...)
	req := WsSubscribeReq{
		Operation: UNSUBSCRIBE,
		Markets:   &m,
	}
	if err := ws.sendMessage(req); err != nil {
		log.Errorf("decrementProtocolSubs user unsubscribe: %v", err)
	}
	if len(ws.protocolSubs) == 0 {
		ws.initialSubSent = false
	}
}

func (ws *UserWsStreamClient) onReconnect() {
	ws.protocolSubsMu.Lock()
	ws.initialSubSent = false
	var specifics []string
	wildcard := false
	for mid, n := range ws.protocolSubs {
		if n <= 0 {
			continue
		}
		if mid == userSubscribeAllMarketsKey {
			wildcard = true
		} else {
			specifics = append(specifics, mid)
		}
	}
	ws.protocolSubsMu.Unlock()

	if !wildcard && len(specifics) == 0 {
		return
	}
	creds := ws.client.ApiKeyCreds
	if creds == nil {
		log.Error("onReconnect user: ApiKeyCreds is nil")
		return
	}
	auth := &WsAuth{
		ApiKey:     creds.APIKey,
		Secret:     creds.Secret,
		Passphrase: creds.Passphrase,
	}
	var req WsSubscribeReq
	if wildcard {
		// 与文档一致：存在全市场订阅时重连仍省略 markets（含全市场 + 按 condition 并存时以全量为上界）
		req = WsSubscribeReq{
			Type: "user",
			Auth: auth,
		}
	} else {
		m := append([]string(nil), specifics...)
		req = WsSubscribeReq{
			Type:    "user",
			Auth:    auth,
			Markets: &m,
		}
	}
	if err := ws.sendMessage(req); err != nil {
		log.Errorf("onReconnect user re-subscribe: %v", err)
		return
	}
	ws.protocolSubsMu.Lock()
	ws.initialSubSent = true
	ws.protocolSubsMu.Unlock()
}

type wsUserEventTypePeek struct {
	EventType string `json:"event_type"`
}

func (ws *UserWsStreamClient) parseAndRoute(data []byte) {
	trimmed := data
	for len(trimmed) > 0 && (trimmed[0] == ' ' || trimmed[0] == '\t' || trimmed[0] == '\r' || trimmed[0] == '\n') {
		trimmed = trimmed[1:]
	}
	if len(trimmed) == 0 {
		return
	}

	var peek wsUserEventTypePeek
	if err := json.Unmarshal(trimmed, &peek); err != nil {
		return
	}

	switch UserEventType(peek.EventType) {
	case UserEventTypeTrade:
		tr := handleWsUserTrade(data)
		if tr == nil {
			return
		}
		ws.tradeMu.RLock()
		if s, ok := ws.tradeStreams[tr.Market]; ok {
			s.send(*tr)
		}
		if s, ok := ws.tradeStreams[userSubscribeAllMarketsKey]; ok {
			s.send(*tr)
		}
		ws.tradeMu.RUnlock()

	case UserEventTypeOrder:
		ord := handleWsUserOrder(data)
		if ord == nil {
			return
		}
		ws.orderMu.RLock()
		if s, ok := ws.orderStreams[ord.Market]; ok {
			s.send(*ord)
		}
		if s, ok := ws.orderStreams[userSubscribeAllMarketsKey]; ok {
			s.send(*ord)
		}
		ws.orderMu.RUnlock()
	}
}

// SubscribeTrade 订阅 trade 事件。
// marketIDs 非空时为按 condition_id 过滤；空切片表示全市场（协议 JSON 省略 markets，见官方文档）。
func (ws *UserWsStreamClient) SubscribeTrade(
	subscriberID string, marketIDs []string, bufSize int,
) (<-chan WsUserTrade, error) {
	if err := ws.validateParams(subscriberID, marketIDs, bufSize); err != nil {
		return nil, err
	}
	if err := ws.validateAllMarketsAfterSpecific(marketIDs); err != nil {
		return nil, err
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsUserTrade, bufSize)

	streamKeys := marketIDs
	if len(streamKeys) == 0 {
		streamKeys = []string{userSubscribeAllMarketsKey}
	}

	ws.tradeMu.Lock()
	for _, id := range streamKeys {
		s, ok := ws.tradeStreams[id]
		if !ok {
			s = newMarketStream[WsUserTrade](inboundUserTradeCap)
			ws.tradeStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.tradeMu.Unlock()

	ws.ensureProtocolSubscribed(marketIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), marketIDs...)
		ws.tradeMu.Lock()
		for _, id := range streamKeys {
			if s, ok := ws.tradeStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.tradeStreams, id)
				}
			}
		}
		ws.tradeMu.Unlock()
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// SubscribeOrder 订阅 order 事件。
// marketIDs 非空时为按 condition_id 过滤；空切片表示全市场（协议 JSON 省略 markets）。
func (ws *UserWsStreamClient) SubscribeOrder(
	subscriberID string, marketIDs []string, bufSize int,
) (<-chan WsUserOrder, error) {
	if err := ws.validateParams(subscriberID, marketIDs, bufSize); err != nil {
		return nil, err
	}
	if err := ws.validateAllMarketsAfterSpecific(marketIDs); err != nil {
		return nil, err
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsUserOrder, bufSize)

	streamKeys := marketIDs
	if len(streamKeys) == 0 {
		streamKeys = []string{userSubscribeAllMarketsKey}
	}

	ws.orderMu.Lock()
	for _, id := range streamKeys {
		s, ok := ws.orderStreams[id]
		if !ok {
			s = newMarketStream[WsUserOrder](inboundUserOrderCap)
			ws.orderStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.orderMu.Unlock()

	ws.ensureProtocolSubscribed(marketIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), marketIDs...)
		ws.orderMu.Lock()
		for _, id := range streamKeys {
			if s, ok := ws.orderStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.orderStreams, id)
				}
			}
		}
		ws.orderMu.Unlock()
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// Unsubscribe 取消 subscriberID 下全部订阅。
func (ws *UserWsStreamClient) Unsubscribe(subscriberID string) {
	ws.ownersMu.Lock()
	rec, ok := ws.owners[subscriberID]
	if ok {
		delete(ws.owners, subscriberID)
	}
	ws.ownersMu.Unlock()
	if ok && rec != nil {
		rec.cancelAll()
	}
}

func (ws *UserWsStreamClient) validateParams(subscriberID string, marketIDs []string, bufSize int) error {
	if subscriberID == "" {
		return errors.New("subscriberID cannot be empty")
	}
	if bufSize <= 0 {
		return errors.New("bufSize must be > 0")
	}
	if ws.isClose {
		return errors.New("websocket is not connected, call OpenConn first")
	}
	if ws.client == nil || ws.client.ApiKeyCreds == nil {
		return errors.New("client with ApiKeyCreds is required for user channel")
	}
	return nil
}

// validateAllMarketsAfterSpecific 禁止在尚未建立「全市场」协议订阅时，在已有按 condition 的协议引用下再申请全市场（文档未定义该首包路径）。
// 若连接上已通过省略 markets 建立全市场，则允许继续增加全市场订阅者（仅增加引用，不发新包）。
func (ws *UserWsStreamClient) validateAllMarketsAfterSpecific(marketIDs []string) error {
	if len(marketIDs) > 0 {
		return nil
	}
	ws.protocolSubsMu.Lock()
	defer ws.protocolSubsMu.Unlock()
	if ws.protocolSubs[userSubscribeAllMarketsKey] > 0 {
		return nil
	}
	for k, n := range ws.protocolSubs {
		if n <= 0 || k == userSubscribeAllMarketsKey {
			continue
		}
		return errors.New("user ws: all-markets subscribe (omit markets) is not supported after condition-specific subscription on the same connection unless all-markets was already active")
	}
	return nil
}

func (ws *UserWsStreamClient) addOwnerCancel(subscriberID string, cancel func()) {
	ws.ownersMu.Lock()
	if ws.owners[subscriberID] == nil {
		ws.owners[subscriberID] = &ownerRecord{}
	}
	ws.owners[subscriberID].add(cancel)
	ws.ownersMu.Unlock()
}
