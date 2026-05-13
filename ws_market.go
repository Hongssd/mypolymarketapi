package mypolymarketapi

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
)

// ---------------------------------------------------------------------------
// 事件类型常量
// ---------------------------------------------------------------------------

type EventType string

const (
	EventTypeBook           EventType = "book"
	EventTypePriceChange    EventType = "price_change"
	EventTypeLastTradePrice EventType = "last_trade_price"
	EventTypeTickSizeChange EventType = "tick_size_change"
	EventTypeBestBidAsk     EventType = "best_bid_ask"
	EventTypeNewMarket      EventType = "new_market"
	EventTypeMarketResolved EventType = "market_resolved"
)

func (e EventType) String() string { return string(e) }

// ---------------------------------------------------------------------------
// 内部分发队列容量
// ---------------------------------------------------------------------------

const (
	inboundBookCap           = 256
	inboundPriceChangeCap    = 1024
	inboundLastTradeCap      = 256
	inboundTickSizeCap       = 64
	inboundBestBidAskCap     = 256
	inboundNewMarketCap      = 64
	inboundMarketResolvedCap = 64
)

// ---------------------------------------------------------------------------
// marketSubKeyCounter：每次 Subscribe 调用生成唯一内部 key
// ---------------------------------------------------------------------------

var marketSubKeyCounter atomic.Int64

func nextMarketSubKey() string {
	return strconv.FormatInt(marketSubKeyCounter.Add(1), 10)
}

// ---------------------------------------------------------------------------
// closeChanGeneric：安全关闭泛型 channel（忽略 already-closed panic）
// ---------------------------------------------------------------------------

func closeChanGeneric[T any](ch chan T) {
	if ch == nil {
		return
	}
	defer func() { recover() }()
	close(ch)
}

// ---------------------------------------------------------------------------
// marketStream[T]：单类型泛型扇出流
//
// 一个 run() goroutine 持续从 inboundCh 读取，并以非阻塞方式广播到全部已注册订阅者。
// ---------------------------------------------------------------------------

type marketStream[T any] struct {
	inboundCh chan T
	mu        sync.RWMutex      // 保护 subs map
	subs      map[string]chan T // subKey → 订阅者 channel
	stopCh    chan struct{}
}

func newMarketStream[T any](inboundCap int) *marketStream[T] {
	return &marketStream[T]{
		inboundCh: make(chan T, inboundCap),
		subs:      make(map[string]chan T),
		stopCh:    make(chan struct{}),
	}
}

// run 为扇出 goroutine，需以 go stream.run() 启动。
func (s *marketStream[T]) run() {
	for {
		select {
		case msg, ok := <-s.inboundCh:
			if !ok {
				return
			}
			s.mu.RLock()
			for _, ch := range s.subs {
				// 非阻塞投递：慢消费者丢消息，不影响其他订阅者或分发 goroutine。
				select {
				case ch <- msg:
				default:
				}
			}
			s.mu.RUnlock()
		case <-s.stopCh:
			return
		}
	}
}

// send 向 inboundCh 非阻塞写入；满则丢弃。
func (s *marketStream[T]) send(msg T) {
	select {
	case s.inboundCh <- msg:
	default:
		log.Warn("marketStream: inboundCh full, dropping message")
	}
}

// addSub 注册新订阅者，返回其专属 channel。
func (s *marketStream[T]) addSub(subKey string, ch chan T) {
	s.mu.Lock()
	s.subs[subKey] = ch
	s.mu.Unlock()
}

// removeSub 移除订阅者。
// 因 run() 在 RLock 期间写 channel，而 removeSub 持 WriteLock，
// 故 close(ch) 可安全在 Unlock 后执行——run() 不会再持有该 ch 的引用。
// 返回 stream 是否已无任何订阅者。
func (s *marketStream[T]) removeSub(subKey string) (ch chan T, empty bool) {
	s.mu.Lock()
	ch = s.subs[subKey]
	delete(s.subs, subKey)
	empty = len(s.subs) == 0
	s.mu.Unlock()
	return ch, empty
}

// isEmpty 报告流是否没有任何订阅者。
func (s *marketStream[T]) isEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.subs) == 0
}

// stop 终止 run() goroutine。
func (s *marketStream[T]) stop() {
	defer func() { recover() }()
	close(s.stopCh)
}

