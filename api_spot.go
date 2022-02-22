package bitmart

// currencies
func (cloudClient *CloudClient) GetSpotCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_CURRENCIES_URL, NONE)
}

// symbols
func (cloudClient *CloudClient) GetSpotSymbol() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_URL, NONE)
}

// symbols/details
func (cloudClient *CloudClient) GetSpotSymbolDetail() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_DETAILS_URL, NONE)
}

// ticker
func (cloudClient *CloudClient) GetSpotTicker(symbol string) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}

	return cloudClient.requestWithParams(GET, API_SPOT_TICKER_URL, params, NONE)
}

// steps
func (cloudClient *CloudClient) GetSpotSteps() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_STEPS_URL, NONE)
}

// symbols/kline
func (cloudClient *CloudClient) GetSpotSymbolKline(symbol string, from int, to int, step int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["from"] = from
	params["to"] = to
	params["step"] = step

	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_KLINE_URL, params, NONE)
}

// symbols/book
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

// symbols/trades
func (cloudClient *CloudClient) GetSpotSymbolTrade(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_TRADES_URL, params, NONE)
}

// wallet
func (cloudClient *CloudClient) GetSpotWallet() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_WALLET_URL, KEYED)
}

// submit_order
type Order struct {
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	Type     string `json:"type"`
	Size     string `json:"size"`
	Price    string `json:"price"`
	Notional string `json:"notional"`
}

func (cloudClient *CloudClient) PostSpotSubmitOrder(order Order) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = order.Side
	params["type"] = order.Type
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

// batch_orders
func (cloudClient *CloudClient) PostSpotBatchOrders(orderParams [1]Order) (*CloudResponse, error) {
	params := NewParams()
	params["orderParams"] = orderParams
	return cloudClient.requestWithParams(POST, API_SPOT_BATCH_ORDERS_URL, params, SIGNED)
}

// cancel_order
func (cloudClient *CloudClient) PostSpotCancelOrder(symbol string, orderId int64) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDER_URL, params, SIGNED)
}

// cancel_orders
func (cloudClient *CloudClient) PostSpotCancelOrders(symbol string, side string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["side"] = side
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDERS_URL, params, SIGNED)
}

// order_detail
func (cloudClient *CloudClient) GetSpotOrderDetail(symbol string, orderId int64) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_SPOT_ORDER_DETAIL_URL, params, KEYED)
}

// orders
func (cloudClient *CloudClient) GetSpotOrders(symbol string, N int, status string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["N"] = N
	params["status"] = status
	return cloudClient.requestWithParams(GET, API_SPOT_ORDERS_URL, params, KEYED)
}

// trades
func (cloudClient *CloudClient) GetSpotHistoryTrades(symbol string, offset int, limit int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["offset"] = offset
	params["limit"] = limit
	return cloudClient.requestWithParams(GET, API_SPOT_TRADES_URL, params, KEYED)
}

func (cloudClient *CloudClient) GetSpotOrderTrades(symbol string, orderId int64) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_SPOT_TRADES_URL, params, KEYED)
}
