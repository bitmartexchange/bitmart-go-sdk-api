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
// Parameters:
// - symbol:  Symbol of the contract(like BTCUSDT)
// - orderId: Order ID
func (cloudClient *CloudClient) GetContractOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_URL, params, KEYED)
}

// GetContractOrderHistory order-history /** Get Order History (KEYED)
// Parameters:
// - symbol:  Symbol of the contract(like BTCUSDT)
// - Options.start_time: Start time, default is the last 7 days
// - Options.end_time: End time, default is the last 7 days
func (cloudClient *CloudClient) GetContractOrderHistory(contractSymbol string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_HISTORY_URL, params, KEYED)
}

// GetContractOpenOrders open orders /** Get All Open Orders (KEYED)
// Parameters:
// - Options.symbol: Symbol of the contract(like BTCUSDT)
// - Options.type: Order type -limit - market - trailing
// - Options.order_state: Order state -all(default) - partially_filled
// - Options.limit: The number of returned results, with a maximum of 100 and a default of 100
func (cloudClient *CloudClient) GetContractOpenOrders(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(GET, API_CONTRACT_OPEN_ORDERS_URL, params, KEYED)
}

// GetContractCurrentPlanOrders  /** Get All Current Plan Orders (KEYED)
// Parameters:
// - Options.symbol: Symbol of the contract(like BTCUSDT)
// - Options.type: Order type -limit - market
// - Options.limit: The number of returned results, with a maximum of 100 and a default of 100
func (cloudClient *CloudClient) GetContractCurrentPlanOrders(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(GET, API_CONTRACT_CURRENT_PLAN_ORDER_URL, params, KEYED)
}

// GetContractPosition position /** Get Current Position (KEYED)
// Parameters:
// - Options.symbol: Symbol of the contract(like BTCUSDT)
func (cloudClient *CloudClient) GetContractPosition(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(GET, API_CONTRACT_POSITION_URL, params, KEYED)
}

// GetContractPositionRisk /** Get Current Position Risk Details(KEYED)
// Parameters:
// - Options.symbol: Symbol of the contract(like BTCUSDT)
func (cloudClient *CloudClient) GetContractPositionRisk(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(GET, API_CONTRACT_POSITION_RISK_URL, params, KEYED)
}

// GetContractTrades trades /** Get Order Trade (KEYED)
// Parameters:
// - symbol: Symbol of the contract(like BTCUSDT)
// - Options.start_time: Start time, default is the last 7 days
// - Options.end_time: End time, default is the last 7 days
func (cloudClient *CloudClient) GetContractTrades(contractSymbol string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_TRADES_URL, params, KEYED)
}

// GetContractTransferList /** Get Transfer List (SIGNED)
// Parameters:
// - Options.currency: Currency (like USDT)
// - Options.time_start: Start time in milliseconds, (e.g. 1681701557927)
// - Options.time_end: End time in milliseconds, (e.g. 1681701557927)
// - Options.page: Number of pages, allowed range [1,1000]
// - Options.limit: Number of queries, allowed range [10,100]
// - Options.recvWindow: Trade time limit, allowed range (0,60000], default: 5000 milliseconds
func (cloudClient *CloudClient) GetContractTransferList(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(POST, API_CONTRACT_TRANSFER_LIST_URL, params, SIGNED)
}

// ContractOrder submit_contract order params
type ContractOrder struct {
	Symbol                    string `json:"symbol,omitempty"`
	ClientOrderId             string `json:"client_order_id,omitempty"`
	Type                      string `json:"type,omitempty"`
	Side                      int    `json:"side,omitempty"`
	Leverage                  string `json:"leverage,omitempty"`
	OpenType                  string `json:"open_type,omitempty"`
	Mode                      int    `json:"mode,omitempty"`
	Price                     string `json:"price,omitempty"`
	Size                      int    `json:"size,omitempty"`
	ActivationPrice           string `json:"activation_price,omitempty"`
	CallbackRate              string `json:"callback_rate,omitempty"`
	ActivationPriceType       int    `json:"activation_price_type,omitempty"`
	PresetTakeProfitPriceType int    `json:"preset_take_profit_price_type,omitempty"`
	PresetStopLossPriceType   int    `json:"preset_stop_loss_price_type,omitempty"`
	PresetTakeProfitPrice     string `json:"preset_take_profit_price,omitempty"`
	PresetStopLossPrice       string `json:"preset_stop_loss_price,omitempty"`
}