// ---------------------------------------------------------------------------
// ownerRecord：记录某一 subscriberID 的全部清理函数
// ---------------------------------------------------------------------------

type ownerRecord struct {
	mu      sync.Mutex
	cancels []func()
}

func (r *ownerRecord) add(cancel func()) {
	r.mu.Lock()
	r.cancels = append(r.cancels, cancel)
	r.mu.Unlock()
}

func (r *ownerRecord) cancelAll() {
	r.mu.Lock()
	cancels := r.cancels
	r.cancels = nil
	r.mu.Unlock()
	for _, c := range cancels {
		c()
	}
}

// ---------------------------------------------------------------------------
// MarketWsStreamClient：带广播-订阅模型的市场 WS 客户端
// ---------------------------------------------------------------------------

type MarketWsStreamClient struct {
	WsStreamClient

	// 连接级别协议参数
	level                int // 1/2/3，默认 2
	customFeatureEnabled bool

	// 协议订阅引用计数：assetID → 当前活跃订阅者总数（跨所有事件类型）
	protocolSubsMu sync.Mutex
	protocolSubs   map[string]int
	initialSubSent bool // 是否已发出过首次 type=market 订阅消息

	// 各事件类型的流 map（key = assetID）
	bookMu      sync.RWMutex
	bookStreams map[string]*marketStream[WsMarketOrderBook]

	pcMu               sync.RWMutex
	priceChangeStreams map[string]*marketStream[WsMarketPriceChange]

	ltMu             sync.RWMutex
	lastTradeStreams map[string]*marketStream[WsMarketLastTradePrice]

	tsMu            sync.RWMutex
	tickSizeStreams map[string]*marketStream[WsMarketTickSizeChange]

	bbaMu             sync.RWMutex
	bestBidAskStreams map[string]*marketStream[WsMarketBestBidAsk]

	// 全局事件流（无 assetID，需 customFeatureEnabled=true）
	nmMu            sync.RWMutex
	newMarketStream *marketStream[WsMarketNewMarket] // 首次订阅时懒创建

	mrMu                 sync.RWMutex
	marketResolvedStream *marketStream[WsMarketMarketResolved] // 首次订阅时懒创建

	// subscriberID → 该 ID 下所有订阅的清理函数集合
	ownersMu sync.Mutex
	owners   map[string]*ownerRecord
}

func (*MyPolymarket) NewMarketWsStreamClient() *MarketWsStreamClient {
	ws := &MarketWsStreamClient{
		WsStreamClient: WsStreamClient{
			isInit:          false,
			client:          &Client{},
			channel:         WS_MARKET,
			writeMu:         &sync.Mutex{},
			isClose:         true,
			waitSubResult:   false,
			waitSubResultMu: &sync.Mutex{},
			reSubscribeMu:   &sync.Mutex{},
			commonSubMap:    NewMySyncMap[string, *Subscription](),
			resultChan:      make(chan []byte),
			errChan:         make(chan error),
		},
		level:              2,
		protocolSubs:       make(map[string]int),
		bookStreams:        make(map[string]*marketStream[WsMarketOrderBook]),
		priceChangeStreams: make(map[string]*marketStream[WsMarketPriceChange]),
		lastTradeStreams:   make(map[string]*marketStream[WsMarketLastTradePrice]),
		tickSizeStreams:    make(map[string]*marketStream[WsMarketTickSizeChange]),
		bestBidAskStreams:  make(map[string]*marketStream[WsMarketBestBidAsk]),
		owners:             make(map[string]*ownerRecord),
	}
	ws.WsStreamClient.dataHandler = ws.parseAndRoute
	ws.WsStreamClient.reconnectHook = ws.onReconnect
	return ws
}

// SetLevel 设置连接级别的订阅深度（1/2/3），须在 OpenConn 前调用。
func (ws *MarketWsStreamClient) SetLevel(level int) {
	ws.level = level
}

// SetCustomFeatureEnabled 启用 custom feature（new_market / market_resolved / best_bid_ask）。
func (ws *MarketWsStreamClient) SetCustomFeatureEnabled(enabled bool) {
	ws.customFeatureEnabled = enabled
}

// ---------------------------------------------------------------------------
// 内部：协议订阅管理
// ---------------------------------------------------------------------------

