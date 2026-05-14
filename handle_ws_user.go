package mypolymarketapi

// WsUserMakerOrder 表示 trade 消息中 maker 侧成交明细。
type WsUserMakerOrder struct {
	AssetID       string `json:"asset_id"`
	MatchedAmount string `json:"matched_amount"`
	OrderID       string `json:"order_id"`
	Outcome       string `json:"outcome"`
	Owner         string `json:"owner"`
	Price         string `json:"price"`
}

// WsUserTrade 对应 user 频道 event_type=trade（type 字段为 TRADE 等业务态）。
type WsUserTrade struct {
	EventType    string              `json:"event_type"` // trade
	Type         string              `json:"type"`       // TRADE
	ID           string              `json:"id"`
	AssetID      string              `json:"asset_id"`
	Market       string              `json:"market"` // condition_id
	LastUpdate   string              `json:"last_update"`
	MatchTime    string              `json:"matchtime"`
	Outcome      string              `json:"outcome"`
	Owner        string              `json:"owner"`
	TradeOwner   string              `json:"trade_owner"`
	Price        string              `json:"price"`
	Side         string              `json:"side"`   // BUY / SELL
	Size         string              `json:"size"`
	Status       string              `json:"status"` // MATCHED / MINED / CONFIRMED / RETRYING / FAILED
	TakerOrderID string              `json:"taker_order_id"`
	Timestamp    string              `json:"timestamp"`
	MakerOrders  []WsUserMakerOrder `json:"maker_orders"`
}

func handleWsUserTrade(data []byte) *WsUserTrade {
	var t WsUserTrade
	if err := json.Unmarshal(data, &t); err != nil {
		log.Error("unmarshal user trade error: ", err)
		return nil
	}
	return &t
}

// WsUserOrder 对应 user 频道 event_type=order（type 为 PLACEMENT / UPDATE / CANCELLATION）。
type WsUserOrder struct {
	EventType       string   `json:"event_type"` // order
	Type            string   `json:"type"`       // PLACEMENT / UPDATE / CANCELLATION
	ID              string   `json:"id"`
	AssetID         string   `json:"asset_id"`
	Market          string   `json:"market"` // condition_id
	OrderOwner      string   `json:"order_owner"`
	Owner           string   `json:"owner"`
	OriginalSize    string   `json:"original_size"`
	SizeMatched     string   `json:"size_matched"`
	Outcome         string   `json:"outcome"`
	Price           string   `json:"price"`
	Side            string   `json:"side"`
	Timestamp       string   `json:"timestamp"`
	AssociateTrades []string `json:"associate_trades"` // 文档为 null 或列表，按 string 兼容
}

func handleWsUserOrder(data []byte) *WsUserOrder {
	var o WsUserOrder
	if err := json.Unmarshal(data, &o); err != nil {
		log.Error("unmarshal user order error: ", err)
		return nil
	}
	return &o
}
