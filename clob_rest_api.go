package mypolymarketapi

import (
	"errors"
	"fmt"
	"strings"
)

// GET Get order book
func (c *CLOBRestClient) NewCLOBGetOrderBook() *CLOBGetOrderBookAPI {
	return &CLOBGetOrderBookAPI{
		client: c,
		req:    &CLOBGetOrderBookReq{},
	}
}

func (api *CLOBGetOrderBookAPI) Do() (*PolyMarketRestRes[CLOBGetOrderBookRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetOrderBook])
	return pmCallAPI[CLOBGetOrderBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Get order books (request body)
func (c *CLOBRestClient) NewCLOBGetOrderBooks() *CLOBGetOrderBooksAPI {
	return &CLOBGetOrderBooksAPI{
		client: c,
		req:    &CLOBGetOrderBooksReq{},
	}
}

func (api *CLOBGetOrderBooksAPI) Do() (*PolyMarketRestRes[CLOBGetOrderBooksRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetOrderBooks])
	return pmCallAPI[CLOBGetOrderBooksRes](api.client.c, url, body, POST)
}

// GET Get market price
func (c *CLOBRestClient) NewCLOBGetMarketPrice() *CLOBGetMarketPriceAPI {
	return &CLOBGetMarketPriceAPI{
		client: c,
		req:    &CLOBGetMarketPriceReq{},
	}
}