// ensureProtocolSubscribed 为新 assetID 增加引用计数，并在首次出现时向交易所发送协议订阅消息。
func (ws *MarketWsStreamClient) ensureProtocolSubscribed(assetIDs []string) {
	ws.protocolSubsMu.Lock()
	defer ws.protocolSubsMu.Unlock()

	newAssets := make([]string, 0, len(assetIDs))
	for _, id := range assetIDs {
		if ws.protocolSubs[id] == 0 {
			newAssets = append(newAssets, id)
		}
		ws.protocolSubs[id]++
	}

	if len(newAssets) == 0 {
		return
	}

	var req WsSubscribeReq
	if !ws.initialSubSent {
		req = WsSubscribeReq{
			Type:                 "market",
			AssetsIDs:            newAssets,
			Level:                ws.level,
			CustomFeatureEnabled: ws.customFeatureEnabled,
		}
	} else {
		req = WsSubscribeReq{
			Operation: SUBSCRIBE,
			AssetsIDs: newAssets,
		}
	}

	if err := ws.sendMessage(req); err != nil {
		log.Errorf("ensureProtocolSubscribed sendMessage: %v", err)
		// 回滚引用计数，避免 ghost 条目占用协议槽
		for _, id := range newAssets {
			ws.protocolSubs[id]--
			if ws.protocolSubs[id] == 0 {
				delete(ws.protocolSubs, id)
			}
		}
		return
	}
	if !ws.initialSubSent {
		ws.initialSubSent = true
	}
}

// decrementProtocolSubs 减少 assetID 的引用计数，归零时向交易所发送取消订阅消息。
func (ws *MarketWsStreamClient) decrementProtocolSubs(assetIDs []string) {
	ws.protocolSubsMu.Lock()
	defer ws.protocolSubsMu.Unlock()

	toUnsubscribe := make([]string, 0, len(assetIDs))
	for _, id := range assetIDs {
		ws.protocolSubs[id]--
		if ws.protocolSubs[id] <= 0 {
			delete(ws.protocolSubs, id)
			toUnsubscribe = append(toUnsubscribe, id)
		}
	}

	if len(toUnsubscribe) == 0 {
		return
	}

	req := WsSubscribeReq{
		Operation: UNSUBSCRIBE,
		AssetsIDs: toUnsubscribe,
	}
	if err := ws.sendMessage(req); err != nil {
		log.Errorf("decrementProtocolSubs unsubscribe: %v", err)
	}
}

// onReconnect 在自动重连成功后重新发送全量协议订阅（由 reconnectHook 触发）。
func (ws *MarketWsStreamClient) onReconnect() {
	ws.protocolSubsMu.Lock()
	ws.initialSubSent = false
	assets := make([]string, 0, len(ws.protocolSubs))
	for assetID, count := range ws.protocolSubs {
		if count > 0 {
			assets = append(assets, assetID)
		}
	}
	ws.protocolSubsMu.Unlock()

	if len(assets) == 0 {
		return
	}

	req := WsSubscribeReq{
		Type:                 "market",
		AssetsIDs:            assets,
		Level:                ws.level,
		CustomFeatureEnabled: ws.customFeatureEnabled,
	}
	if err := ws.sendMessage(req); err != nil {
		log.Errorf("onReconnect re-subscribe: %v", err)
		return
	}

	ws.protocolSubsMu.Lock()
	ws.initialSubSent = true
	ws.protocolSubsMu.Unlock()
}

// ---------------------------------------------------------------------------
// 内部：消息解析与路由
// ---------------------------------------------------------------------------

// wsEventTypePeek 仅用于快速提取 event_type 字段。
type wsEventTypePeek struct {
	EventType string `json:"event_type"`
}

