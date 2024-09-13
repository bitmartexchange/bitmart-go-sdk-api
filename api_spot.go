package bitmart

// GetSpotCurrencies /** Get Currency List (v1)
func (cloudClient *CloudClient) GetSpotCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_CURRENCIES_URL, NONE)
}

// GetSpotSymbol /** Get List of Trading Pairs (v1)
func (cloudClient *CloudClient) GetSpotSymbol() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_URL, NONE)
}

// GetSpotSymbolDetail /** Get List of Trading Pair Details (v1)
func (cloudClient *CloudClient) GetSpotSymbolDetail() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_DETAILS_URL, NONE)
}

// GetSpotV3Tickers /** Get Ticker of All Pairs (V3)
func (cloudClient *CloudClient) GetSpotV3Tickers() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_V3_TICKERS_URL, NONE)
}

// GetSpotV3Ticker /** Get Ticker of a Trading Pair (V3)
func (cloudClient *CloudClient) GetSpotV3Ticker(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_V3_TICKER_URL, params, NONE)
}

// GetSpotV3LatestKline /** Get Latest K-Line (V3)
//
// Parameters:
// - symbol - Trading pair (e.g. BMX_USDT)
// - Options: before - Query timestamp (unit: second, e.g. 1525760116), query the data before this time
// - Options: after - Query timestamp (unit: second, e.g. 1525769116), query the data after this time
// - Options: step - k-line step, value [1, 3, 5, 15, 30, 45, 60, 120, 180, 240, 1440, 10080, 43200] unit: minute, default 1
// - Options: limit - Return number, the maximum value is 200, default is 100
func (cloudClient *CloudClient) GetSpotV3LatestKline(symbol string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_V3_LATEST_KLINE_URL, params, NONE)
}

// GetSpotV3HistoryKline /** Get History K-Line (V3)
//
// Parameters:
// - symbol - Trading pair (e.g. BMX_USDT)
// - Options: before - Query timestamp (unit: second, e.g. 1525760116), query the data before this time
// - Options: after - Query timestamp (unit: second, e.g. 1525769116), query the data after this time
// - Options: step - k-line step, value [1, 3, 5, 15, 30, 45, 60, 120, 180, 240, 1440, 10080, 43200] unit: minute, default 1
// - Options: limit - Return number, the maximum value is 200, default is 100
func (cloudClient *CloudClient) GetSpotV3HistoryKline(symbol string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_V3_HISTORY_KLINE_URL, params, NONE)
}

// GetSpotV3Book /** Get Depth (V3)
// Parameters:
// - symbol - Trading pair (e.g. BMX_USDT)
// - Options: limit - Order book depth per side. Maximum 50, e.g. 50 bids + 50 asks. Default returns to 35 depth data, e.g. 35 bids + 35 asks.
func (cloudClient *CloudClient) GetSpotV3Book(symbol string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_V3_BOOKS_URL, params, NONE)
}

// GetSpotV3Trade /** Get Recent Trades (V3)
// Parameters:
// - symbol - Trading pair (e.g. BMX_USDT)
// - Options: limit - Number of returned items, maximum is 50, default 50
func (cloudClient *CloudClient) GetSpotV3Trade(symbol string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_V3_TRADES_URL, params, NONE)
}

// GetSpotWallet /** Get Account Balance (KEYED)
func (cloudClient *CloudClient) GetSpotWallet() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_WALLET_URL, KEYED)
}

// Order /** Spot Order Parameters
type Order struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	ClientOrderId string `json:"client_order_id"`
	Size          string `json:"size"`
	Price         string `json:"price"`
	Notional      string `json:"notional"`
}

// PostSpotSubmitOrder /** New Order(v2) (SIGNED)
func (cloudClient *CloudClient) PostSpotSubmitOrder(order Order) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = order.Side
	params["type"] = order.Type
	if order.ClientOrderId != "" {
		params["client_order_id"] = order.ClientOrderId
	}
	if order.Size != "" {
		params["size"] = order.Size
	}
	if order.Price != "" {
		params["price"] = order.Price
	}
	if order.Notional != "" {
		params["notional"] = order.Notional
	}
	return cloudClient.requestWithParams(POST, API_SPOT_SUBMIT_ORDER_URL, params, SIGNED)
}

