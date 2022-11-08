package bitmart

// contracts

// GetContractTickers tickers
func (cloudClient *CloudClient) GetContractTickers(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["contract_symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_TICKERS_URL, params, NONE)
}

// GetContractDetails details
func (cloudClient *CloudClient) GetContractDetails(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_DETAILS_URL, params, NONE)
}

// GetContractDepth depth
func (cloudClient *CloudClient) GetContractDepth(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_DEPTH_URL, params, NONE)
}

// GetContractOpenInterest open-interest
func (cloudClient *CloudClient) GetContractOpenInterest(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_OPEN_INTEREST_URL, params, NONE)
}

// GetContractFundingRate funding-rate
func (cloudClient *CloudClient) GetContractFundingRate(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_FUNDING_RATE_URL, params, NONE)
}

// GetContractKline kline
func (cloudClient *CloudClient) GetContractKline(contractSymbol string, from, to, step int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["start_time"] = from
	params["end_time"] = to
	params["step"] = step
	return cloudClient.requestWithParams(GET, API_CONTRACT_KLINE_URL, params, NONE)
}

// GetContractAssetsDetail assets-detail
func (cloudClient *CloudClient) GetContractAssetsDetail() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_CONTRACT_ASSETS_DETAIL_URL, KEYED)
}

// GetContractOrder order
func (cloudClient *CloudClient) GetContractOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_URL, params, KEYED)
}

// GetContractOrderHistory order-history
func (cloudClient *CloudClient) GetContractOrderHistory(contractSymbol string, from, to int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["start_time"] = from
	params["end_time"] = to
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_HISTORY_URL, params, KEYED)
}

// GetContractPosition position
func (cloudClient *CloudClient) GetContractPosition(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_POSITION_URL, params, KEYED)
}

// GetContractTrades trades
func (cloudClient *CloudClient) GetContractTrades(contractSymbol string, from, to int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["start_time"] = from
	params["end_time"] = to
	return cloudClient.requestWithParams(GET, API_CONTRACT_TRADES_URL, params, KEYED)
}

// ContractOrder submit_contract order
type ContractOrder struct {
	Symbol   string `json:"symbol"`
	Type     string `json:"type,omitempty"`
	Side     int    `json:"side"`
	Leverage string `json:"leverage"`
	OpenType string `json:"open_type"`
	Price    string `json:"price"`
	Size     int    `json:"size"`
}

// PostContractSubmitOrder submit-order
func (cloudClient *CloudClient) PostContractSubmitOrder(order ContractOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	if order.Type == "" {
		params["type"] = "limit"
	}
	params["side"] = order.Side
	params["leverage"] = order.Leverage
	params["open_type"] = order.OpenType
	params["price"] = order.Price
	params["size"] = order.Size

	if order.Price != "" {
		params["price"] = order.Price
	}

	return cloudClient.requestWithParams(POST, API_CONTRACT_SUBMIT_ORDER_URL, params, SIGNED)
}

// PostContractCancelOrder cancel-order
func (cloudClient *CloudClient) PostContractCancelOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDER_URL, params, SIGNED)
}

// PostContractCancelOrders cancel-orders
func (cloudClient *CloudClient) PostContractCancelOrders(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDERS_URL, params, SIGNED)
}
