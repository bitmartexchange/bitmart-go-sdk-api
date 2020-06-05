package bitmart

const EXCHANGE  = "bitmart"

// contracts
func (cloudClient *CloudClient) getContractAllContracts() (*CloudResponse, error) {
	params := NewParams()
	params["exchange"] = EXCHANGE
	return cloudClient.requestWithParams(GET, API_CONTRACT_CURRENCIES_URL, params, NONE)
}

func (cloudClient *CloudClient) getContractContracts(contractId int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = IntToString(contractId)
	params["exchange"] = EXCHANGE
	return cloudClient.requestWithParams(GET, API_CONTRACT_CURRENCIES_URL, params, NONE)
}


// pnls
func (cloudClient *CloudClient) getContractPnls(contractId int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	return cloudClient.requestWithParams(GET, API_CONTRACT_PNLS_URL, params, NONE)
}


// indexes
func (cloudClient *CloudClient) getContractIndexes() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_CONTRACT_INDEXES_URL, NONE)
}


// tickers
func (cloudClient *CloudClient) getContractTickers() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_CONTRACT_TICKERS_URL, NONE)
}

func (cloudClient *CloudClient) getContractTickersByContractId(contractId int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	return cloudClient.requestWithParams(GET, API_CONTRACT_TICKERS_URL, params, NONE)
}


// quote
func (cloudClient *CloudClient) getContractQuote(contractId int, startTime int, endTime int, unit int, resolution string) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["unit"] = unit
	params["resolution"] = resolution
	return cloudClient.requestWithParams(GET, API_CONTRACT_QUOTE_URL, params, NONE)
}


// index quote
func (cloudClient *CloudClient) getContractIndexQuote(indexId int, startTime int, endTime int, unit int, resolution string) (*CloudResponse, error) {
	params := NewParams()
	params["indexID"] = indexId
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["unit"] = unit
	params["resolution"] = resolution
	return cloudClient.requestWithParams(GET, API_CONTRACT_INDEX_QUOTE_URL, params, NONE)
}

// trades
func (cloudClient *CloudClient) getContractTrades(contractId int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	return cloudClient.requestWithParams(GET, API_CONTRACT_TRADES_URL, params, NONE)
}

// depth
func (cloudClient *CloudClient) getContractDepth(contractId int, count int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	if count != 0 {
		params["count"] = count
	}
	return cloudClient.requestWithParams(GET, API_CONTRACT_DEPTH_URL, params, NONE)
}

// funding rate
func (cloudClient *CloudClient) getContractFundingRate(contractId int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	return cloudClient.requestWithParams(GET, API_CONTRACT_FUNDING_RATE_URL, params, NONE)
}

// --------------------- Trading API
// userOrders
func (cloudClient *CloudClient) getContractUserOrders(contractId int, status int, offset int, size int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	params["status"] = status
	if offset != 0 && size != 0 {
		params["offset"] = offset
		params["size"] = size
	}
	return cloudClient.requestWithParams(GET, API_CONTRACT_USER_ORDERS_URL, params, KEYED)
}

// userOrderInfo
func (cloudClient *CloudClient) getContractUserOrderInfo(contractId int, orderId int64) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	params["orderID"] = orderId
	return cloudClient.requestWithParams(GET, API_CONTRACT_USER_ORDER_INFO_URL, params, KEYED)
}

// userTrades
func (cloudClient *CloudClient) getContractUserTrades(contractId int, offset int, size int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	if offset == 0 && size != 0 {
		params["offset"] = offset
		params["size"] = size
	}

	return cloudClient.requestWithParams(GET, API_CONTRACT_USER_TRADES_URL, params, KEYED)
}

// orderTrades
func (cloudClient *CloudClient) getContractOrderTrades(contractId int, orderId int64) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	params["orderID"] = orderId
	return cloudClient.requestWithParams(GET, API_CONTRACT_ORDER_TRADES_URL, params, KEYED)
}

type CreateOrder struct {
	ContractId int    `json:"contract_id"`
	Category   int    `json:"category"`
	Way        int    `json:"way"`
	CustomId   int    `json:"custom_id"`
	OpenType   int    `json:"open_type"`
	Leverage   int    `json:"leverage"`
	Price      string `json:"price"`
	Vol        string `json:"vol"`
}

// batchOrders
func (cloudClient *CloudClient) postContractBatchOrders(orders []CreateOrder) (*CloudResponse, error) {
	params := NewParams()
	params["orders"] = orders
	return cloudClient.requestWithParams(POST, API_CONTRACT_USER_BATCH_ORDERS_URL, params, SIGNED)
}

// submitOrder
func (cloudClient *CloudClient) postContractSubmitOrders(order CreateOrder) (*CloudResponse, error) {
	params := NewParams()
	params["contract_id"] = order.ContractId
	params["category"] = order.Category
	params["way"] = order.Way
	params["custom_id"] = order.CustomId
	params["open_type"] = order.OpenType
	params["leverage"] = order.Leverage
	params["price"] = order.Price
	params["vol"] = order.Vol
	return cloudClient.requestWithParams(POST, API_CONTRACT_USER_SUBMIT_ORDER_URL, params, SIGNED)
}

type CancelOrder struct {
	ContractId int     `json:"contract_id"`
	Orders     []int64 `json:"orders"`
}

// cancelOrders
func (cloudClient *CloudClient) postContractCancelOrders(cancelOrders []CancelOrder) (*CloudResponse, error) {
	params := NewParams()
	params["orders"] = cancelOrders
	return cloudClient.requestWithParams(POST, API_CONTRACT_CANCEL_ORDERS_URL, params, SIGNED)
}


// accounts
func (cloudClient *CloudClient) getContractAccounts(coinCode string) (*CloudResponse, error) {
	params := NewParams()
	params["coinCode"] = coinCode
	return cloudClient.requestWithParams(GET, API_CONTRACT_ACCOUNTS_URL, params, KEYED)
}


// userPositions
func (cloudClient *CloudClient) getContractUserPositions(contractId int) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	return cloudClient.requestWithParams(GET, API_CONTRACT_USER_POSITIONS_URL, params, KEYED)
}

// userLiqRecords
func (cloudClient *CloudClient) getContractUserLiqRecords(contractId int, orderId int64) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	if orderId != 0 {
		params["orderID"] = orderId
	}
	return cloudClient.requestWithParams(GET, API_CONTRACT_USER_LIQ_RECORDS_URL, params, KEYED)
}

// positionFee
func (cloudClient *CloudClient) getContractPositionFee(contractId int, positionId int64) (*CloudResponse, error) {
	params := NewParams()
	params["contractID"] = contractId
	params["positionID"] = positionId
	return cloudClient.requestWithParams(GET, API_CONTRACT_POSITION_FEE_URL, params, KEYED)
}

// marginOper
func (cloudClient *CloudClient) postContractMarginOper(contractId int, positionId int64, vol int, operType int) (*CloudResponse, error) {
	params := NewParams()
	params["contract_id"] = contractId
	params["position_id"] = positionId
	params["vol"] = vol
	params["oper_type"] = operType
	return cloudClient.requestWithParams(POST, API_CONTRACT_MARGIN_OPER_URL, params, SIGNED)
}