// parseAndRoute 解析原始 WS 消息并分发到对应 eventType 的流。
// 由 WsStreamClient.dataHandler 调用，运行在 handleResult goroutine 中。
func (ws *MarketWsStreamClient) parseAndRoute(data []byte) {
	// 去掉首尾空白后快速判断是否为 JSON 数组（initial_dump 会推送 book 数组）
	trimmed := data
	for len(trimmed) > 0 && (trimmed[0] == ' ' || trimmed[0] == '\t' || trimmed[0] == '\r' || trimmed[0] == '\n') {
		trimmed = trimmed[1:]
	}
	if len(trimmed) == 0 {
		return
	}

	if trimmed[0] == '[' {
		// book 数组（initial_dump）
		books := handleWsMarketOrderBook(data)
		ws.bookMu.RLock()
		for i := range books {
			if s, ok := ws.bookStreams[books[i].AssetID]; ok {
				s.send(books[i])
			}
		}
		ws.bookMu.RUnlock()
		return
	}

	var peek wsEventTypePeek
	if err := json.Unmarshal(trimmed, &peek); err != nil {
		return
	}

	switch EventType(peek.EventType) {
	case EventTypeBook:
		books := handleWsMarketOrderBook(data)
		ws.bookMu.RLock()
		for i := range books {
			if s, ok := ws.bookStreams[books[i].AssetID]; ok {
				s.send(books[i])
			}
		}
		ws.bookMu.RUnlock()

	case EventTypePriceChange:
		items := handleWsMarketPriceChange(data)
		ws.pcMu.RLock()
		for i := range items {
			if s, ok := ws.priceChangeStreams[items[i].AssetID]; ok {
				s.send(items[i])
			}
		}
		ws.pcMu.RUnlock()

	case EventTypeLastTradePrice:
		ltp := handleWsMarketLastTradePrice(data)
		if ltp == nil {
			return
		}
		ws.ltMu.RLock()
		if s, ok := ws.lastTradeStreams[ltp.AssetID]; ok {
			s.send(*ltp)
		}
		ws.ltMu.RUnlock()

	case EventTypeTickSizeChange:
		tsc := handleWsMarketTickSizeChange(data)
		if tsc == nil {
			return
		}
		ws.tsMu.RLock()
		if s, ok := ws.tickSizeStreams[tsc.AssetID]; ok {
			s.send(*tsc)
		}
		ws.tsMu.RUnlock()

	case EventTypeBestBidAsk:
		bba := handleWsMarketBestBidAsk(data)
		if bba == nil {
			return
		}
		ws.bbaMu.RLock()
		if s, ok := ws.bestBidAskStreams[bba.AssetID]; ok {
			s.send(*bba)
		}
		ws.bbaMu.RUnlock()

	case EventTypeNewMarket:
		nm := handleWsMarketNewMarket(data)
		if nm == nil {
			return
		}
		ws.nmMu.RLock()
		if ws.newMarketStream != nil {
			ws.newMarketStream.send(*nm)
		}
		ws.nmMu.RUnlock()

	case EventTypeMarketResolved:
		mr := handleWsMarketMarketResolved(data)
		if mr == nil {
			return
		}
		ws.mrMu.RLock()
		if ws.marketResolvedStream != nil {
			ws.marketResolvedStream.send(*mr)
		}
		ws.mrMu.RUnlock()
	}
}

// ---------------------------------------------------------------------------
// 公开 Subscribe/Unsubscribe 接口
// ---------------------------------------------------------------------------