func (api *CLOBGetMarketPriceAPI) Do() (*PolyMarketRestRes[CLOBGetMarketPriceRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	if api.req.Side == nil {
		return nil, errors.New("side is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetMarketPrice])
	return pmCallAPI[CLOBGetMarketPriceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get market prices (query parameters)
func (c *CLOBRestClient) NewCLOBGetMarketPricesGet() *CLOBGetMarketPricesGetAPI {
	return &CLOBGetMarketPricesGetAPI{
		client: c,
		req:    &CLOBGetMarketPricesGetReq{},
	}
}

func (api *CLOBGetMarketPricesGetAPI) Do() (*PolyMarketRestRes[CLOBGetMarketPricesGetRes], error) {
	if api.req.TokenIDs == nil {
		return nil, errors.New("token_ids is required")
	}
	if api.req.Sides == nil {
		return nil, errors.New("sides is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetMarketPricesGet])
	return pmCallAPI[CLOBGetMarketPricesGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Get market prices (request body)
func (c *CLOBRestClient) NewCLOBGetMarketPricesPost() *CLOBGetMarketPricesPostAPI {
	return &CLOBGetMarketPricesPostAPI{
		client: c,
		req:    &CLOBGetMarketPricesPostReq{},
	}
}

func (api *CLOBGetMarketPricesPostAPI) Do() (*PolyMarketRestRes[CLOBGetMarketPricesPostRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetMarketPricesPost])
	return pmCallAPI[CLOBGetMarketPricesPostRes](api.client.c, url, body, POST)
}

// GET Get midpoint price
func (c *CLOBRestClient) NewCLOBGetMidpointPrice() *CLOBGetMidpointPriceAPI {
	return &CLOBGetMidpointPriceAPI{
		client: c,
		req:    &CLOBGetMidpointPriceReq{},
	}
}

func (api *CLOBGetMidpointPriceAPI) Do() (*PolyMarketRestRes[CLOBGetMidpointPriceRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetMidpointPrice])
	return pmCallAPI[CLOBGetMidpointPriceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get midpoint prices (query parameters)
func (c *CLOBRestClient) NewCLOBGetMidpointPricesGet() *CLOBGetMidpointPricesGetAPI {
	return &CLOBGetMidpointPricesGetAPI{
		client: c,
		req:    &CLOBGetMidpointPricesGetReq{},
	}
}

func (api *CLOBGetMidpointPricesGetAPI) Do() (*PolyMarketRestRes[CLOBGetMidpointPricesGetRes], error) {
	if api.req.TokenIDs == nil {
		return nil, errors.New("token_ids is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetMidpointPricesGet])
	return pmCallAPI[CLOBGetMidpointPricesGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Get midpoint prices (request body)
func (c *CLOBRestClient) NewCLOBGetMidpointPricesPost() *CLOBGetMidpointPricesPostAPI {
	return &CLOBGetMidpointPricesPostAPI{
		client: c,
		req:    &CLOBGetMidpointPricesPostReq{},
	}
}

func (api *CLOBGetMidpointPricesPostAPI) Do() (*PolyMarketRestRes[CLOBGetMidpointPricesPostRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetMidpointPricesPost])
	return pmCallAPI[CLOBGetMidpointPricesPostRes](api.client.c, url, body, POST)
}

// GET Get spread
func (c *CLOBRestClient) NewCLOBGetSpread() *CLOBGetSpreadAPI {
	return &CLOBGetSpreadAPI{
		client: c,
		req:    &CLOBGetSpreadReq{},
	}
}

func (api *CLOBGetSpreadAPI) Do() (*PolyMarketRestRes[CLOBGetSpreadRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetSpread])
	return pmCallAPI[CLOBGetSpreadRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Get spreads
func (c *CLOBRestClient) NewCLOBGetSpreads() *CLOBGetSpreadsAPI {
	return &CLOBGetSpreadsAPI{
		client: c,
		req:    &CLOBGetSpreadsReq{},
	}
}

func (api *CLOBGetSpreadsAPI) Do() (*PolyMarketRestRes[CLOBGetSpreadsRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetSpreads])
	return pmCallAPI[CLOBGetSpreadsRes](api.client.c, url, body, POST)
}

// GET Get last trade price
func (c *CLOBRestClient) NewCLOBGetLastTradePrice() *CLOBGetLastTradePriceAPI {
	return &CLOBGetLastTradePriceAPI{
		client: c,
		req:    &CLOBGetLastTradePriceReq{},
	}
}

func (api *CLOBGetLastTradePriceAPI) Do() (*PolyMarketRestRes[CLOBGetLastTradePriceRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetLastTradePrice])
	return pmCallAPI[CLOBGetLastTradePriceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get last trade prices (query parameters)
func (c *CLOBRestClient) NewCLOBGetLastTradesPricesGet() *CLOBGetLastTradesPricesGetAPI {
	return &CLOBGetLastTradesPricesGetAPI{
		client: c,
		req:    &CLOBGetLastTradesPricesGetReq{},
	}
}

func (api *CLOBGetLastTradesPricesGetAPI) Do() (*PolyMarketRestRes[CLOBGetLastTradesPricesGetRes], error) {
	if api.req.TokenIDs == nil {
		return nil, errors.New("token_ids is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetLastTradesPricesGet])
	return pmCallAPI[CLOBGetLastTradesPricesGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Get last trade prices (request body)
func (c *CLOBRestClient) NewCLOBGetLastTradesPricesPost() *CLOBGetLastTradesPricesPostAPI {
	return &CLOBGetLastTradesPricesPostAPI{
		client: c,
		req:    &CLOBGetLastTradesPricesPostReq{},
	}
}

func (api *CLOBGetLastTradesPricesPostAPI) Do() (*PolyMarketRestRes[CLOBGetLastTradesPricesPostRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetLastTradesPricesPost])
	return pmCallAPI[CLOBGetLastTradesPricesPostRes](api.client.c, url, body, POST)
}

// GET Get prices history
func (c *CLOBRestClient) NewCLOBGetPricesHistory() *CLOBGetPricesHistoryAPI {
	return &CLOBGetPricesHistoryAPI{
		client: c,
		req:    &CLOBGetPricesHistoryReq{},
	}
}

func (api *CLOBGetPricesHistoryAPI) Do() (*PolyMarketRestRes[CLOBGetPricesHistoryRes], error) {
	if api.req.Market == nil {
		return nil, errors.New("market is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetPricesHistory])
	return pmCallAPI[CLOBGetPricesHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Get batch prices history
func (c *CLOBRestClient) NewCLOBGetBatchPricesHistory() *CLOBGetBatchPricesHistoryAPI {
	return &CLOBGetBatchPricesHistoryAPI{
		client: c,
		req:    &CLOBGetBatchPricesHistoryReq{},
	}
}

func (api *CLOBGetBatchPricesHistoryAPI) Do() (*PolyMarketRestRes[CLOBGetBatchPricesHistoryRes], error) {
	if len(api.req.Markets) == 0 {
		return nil, errors.New("markets is required")
	}
	if len(api.req.Markets) > 20 {
		return nil, errors.New("markets exceeds maximum of 20 per request")
	}

	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetBatchPricesHistory])
	return pmCallAPI[CLOBGetBatchPricesHistoryRes](api.client.c, url, body, POST)
}

// GET Get fee rate (query parameter)
func (c *CLOBRestClient) NewCLOBGetFeeRate() *CLOBGetFeeRateAPI {
	return &CLOBGetFeeRateAPI{
		client: c,
		req:    &CLOBGetFeeRateReq{},
	}
}

func (api *CLOBGetFeeRateAPI) Do() (*PolyMarketRestRes[CLOBGetFeeRateRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetFeeRate])
	return pmCallAPI[CLOBGetFeeRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get fee rate (path parameter)
func (c *CLOBRestClient) NewCLOBGetFeeRateByPath() *CLOBGetFeeRateByPathAPI {
	return &CLOBGetFeeRateByPathAPI{
		client: c,
		req:    &CLOBGetFeeRateByPathReq{},
	}
}

func (api *CLOBGetFeeRateByPathAPI) Do() (*PolyMarketRestRes[CLOBGetFeeRateRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	path := strings.Replace(CLOBAPITypeMap[CLOBGetFeeRateByPath], "{token_id}", *api.req.TokenID, 1)
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, path)
	return pmCallAPI[CLOBGetFeeRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get tick size (query parameter)
func (c *CLOBRestClient) NewCLOBGetTickSize() *CLOBGetTickSizeAPI {
	return &CLOBGetTickSizeAPI{
		client: c,
		req:    &CLOBGetTickSizeReq{},
	}
}

func (api *CLOBGetTickSizeAPI) Do() (*PolyMarketRestRes[CLOBGetTickSizeRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetTickSize])
	return pmCallAPI[CLOBGetTickSizeRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get tick size (path parameter)
func (c *CLOBRestClient) NewCLOBGetTickSizeByPath() *CLOBGetTickSizeByPathAPI {
	return &CLOBGetTickSizeByPathAPI{
		client: c,
		req:    &CLOBGetTickSizeByPathReq{},
	}
}

func (api *CLOBGetTickSizeByPathAPI) Do() (*PolyMarketRestRes[CLOBGetTickSizeRes], error) {
	if api.req.TokenID == nil {
		return nil, errors.New("token_id is required")
	}
	path := strings.Replace(CLOBAPITypeMap[CLOBGetTickSizeByPath], "{token_id}", *api.req.TokenID, 1)
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, path)
	return pmCallAPI[CLOBGetTickSizeRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get CLOB market info (path parameter: condition_id)
func (c *CLOBRestClient) NewCLOBGetClobMarketInfo() *CLOBGetClobMarketInfoAPI {
	return &CLOBGetClobMarketInfoAPI{
		client: c,
		req:    &CLOBGetClobMarketInfoReq{},
	}
}

func (api *CLOBGetClobMarketInfoAPI) Do() (*PolyMarketRestRes[CLOBGetClobMarketInfoRes], error) {
	if api.req.ConditionID == nil {
		return nil, errors.New("condition_id is required")
	}
	path := strings.Replace(CLOBAPITypeMap[CLOBGetClobMarketInfo], "{condition_id}", *api.req.ConditionID, 1)
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, path)
	return pmCallAPI[CLOBGetClobMarketInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get server time
func (c *CLOBRestClient) NewCLOBGetServerTime() *CLOBGetServerTimeAPI {
	return &CLOBGetServerTimeAPI{client: c}
}

func (api *CLOBGetServerTimeAPI) Do() (*PolyMarketRestRes[CLOBGetServerTimeRes], error) {
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBGetServerTime])
	return pmCallAPI[CLOBGetServerTimeRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get trades
func (c *CLOBRestClient) NewCLOBGetTrades() *CLOBGetTradesAPI {
	return &CLOBGetTradesAPI{
		client: c,
		req:    &CLOBGetTradesReq{},
	}
}

func (api *CLOBGetTradesAPI) Do() (*PolyMarketRestRes[CLOBGetTradesRes], error) {
	if api.req == nil || api.req.MakerAddress == nil {
		return nil, errors.New("maker_address is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetTrades])
	return pmCallAPI[CLOBGetTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get balance and allowance
func (c *CLOBRestClient) NewCLOBGetBalanceAllowance() *CLOBGetBalanceAllowanceAPI {
	return &CLOBGetBalanceAllowanceAPI{
		client: c,
		req:    &CLOBGetBalanceAllowanceReq{},
	}
}

func (api *CLOBGetBalanceAllowanceAPI) Do() (*PolyMarketRestRes[CLOBGetBalanceAllowanceRes], error) {
	if api.req == nil || api.req.AssetType == nil {
		return nil, errors.New("asset_type is required")
	}
	if api.req.SignatureType != nil {
		if *api.req.SignatureType < 0 || *api.req.SignatureType > 2 {
			return nil, errors.New("signature_type must be 0/1/2")
		}
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetBalanceAllowance])
	return pmCallAPIWithSecret[CLOBGetBalanceAllowanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST Post a new order
func (c *CLOBRestClient) NewCLOBPostOrder() *CLOBPostOrderAPI {
	return &CLOBPostOrderAPI{
		client: c,
		req: &CLOBPostOrderReq{
			// 与 @polymarket/clob-client-v2 的 orderToJsonV2 默认一致，避免 omit 掉 false 或与官方 JSON 形状不一致。
			PostOnly:  GetPointer(false),
			DeferExec: GetPointer(false),
		},
	}
}

// CLOB V2 迁移文档要求 wire 上 metadata/builder 为完整 bytes32 hex；空串易导致服务端按非 V2 解析并返回 order_version_mismatch。
const clobV2Bytes32ZeroWire = "0x0000000000000000000000000000000000000000000000000000000000000000"

func (api *CLOBPostOrderAPI) Do() (*PolyMarketRestRes[CLOBPostOrderRes], error) {
	if api.req.Order == nil {
		return nil, errors.New("order is required")
	}

	order := *api.req.Order
	err := order.SignCLOBOrderV2EIP712(api.client.c.Wallet.PrivateKey, POLYGON_MAINNET_CHAIN_ID, false)
	if err != nil {
		return nil, fmt.Errorf("sign order: %w", err)
	}
	if order.Metadata == nil || strings.TrimSpace(*order.Metadata) == "" {
		order.Metadata = GetPointer(clobV2Bytes32ZeroWire)
	}
	if order.Builder == nil || strings.TrimSpace(*order.Builder) == "" {
		order.Builder = GetPointer(clobV2Bytes32ZeroWire)
	}
	// 签名与 wire 修正写在局部副本上，须写回 req，否则 Marshal 仍为无签名旧载荷 → order_version_mismatch
	*api.req.Order = order

	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBPostOrder])
	return pmCallAPIWithSecret[CLOBPostOrderRes](api.client.c, url, body, POST)
}

// POST Post multiple orders
func (c *CLOBRestClient) NewCLOBPostMultipleOrders() *CLOBPostMultipleOrdersAPI {
	return &CLOBPostMultipleOrdersAPI{
		client: c,
		req:    &CLOBPostMultipleOrdersReq{},
	}
}

func (api *CLOBPostMultipleOrdersAPI) Do() (*PolyMarketRestRes[CLOBPostMultipleOrdersRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}
	if len(*api.req) > 15 {
		return nil, errors.New("orders exceeds maximum of 15 per request")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBPostMultipleOrders])
	return pmCallAPIWithSecret[CLOBPostMultipleOrdersRes](api.client.c, url, body, POST)
}

// DELETE Cancel single order
func (c *CLOBRestClient) NewCLOBCancelSingleOrder() *CLOBCancelSingleOrderAPI {
	return &CLOBCancelSingleOrderAPI{
		client: c,
		req:    &CLOBCancelSingleOrderReq{},
	}
}

func (api *CLOBCancelSingleOrderAPI) Do() (*PolyMarketRestRes[CLOBCancelSingleOrderRes], error) {
	if api.req.OrderID == nil {
		return nil, errors.New("order_id is required")
	}

	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBCancelSingleOrder])
	return pmCallAPIWithSecret[CLOBCancelSingleOrderRes](api.client.c, url, body, DELETE)
}

// DELETE Cancel multiple orders
func (c *CLOBRestClient) NewCLOBCancelMultipleOrders() *CLOBCancelMultipleOrdersAPI {
	return &CLOBCancelMultipleOrdersAPI{
		client: c,
		req:    &CLOBCancelMultipleOrdersReq{},
	}
}

func (api *CLOBCancelMultipleOrdersAPI) Do() (*PolyMarketRestRes[CLOBCancelMultipleOrdersRes], error) {
	if api.req == nil || len(*api.req) == 0 {
		return nil, errors.New("request body is required")
	}
	if len(*api.req) > 3000 {
		return nil, errors.New("order_ids exceeds maximum of 3000 per request")
	}

	body, err := json.Marshal(*api.req)
	if err != nil {
		return nil, err
	}

	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBCancelMultiple])
	return pmCallAPIWithSecret[CLOBCancelMultipleOrdersRes](api.client.c, url, body, DELETE)
}

// DELETE Cancel all orders
func (c *CLOBRestClient) NewCLOBCancelAllOrders() *CLOBCancelAllOrdersAPI {
	return &CLOBCancelAllOrdersAPI{client: c}
}

func (api *CLOBCancelAllOrdersAPI) Do() (*PolyMarketRestRes[CLOBCancelAllOrdersRes], error) {
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBCancelAll])
	return pmCallAPIWithSecret[CLOBCancelAllOrdersRes](api.client.c, url, NIL_REQBODY, DELETE)
}

// GET Get single order by ID
func (c *CLOBRestClient) NewCLOBGetSingleOrderByID() *CLOBGetSingleOrderByIDAPI {
	return &CLOBGetSingleOrderByIDAPI{
		client: c,
		req:    &CLOBGetSingleOrderByIDReq{},
	}
}

func (api *CLOBGetSingleOrderByIDAPI) Do() (*PolyMarketRestRes[CLOBGetSingleOrderByIDRes], error) {
	if api.req.OrderID == nil {
		return nil, errors.New("order_id is required")
	}

	path := strings.Replace(CLOBAPITypeMap[CLOBGetSingleOrderByID], "{orderID}", *api.req.OrderID, 1)
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, path)
	return pmCallAPIWithSecret[CLOBGetSingleOrderByIDRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get user orders
func (c *CLOBRestClient) NewCLOBGetUserOrders() *CLOBGetUserOrdersAPI {
	return &CLOBGetUserOrdersAPI{
		client: c,
		req:    &CLOBGetUserOrdersReq{},
	}
}

func (api *CLOBGetUserOrdersAPI) Do() (*PolyMarketRestRes[CLOBGetUserOrdersRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetUserOrders])
	return pmCallAPIWithSecret[CLOBGetUserOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get order scoring status
func (c *CLOBRestClient) NewCLOBGetOrderScoring() *CLOBGetOrderScoringAPI {
	return &CLOBGetOrderScoringAPI{
		client: c,
		req:    &CLOBGetOrderScoringReq{},
	}
}

func (api *CLOBGetOrderScoringAPI) Do() (*PolyMarketRestRes[CLOBGetOrderScoringRes], error) {
	if api.req.OrderID == nil {
		return nil, errors.New("order_id is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, CLOB_REST, api.req, CLOBAPITypeMap[CLOBGetOrderScoring])
	return pmCallAPIWithSecret[CLOBGetOrderScoringRes](api.client.c, url, NIL_REQBODY, GET)
}