// MarginOrder /** Margin Order Parameters
type MarginOrder struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	ClientOrderId string `json:"clientOrderId"`
	Size          string `json:"size"`
	Price         string `json:"price"`
	Notional      string `json:"notional"`
}

// PostMarginSubmitOrder /** New Margin Order(v1) (SIGNED)
func (cloudClient *CloudClient) PostMarginSubmitOrder(order MarginOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = order.Side
	params["type"] = order.Type
	if order.ClientOrderId != "" {
		params["clientOrderId"] = order.ClientOrderId
	}
	if order.Size != "" {
		params["size"] = order.Size
	}
	if order.Price != "" {
		params["price"] = order.Price
	}
	if order.Notional != "" {
		params["notional"] = order.Notional
	}
	return cloudClient.requestWithParams(POST, API_SPOT_SUBMIT_MARGIN_ORDER_URL, params, SIGNED)
}

// PostSpotBatchOrders /** Batch New Order(v2) (SIGNED)
func (cloudClient *CloudClient) PostSpotBatchOrders(orderParams []Order) (*CloudResponse, error) {
	params := NewParams()
	params["order_params"] = orderParams
	return cloudClient.requestWithParams(POST, API_SPOT_BATCH_ORDERS_URL, params, SIGNED)
}

// PostSpotCancelOrder /** Cancel Order(v3) (SIGNED)
func (cloudClient *CloudClient) PostSpotCancelOrder(symbol string, orderId string, clientOrderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if orderId != "" {
		params["order_id"] = orderId
	}
	if clientOrderId != "" {
		params["client_order_id"] = clientOrderId
	}
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDER_URL, params, SIGNED)
}

// PostSpotCancelOrders /** Cancel Batch Order(v1) (SIGNED)
func (cloudClient *CloudClient) PostSpotCancelOrders(symbol string, side string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["side"] = side
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDERS_URL, params, SIGNED)
}

// GetSpotOrderByOrderId /** Query Order By Id (v4) (SIGNED)
func (cloudClient *CloudClient) GetSpotOrderByOrderId(orderId string, queryState string, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	params["orderId"] = orderId
	if queryState != "" {
		params["queryState"] = queryState
	}
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_SPOT_V4_QUERY_ORDER_BY_ID_URL, params, SIGNED)
}

// GetSpotOrderByClientOrderId /** Query Order By clientOrderId(v4) (SIGNED)
func (cloudClient *CloudClient) GetSpotOrderByClientOrderId(clientOrderId string, queryState string, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	params["clientOrderId"] = clientOrderId
	if queryState != "" {
		params["queryState"] = queryState
	}
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_SPOT_V4_QUERY_ORDER_BY_CLIENT_URL, params, SIGNED)
}

// GetSpotOpenOrders /** Query Open Orders (v4) (SIGNED)
func (cloudClient *CloudClient) GetSpotOpenOrders(symbol string, orderMode string, startTime int64, endTime int64, limit int, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}
	if orderMode != "" {
		params["orderMode"] = orderMode
	}
	if startTime != 0 {
		params["startTime"] = startTime
	}
	if endTime != 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_SPOT_V4_QUERY_OPEN_ORDERS_URL, params, SIGNED)
}

// GetSpotAccountOrders /** Query Account Orders (v4) (SIGNED)
func (cloudClient *CloudClient) GetSpotAccountOrders(symbol string, orderMode string, startTime int64, endTime int64, limit int, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}
	if orderMode != "" {
		params["orderMode"] = orderMode
	}
	if startTime != 0 {
		params["startTime"] = startTime
	}
	if endTime != 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_SPOT_V4_QUERY_HISTORY_ORDERS_URL, params, SIGNED)
}

// GetSpotAccountTradeList /** Account Trade List (v4) (SIGNED)
func (cloudClient *CloudClient) GetSpotAccountTradeList(symbol string, orderMode string, startTime int64, endTime int64, limit int, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}
	if orderMode != "" {
		params["orderMode"] = orderMode
	}
	if startTime != 0 {
		params["startTime"] = startTime
	}
	if endTime != 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_SPOT_V4_QUERY_TRADES_URL, params, SIGNED)
}

// GetSpotOrderTradeList /** Order Trade List(v4) (SIGNED)
func (cloudClient *CloudClient) GetSpotOrderTradeList(orderId string, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	params["orderId"] = orderId
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_SPOT_V4_QUERY_ORDER_TRADES_URL, params, SIGNED)
}