// SubscribeOrderBook 订阅一组 assetID 的 orderbook 事件。
// 同一 subscriberID 可多次调用（不同 assetIDs），每次返回独立 channel。
// 调用 Unsubscribe(subscriberID) 可一次性取消该 ID 下所有订阅。
func (ws *MarketWsStreamClient) SubscribeOrderBook(
	subscriberID string, assetIDs []string, bufSize int,
) (<-chan WsMarketOrderBook, error) {
	if err := ws.validateParams(subscriberID, assetIDs, bufSize); err != nil {
		return nil, err
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketOrderBook, bufSize)

	ws.bookMu.Lock()
	for _, id := range assetIDs {
		s, ok := ws.bookStreams[id]
		if !ok {
			s = newMarketStream[WsMarketOrderBook](inboundBookCap)
			ws.bookStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.bookMu.Unlock()

	ws.ensureProtocolSubscribed(assetIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), assetIDs...)
		ws.bookMu.Lock()
		for _, id := range idsCopy {
			if s, ok := ws.bookStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.bookStreams, id)
				}
			}
		}
		ws.bookMu.Unlock()
		// ch 已从所有流的 subs 中移除后再关闭，确保 run() 不会向已关闭的 ch 写入
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// SubscribePriceChange 订阅一组 assetID 的 price_change 事件。
func (ws *MarketWsStreamClient) SubscribePriceChange(
	subscriberID string, assetIDs []string, bufSize int,
) (<-chan WsMarketPriceChange, error) {
	if err := ws.validateParams(subscriberID, assetIDs, bufSize); err != nil {
		return nil, err
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketPriceChange, bufSize)

	ws.pcMu.Lock()
	for _, id := range assetIDs {
		s, ok := ws.priceChangeStreams[id]
		if !ok {
			s = newMarketStream[WsMarketPriceChange](inboundPriceChangeCap)
			ws.priceChangeStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.pcMu.Unlock()

	ws.ensureProtocolSubscribed(assetIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), assetIDs...)
		ws.pcMu.Lock()
		for _, id := range idsCopy {
			if s, ok := ws.priceChangeStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.priceChangeStreams, id)
				}
			}
		}
		ws.pcMu.Unlock()
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// SubscribeLastTradePrice 订阅一组 assetID 的 last_trade_price 事件。
func (ws *MarketWsStreamClient) SubscribeLastTradePrice(
	subscriberID string, assetIDs []string, bufSize int,
) (<-chan WsMarketLastTradePrice, error) {
	if err := ws.validateParams(subscriberID, assetIDs, bufSize); err != nil {
		return nil, err
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketLastTradePrice, bufSize)

	ws.ltMu.Lock()
	for _, id := range assetIDs {
		s, ok := ws.lastTradeStreams[id]
		if !ok {
			s = newMarketStream[WsMarketLastTradePrice](inboundLastTradeCap)
			ws.lastTradeStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.ltMu.Unlock()

	ws.ensureProtocolSubscribed(assetIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), assetIDs...)
		ws.ltMu.Lock()
		for _, id := range idsCopy {
			if s, ok := ws.lastTradeStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.lastTradeStreams, id)
				}
			}
		}
		ws.ltMu.Unlock()
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// SubscribeTickSizeChange 订阅一组 assetID 的 tick_size_change 事件。
func (ws *MarketWsStreamClient) SubscribeTickSizeChange(
	subscriberID string, assetIDs []string, bufSize int,
) (<-chan WsMarketTickSizeChange, error) {
	if err := ws.validateParams(subscriberID, assetIDs, bufSize); err != nil {
		return nil, err
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketTickSizeChange, bufSize)

	ws.tsMu.Lock()
	for _, id := range assetIDs {
		s, ok := ws.tickSizeStreams[id]
		if !ok {
			s = newMarketStream[WsMarketTickSizeChange](inboundTickSizeCap)
			ws.tickSizeStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.tsMu.Unlock()

	ws.ensureProtocolSubscribed(assetIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), assetIDs...)
		ws.tsMu.Lock()
		for _, id := range idsCopy {
			if s, ok := ws.tickSizeStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.tickSizeStreams, id)
				}
			}
		}
		ws.tsMu.Unlock()
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// SubscribeBestBidAsk 订阅一组 assetID 的 best_bid_ask 事件。
// 注意：此事件需要 custom_feature_enabled=true；调用前请先 SetCustomFeatureEnabled(true)。
func (ws *MarketWsStreamClient) SubscribeBestBidAsk(
	subscriberID string, assetIDs []string, bufSize int,
) (<-chan WsMarketBestBidAsk, error) {
	if err := ws.validateParams(subscriberID, assetIDs, bufSize); err != nil {
		return nil, err
	}
	if !ws.customFeatureEnabled {
		return nil, errors.New("best_bid_ask requires SetCustomFeatureEnabled(true)")
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketBestBidAsk, bufSize)

	ws.bbaMu.Lock()
	for _, id := range assetIDs {
		s, ok := ws.bestBidAskStreams[id]
		if !ok {
			s = newMarketStream[WsMarketBestBidAsk](inboundBestBidAskCap)
			ws.bestBidAskStreams[id] = s
			go s.run()
		}
		s.addSub(subKey, ch)
	}
	ws.bbaMu.Unlock()

	ws.ensureProtocolSubscribed(assetIDs)
	ws.addOwnerCancel(subscriberID, func() {
		idsCopy := append([]string(nil), assetIDs...)
		ws.bbaMu.Lock()
		for _, id := range idsCopy {
			if s, ok := ws.bestBidAskStreams[id]; ok {
				if _, empty := s.removeSub(subKey); empty {
					s.stop()
					delete(ws.bestBidAskStreams, id)
				}
			}
		}
		ws.bbaMu.Unlock()
		closeChanGeneric(ch)
		ws.decrementProtocolSubs(idsCopy)
	})
	return ch, nil
}