// PostContractSubmitOrder submit-order /** Submit Order (SIGNED)
func (cloudClient *CloudClient) PostContractSubmitOrder(order ContractOrder) (*CloudResponse, error) {
	params := NewParams()
	AddToParams("symbol", order.Symbol, params)
	AddToParams("client_order_id", order.ClientOrderId, params)
	AddToParams("type", order.Type, params)
	AddToParams("side", order.Side, params)
	AddToParams("leverage", order.Leverage, params)
	AddToParams("open_type", order.OpenType, params)
	AddToParams("mode", order.Mode, params)
	AddToParams("price", order.Price, params)
	AddToParams("size", order.Size, params)
	AddToParams("activation_price", order.ActivationPrice, params)
	AddToParams("callback_rate", order.CallbackRate, params)
	AddToParams("activation_price_type", order.ActivationPriceType, params)
	AddToParams("preset_take_profit_price_type", order.PresetTakeProfitPriceType, params)
	AddToParams("preset_stop_loss_price_type", order.PresetStopLossPriceType, params)
	AddToParams("preset_take_profit_price", order.PresetTakeProfitPrice, params)
	AddToParams("preset_stop_loss_price", order.PresetStopLossPrice, params)

	return cloudClient.requestWithParams(POST, API_CONTRACT_SUBMIT_ORDER_URL, params, SIGNED)
}

// PostContractCancelOrder cancel-order /** Cancel Order (SIGNED)
// Parameters:
// - symbol: Symbol of the contract(like BTCUSDT)
// - orderId: Order ID
func (cloudClient *CloudClient) PostContractCancelOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDER_URL, params, SIGNED)
}

// PostContractCancelOrders cancel-orders /** Cancel All Orders (SIGNED)
// Parameters:
// - symbol: Symbol of the contract(like BTCUSDT)
func (cloudClient *CloudClient) PostContractCancelOrders(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDERS_URL, params, SIGNED)
}

type ContractPlanOrder struct {
	Symbol                    string `json:"symbol"`
	Type                      string `json:"type,omitempty"`
	Side                      int    `json:"side"`
	Leverage                  string `json:"leverage"`
	OpenType                  string `json:"open_type"`
	Mode                      int    `json:"mode,omitempty"`
	Size                      int    `json:"size"`
	TriggerPrice              string `json:"trigger_price"`
	ExecutivePrice            string `json:"executive_price"`
	PriceWay                  int    `json:"price_way"`
	PriceType                 int    `json:"price_type"`
	PlanCategory              int    `json:"plan_category"`
	PresetTakeProfitPriceType int    `json:"preset_take_profit_price_type"`
	PresetStopLossPriceType   int    `json:"preset_stop_loss_price_type"`
	PresetTakeProfitPrice     string `json:"preset_take_profit_price"`
	PresetStopLossPrice       string `json:"preset_stop_loss_price"`
}

// PostContractPlanOrder plan-order /** Submit Plan Order (SIGNED)
func (cloudClient *CloudClient) PostContractPlanOrder(planOrder ContractPlanOrder) (*CloudResponse, error) {
	params := NewParams()
	AddToParams("symbol", planOrder.Symbol, params)
	AddToParams("type", planOrder.Type, params)
	AddToParams("side", planOrder.Side, params)
	AddToParams("leverage", planOrder.Leverage, params)
	AddToParams("open_type", planOrder.OpenType, params)
	AddToParams("mode", planOrder.Mode, params)
	AddToParams("size", planOrder.Size, params)
	AddToParams("trigger_price", planOrder.TriggerPrice, params)
	AddToParams("executive_price", planOrder.ExecutivePrice, params)
	AddToParams("price_way", planOrder.PriceWay, params)
	AddToParams("price_type", planOrder.PriceType, params)
	AddToParams("plan_category", planOrder.PlanCategory, params)
	AddToParams("preset_take_profit_price_type", planOrder.PresetTakeProfitPriceType, params)
	AddToParams("preset_stop_loss_price_type", planOrder.PresetStopLossPriceType, params)
	AddToParams("preset_take_profit_price", planOrder.PresetTakeProfitPrice, params)
	AddToParams("preset_stop_loss_price", planOrder.PresetStopLossPrice, params)
	return cloudClient.requestWithParams(POST, API_CONTRACT_SUBMIT_PLAN_ORDER_URL, params, SIGNED)
}

// PostContractCancelPlanOrder cancel-plan-order /** Cancel Plan Order (SIGNED)
// Parameters:
// - symbol: Symbol of the contract(like BTCUSDT)
// - orderId: Order ID
func (cloudClient *CloudClient) PostContractCancelPlanOrder(contractSymbol string, orderId string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["order_id"] = orderId
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_PLAN_ORDER_URL, params, SIGNED)
}

// PostContractTransfer transfer /** Transfer (SIGNED)
// Parameters:
// - currency: Currency
// - amount: Transfer amount，allowed range[0.01,10000000000]
// - type: Transfer type  -spot_to_contract  -contract_to_spot
// - recvWindow: Trade time limit, allowed range (0,60000], default: 5000 milliseconds
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

// PostContractSubmitLeverage submit-leverage /** Submit Leverage (SIGNED)
// Parameters:
// - symbol: Symbol of the contract(like BTCUSDT)
// - leverage: Order leverage
// - openType: Open type, required at close position  -cross  -isolated
func (cloudClient *CloudClient) PostContractSubmitLeverage(contractSymbol string, leverage string, openType string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = contractSymbol
	params["open_type"] = openType
	if leverage != "" {
		params["leverage"] = leverage
	}
	return cloudClient.requestWithParams(POST, API_CONTRACT_SUBMIT_LEVERAGE_URL, params, SIGNED)
}
