package bitmart

// currencies
func (cloudClient *CloudClient) getSpotCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_CURRENCIES_URL, NONE)
}

// symbols
func (cloudClient *CloudClient) getSpotSymbol() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_URL, NONE)
}

// symbols/details
func (cloudClient *CloudClient) getSpotSymbolDetail() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_SYMBOLS_DETAILS_URL, NONE)
}

// ticker
func (cloudClient *CloudClient) getSpotTicker(symbol string) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}

	return cloudClient.requestWithParams(GET, API_SPOT_TICKER_URL, params, NONE)
}

// steps
func (cloudClient *CloudClient) getSpotSteps() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_STEPS_URL, NONE)
}

// symbols/kline
func (cloudClient *CloudClient) getSpotSymbolKline(symbol string, from int64, to int64, step int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["from"] = from
	params["to"] = to
	params["step"] = step

	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_KLINE_URL, params, NONE)
}

// symbols/book
func (cloudClient *CloudClient) getSpotSymbolBook(symbol string, precision int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if precision != 0 {
		params["precision"] = precision
	}

	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_BOOK_URL, params, NONE)
}

// symbols/trades
func (cloudClient *CloudClient) getSpotSymbolTrade(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_SYMBOLS_TRADES_URL, params, NONE)
}

// wallet
func (cloudClient *CloudClient) getSpotWallet() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_WALLET_URL, KEYED)
}

// submit_order

type LimitBuyOrder struct {
	Symbol string `json:"symbol"`
	Size   string `json:"size"`
	Price  string `json:"price"`
}

func (cloudClient *CloudClient) postSpotSubmitLimitBuyOrder(order LimitBuyOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = "buy"
	params["type"] = "limit"
	params["size"] = order.Size
	params["price"] = order.Price
	return cloudClient.requestWithParams(POST, API_SPOT_SUBMIT_ORDER_URL, params, SIGNED)
}

type LimitSellOrder struct {
	Symbol string `json:"symbol"`
	Size   string `json:"size"`
	Price  string `json:"price"`
}

func (cloudClient *CloudClient) postSpotSubmitLimitSellOrder(order LimitSellOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = "sell"
	params["type"] = "limit"
	params["size"] = order.Size
	params["price"] = order.Price
	return cloudClient.requestWithParams(POST, API_SPOT_SUBMIT_ORDER_URL, params, SIGNED)
}

type MarketBuyOrder struct {
	Symbol   string `json:"symbol"`
	Notional string `json:"notional"`
}

func (cloudClient *CloudClient) postSpotSubmitMarketBuyOrder(order MarketBuyOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = "buy"
	params["type"] = "market"
	params["notional"] = order.Notional
	return cloudClient.requestWithParams(POST, API_SPOT_SUBMIT_ORDER_URL, params, SIGNED)
}

type MarketSellOrder struct {
	Symbol string `json:"symbol"`
	Size   string `json:"size"`
}

func (cloudClient *CloudClient) postSpotSubmitMarketSellOrder(order MarketSellOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	params["side"] = "sell"
	params["type"] = "market"
	params["size"] = order.Size
	return cloudClient.requestWithParams(POST, API_SPOT_SUBMIT_ORDER_URL, params, SIGNED)
}

// cancel_order
func (cloudClient *CloudClient) postSpotCancelOrder(symbol string, entrustId int64) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["entrust_id"] = entrustId
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDER_URL, params, SIGNED)
}


// cancel_orders
func (cloudClient *CloudClient) postSpotCancelOrders(symbol string, side string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["side"] = side
	return cloudClient.requestWithParams(POST, API_SPOT_CANCEL_ORDERS_URL, params, SIGNED)
}

// order_detail
func (cloudClient *CloudClient) getSpotOrderDetail(symbol string, entrustId int64) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["entrust_id"] = entrustId
	return cloudClient.requestWithParams(GET, API_SPOT_ORDER_DETAIL_URL, params, KEYED)
}


// orders
func (cloudClient *CloudClient) getSpotOrders(symbol string, offset int, limit int, status string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["offset"] = offset
	params["limit"] = limit
	params["status"] = status
	return cloudClient.requestWithParams(GET, API_SPOT_ORDERS_URL, params, KEYED)
}


// trades
func (cloudClient *CloudClient) getSpotHistoryTrades(symbol string, offset int, limit int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["offset"] = offset
	params["limit"] = limit
	return cloudClient.requestWithParams(GET, API_SPOT_TRADES_URL, params, KEYED)
}

func (cloudClient *CloudClient) getSpotOrderTrades(symbol string, entrustId int64) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["entrust_id"] = entrustId
	return cloudClient.requestWithParams(GET, API_SPOT_TRADES_URL, params, KEYED)
}
