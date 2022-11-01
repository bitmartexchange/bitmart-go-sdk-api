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
func (cloudClient *CloudClient) GetSpotTicker() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_TICKER_URL, NONE)
}

// ticker_detail
func (cloudClient *CloudClient) GetSpotTickerDetail(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol

	return cloudClient.requestWithParams(GET, API_SPOT_TICKER_DETAIL_URL, params, NONE)
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
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	ClientOrderId string `json:"client_order_id"`
	Size          string `json:"size"`
	Price         string `json:"price"`
	Notional      string `json:"notional"`
}

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

// margin/submit_order
type MarginOrder struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	ClientOrderId string `json:"clientOrderId"`
	Size          string `json:"size"`
	Price         string `json:"price"`
	Notional      string `json:"notional"`
}

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

// batch_orders
func (cloudClient *CloudClient) PostSpotBatchOrders(orderParams [1]Order) (*CloudResponse, error) {
	params := NewParams()
	params["order_params"] = orderParams
	return cloudClient.requestWithParams(POST, API_SPOT_BATCH_ORDERS_URL, params, SIGNED)
}

// cancel_order
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

// cancel_orders
func (cloudClient *CloudClient) PostSpotCancelOrders(symbol string, side string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["side"] = side
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDERS_URL, params, SIGNED)
}

// order_detail
func (cloudClient *CloudClient) GetSpotOrderDetail(orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_SPOT_ORDER_DETAIL_URL, params, KEYED)
}

// orders
func (cloudClient *CloudClient) GetSpotOrders(symbol string, orderMode string, N int, status string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if orderMode != "" {
		params["order_mode"] = orderMode
	}
	params["N"] = N
	params["status"] = status
	return cloudClient.requestWithParams(GET, API_SPOT_ORDERS_URL, params, KEYED)
}

// trades
func (cloudClient *CloudClient) GetSpotHistoryTrades(symbol string, orderMode string, N int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if orderMode != "" {
		params["order_mode"] = orderMode
	}
	params["N"] = N
	return cloudClient.requestWithParams(GET, API_SPOT_TRADES_URL, params, KEYED)
}

func (cloudClient *CloudClient) GetSpotOrderTrades(symbol string, orderMode string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if orderMode != "" {
		params["order_mode"] = orderMode
	}
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_SPOT_TRADES_URL, params, KEYED)
}
