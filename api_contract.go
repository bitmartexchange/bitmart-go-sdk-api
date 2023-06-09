package bitmart

// contracts

// GetContractDetails details /** Get Contract Details
func (cloudClient *CloudClient) GetContractDetails(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_DETAILS_URL, params, NONE)
}

// GetContractDepth depth /** Get Market Depth
func (cloudClient *CloudClient) GetContractDepth(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_DEPTH_URL, params, NONE)
}

// GetContractOpenInterest open-interest /** Get Futures Open Interest
func (cloudClient *CloudClient) GetContractOpenInterest(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_OPEN_INTEREST_URL, params, NONE)
}

// GetContractFundingRate funding-rate /** Get Current Funding Rate
func (cloudClient *CloudClient) GetContractFundingRate(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_FUNDING_RATE_URL, params, NONE)
}

// GetContractKline kline /** Get K-line
func (cloudClient *CloudClient) GetContractKline(contractSymbol string, from, to, step int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["start_time"] = from
	params["end_time"] = to
	params["step"] = step
	return cloudClient.requestWithParams(GET, API_CONTRACT_KLINE_URL, params, NONE)
}

// GetContractAssetsDetail assets-detail /** Get Contract Assets (KEYED)
func (cloudClient *CloudClient) GetContractAssetsDetail() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_CONTRACT_ASSETS_DETAIL_URL, KEYED)
}

// GetContractOrder order /** Get Order Detail (KEYED)
func (cloudClient *CloudClient) GetContractOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_URL, params, KEYED)
}

// GetContractOrderHistory order-history /** Get Order History (KEYED)
func (cloudClient *CloudClient) GetContractOrderHistory(contractSymbol string, from, to int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["start_time"] = from
	params["end_time"] = to
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_HISTORY_URL, params, KEYED)
}

// GetContractPosition position /** Get Current Position (KEYED)
func (cloudClient *CloudClient) GetContractPosition(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_POSITION_URL, params, KEYED)
}

// GetContractTrades trades /** Get Order Trade (KEYED)
func (cloudClient *CloudClient) GetContractTrades(contractSymbol string, from, to int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["start_time"] = from
	params["end_time"] = to
	return cloudClient.requestWithParams(GET, API_CONTRACT_TRADES_URL, params, KEYED)
}

func (cloudClient *CloudClient) GetContractTransferList(contractSymbol string, timeStart, timeEnd int64, page, limit, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	if timeStart > 0 {
		params["time_start"] = timeStart
	}
	if timeEnd > 0 {
		params["time_end"] = timeEnd
	}
	params["page"] = page
	params["limit"] = limit
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_CONTRACT_TRANSFER_LIST_URL, params, SIGNED)
}

// ContractOrder submit_contract order params
type ContractOrder struct {
	Symbol        string `json:"symbol"`
	ClientOrderId string `json:"client_order_id"`
	Type          string `json:"type,omitempty"`
	Side          int    `json:"side"`
	Leverage      string `json:"leverage"`
	OpenType      string `json:"open_type"`
	Mode          int    `json:"mode"`
	Price         string `json:"price"`
	Size          int    `json:"size"`
}

// PostContractSubmitOrder submit-order /** Submit Order (SIGNED)
func (cloudClient *CloudClient) PostContractSubmitOrder(order ContractOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = order.Symbol
	if order.ClientOrderId != "" {
		params["client_order_id"] = order.ClientOrderId
	}
	if order.Type == "" {
		params["type"] = "limit"
	}
	params["side"] = order.Side
	params["leverage"] = order.Leverage
	params["open_type"] = order.OpenType
	if order.Mode == 0 {
		params["mode"] = 1
	}
	params["size"] = order.Size
	if order.Price != "" {
		params["price"] = order.Price
	}

	return cloudClient.requestWithParams(POST, API_CONTRACT_SUBMIT_ORDER_URL, params, SIGNED)
}

// PostContractCancelOrder cancel-order /** Cancel Order (SIGNED)
func (cloudClient *CloudClient) PostContractCancelOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDER_URL, params, SIGNED)
}

// PostContractCancelOrders cancel-orders /** Cancel All Orders (SIGNED)
func (cloudClient *CloudClient) PostContractCancelOrders(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDERS_URL, params, SIGNED)
}

type ContractPlanOrder struct {
	Symbol         string `json:"symbol"`
	Type           string `json:"type,omitempty"`
	Side           int    `json:"side"`
	Leverage       string `json:"leverage"`
	OpenType       string `json:"open_type"`
	Mode           int    `json:"mode,omitempty"`
	Size           int    `json:"size"`
	TriggerPrice   string `json:"trigger_price"`
	ExecutivePrice string `json:"executive_price"`
	PriceWay       int    `json:"price_way"`
	PriceType      int    `json:"price_type"`
}

// PostContractPlanOrder plan-order /** Submit Plan Order (SIGNED)
func (cloudClient *CloudClient) PostContractPlanOrder(planOrder ContractPlanOrder) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = planOrder.Symbol
	if planOrder.Type == "" {
		params["type"] = "limit"
	}
	params["side"] = planOrder.Side
	params["leverage"] = planOrder.Leverage
	params["open_type"] = planOrder.OpenType
	if planOrder.Mode == 0 {
		params["mode"] = 1
	}
	params["size"] = planOrder.Size
	params["trigger_price"] = planOrder.TriggerPrice
	if planOrder.ExecutivePrice != "" {
		params["executive_price"] = planOrder.ExecutivePrice
	}
	params["price_way"] = planOrder.PriceWay
	params["price_type"] = planOrder.PriceType
	return cloudClient.requestWithParams(POST, API_CONTRACT_SUBMIT_PLAN_ORDER_URL, params, SIGNED)
}

// PostContractCancelPlanOrder cancel-plan-order /** Cancel Plan Order (SIGNED)
func (cloudClient *CloudClient) PostContractCancelPlanOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_PLAN_ORDER_URL, params, SIGNED)
}

// PostContractTransfer transfer /** Transfer (SIGNED)
func (cloudClient *CloudClient) PostContractTransfer(currency string, amount string, transferType string, recvWindow int) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = currency
	params["amount"] = amount
	params["type"] = transferType
	if recvWindow > 0 {
		params["recvWindow"] = recvWindow
	}
	return cloudClient.requestWithParams(POST, API_CONTRACT_TRANSFER_URL, params, SIGNED)
}