// SubscribeNewMarket 订阅全局 new_market 事件（需 custom_feature_enabled=true）。
func (ws *MarketWsStreamClient) SubscribeNewMarket(
	subscriberID string, bufSize int,
) (<-chan WsMarketNewMarket, error) {
	if subscriberID == "" {
		return nil, errors.New("subscriberID cannot be empty")
	}
	if bufSize <= 0 {
		return nil, errors.New("bufSize must be > 0")
	}
	if !ws.customFeatureEnabled {
		return nil, errors.New("new_market requires SetCustomFeatureEnabled(true)")
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketNewMarket, bufSize)

	ws.nmMu.Lock()
	if ws.newMarketStream == nil {
		ws.newMarketStream = newMarketStream[WsMarketNewMarket](inboundNewMarketCap)
		go ws.newMarketStream.run()
	}
	ws.newMarketStream.addSub(subKey, ch)
	ws.nmMu.Unlock()

	ws.addOwnerCancel(subscriberID, func() {
		ws.nmMu.Lock()
		if ws.newMarketStream != nil {
			if _, empty := ws.newMarketStream.removeSub(subKey); empty {
				ws.newMarketStream.stop()
				ws.newMarketStream = nil
			}
		}
		ws.nmMu.Unlock()
		closeChanGeneric(ch)
	})
	return ch, nil
}

// SubscribeMarketResolved 订阅全局 market_resolved 事件（需 custom_feature_enabled=true）。
func (ws *MarketWsStreamClient) SubscribeMarketResolved(
	subscriberID string, bufSize int,
) (<-chan WsMarketMarketResolved, error) {
	if subscriberID == "" {
		return nil, errors.New("subscriberID cannot be empty")
	}
	if bufSize <= 0 {
		return nil, errors.New("bufSize must be > 0")
	}
	if !ws.customFeatureEnabled {
		return nil, errors.New("market_resolved requires SetCustomFeatureEnabled(true)")
	}
	subKey := nextMarketSubKey()
	ch := make(chan WsMarketMarketResolved, bufSize)

	ws.mrMu.Lock()
	if ws.marketResolvedStream == nil {
		ws.marketResolvedStream = newMarketStream[WsMarketMarketResolved](inboundMarketResolvedCap)
		go ws.marketResolvedStream.run()
	}
	ws.marketResolvedStream.addSub(subKey, ch)
	ws.mrMu.Unlock()

	ws.addOwnerCancel(subscriberID, func() {
		ws.mrMu.Lock()
		if ws.marketResolvedStream != nil {
			if _, empty := ws.marketResolvedStream.removeSub(subKey); empty {
				ws.marketResolvedStream.stop()
				ws.marketResolvedStream = nil
			}
		}
		ws.mrMu.Unlock()
		closeChanGeneric(ch)
	})
	return ch, nil
}

// Unsubscribe 取消 subscriberID 下的全部订阅：
//   - 从各 assetID 流中移除订阅者
//   - 关闭所有已分配的 channel（消费方 range ch 会自动退出）
//   - 若某 assetID 已无任何订阅者，向交易所发送 operation=unsubscribe
func (ws *MarketWsStreamClient) Unsubscribe(subscriberID string) {
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

// ---------------------------------------------------------------------------
// 内部辅助方法
// ---------------------------------------------------------------------------

func (ws *MarketWsStreamClient) validateParams(subscriberID string, assetIDs []string, bufSize int) error {
	if subscriberID == "" {
		return errors.New("subscriberID cannot be empty")
	}
	if len(assetIDs) == 0 {
		return errors.New("assetIDs cannot be empty")
	}
	if bufSize <= 0 {
		return errors.New("bufSize must be > 0")
	}
	if ws.isClose {
		return errors.New("websocket is not connected, call OpenConn first")
	}
	return nil
}

func (ws *MarketWsStreamClient) addOwnerCancel(subscriberID string, cancel func()) {
	ws.ownersMu.Lock()
	if ws.owners[subscriberID] == nil {
		ws.owners[subscriberID] = &ownerRecord{}
	}
	ws.owners[subscriberID].add(cancel)
	ws.ownersMu.Unlock()
}
