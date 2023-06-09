package bitmart

// GetSpotCurrencies /** Get Currency List
func (cloudClient *CloudClient) GetSpotCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_CURRENCIES_URL, NONE)
}

// GetSpotSymbol /** Get List of Trading Pairs
func (cloudClient *CloudClient) GetSpotSymbol() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_URL, NONE)
}

// GetSpotSymbolDetail /** Get List of Trading Pair Details
func (cloudClient *CloudClient) GetSpotSymbolDetail() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_DETAILS_URL, NONE)
}

// GetSpotTicker /** Get Ticker of All Pairs (V2)
func (cloudClient *CloudClient) GetSpotTicker() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_TICKER_URL, NONE)
}

// GetSpotTickerDetail /** Get Ticker of a Trading Pair
func (cloudClient *CloudClient) GetSpotTickerDetail(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol

	return cloudClient.requestWithParams(GET, API_SPOT_TICKER_DETAIL_URL, params, NONE)
}

// GetSpotSteps /** Get K-Line Step
func (cloudClient *CloudClient) GetSpotSteps() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_STEPS_URL, NONE)
}

// GetSpotSymbolKline /** Get K-Line
func (cloudClient *CloudClient) GetSpotSymbolKline(symbol string, from int64, to int64, step int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["from"] = from
	params["to"] = to
	params["step"] = step

	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_KLINE_URL, params, NONE)
}

// GetSpotSymbolBook /** Get Depth
func (cloudClient *CloudClient) GetSpotSymbolBook(symbol string, precision int, size int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if precision != 0 {
		params["precision"] = precision
	}

	if size != 0 {
		params["size"] = size
	}

	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_BOOK_URL, params, NONE)
}

// GetSpotSymbolTrade /** Get Recent Trades
func (cloudClient *CloudClient) GetSpotSymbolTrade(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_TRADES_URL, params, NONE)
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
func (cloudClient *CloudClient) PostSpotBatchOrders(orderParams [1]Order) (*CloudResponse, error) {
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
