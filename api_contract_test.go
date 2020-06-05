package bitmart

import (
	"log"
	"testing"
)

const (
	CONTRACT_ID = 1
)

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/contracts
func TestGetContractContracts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractContracts(CONTRACT_ID)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/pnls
func TestGetContractPnl(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractPnls(CONTRACT_ID)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}


// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/indexes
func TestGetContractIndexes(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractIndexes()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/tickers
func TestGetContractTickers(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractTickers()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/quote
func TestGetContractQuote(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractQuote(CONTRACT_ID, 1584343602, 1585343602, 5, "M")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/indexquote
func TestGetContractIndexQuote(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractIndexQuote(1, 1584343602, 1585343602, 1, "H")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/trades
func TestGetContractTrade(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractTrades(CONTRACT_ID)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/depth
func TestGetContractDepth(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractDepth(CONTRACT_ID, 1)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/fundingrate
func TestGetContractFundingRate(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractFundingRate(CONTRACT_ID)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

//GET https://api-cloud.bitmart.com/contract/v1/ifcontract/userOrders
func TestGetContractUserOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractUserOrders(CONTRACT_ID, 0, 0, 0)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/userOrderInfo
func TestGetContractUserOrderInfo(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractUserOrderInfo(CONTRACT_ID, 2802614433)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/userTrades
func TestGetContractUserTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractUserTrades(CONTRACT_ID, 1, 10)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/orderTrades
func TestGetContractOrderTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOrderTrades(CONTRACT_ID, 2802614433)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/contract/v1/ifcontract/batchOrders
// {"orders":[{"contract_id":1,"category":1,"way":1,"custom_id":100,"open_type":1,"leverage":10,"price":"9000","vol":"1"},
// {"contract_id":1,"category":1,"way":1,"custom_id":100,"open_type":1,"leverage":10,"price":"9100","vol":"1"}]}
func TestPostContractBatchOrders(t *testing.T) {
	c := NewTestClient()
	var orderList = []CreateOrder{
		{ContractId: 1, Category: 1, Way: 1, CustomId: 100, OpenType: 1, Leverage: 10, Price: "9000", Vol: "1"},
		{ContractId: 1, Category: 1, Way: 1, CustomId: 100, OpenType: 1, Leverage: 10, Price: "9100", Vol: "1"},
	}
	ac, err := c.PostContractBatchOrders(orderList)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}


// POST https://api-cloud.bitmart.com/contract/v1/ifcontract/submitOrder
// {"message":"OK","code":1000,"trace":"8c85fa6f-fab8-4f12-89b2-84ba113ba9db","data":{"order_id":3826296406}}
func TestPostContractSubmitOrders(t *testing.T) {
	c := NewTestClient()
	var order = CreateOrder{ContractId: 1, Category: 1, Way: 1, CustomId: 10000, OpenType: 1, Leverage: 10, Price: "9000", Vol: "1"}
	ac, err := c.PostContractSubmitOrders(order)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/contract/v1/ifcontract/cancelOrders
// {"message":"OK","code":1000,"trace":"f48e7528-6db5-46ac-a5cd-4135405e2a5a","data":{"succeed":[3826296405,3826296406],"failed":null}}
func TestPostContractCancelOrders(t *testing.T) {
	c := NewTestClient()
	var orderIds = []CancelOrder{
		{ContractId: CONTRACT_ID, Orders: []int64{3826296405, 3826296406}},
	}
	ac, err := c.PostContractCancelOrders(orderIds)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/accounts
func TestGetContractAccounts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractAccounts("USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/userPositions
func TestGetContractUserPositions(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractUserPositions(CONTRACT_ID)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/userLiqRecords
func TestGetContractUserLiqRecords(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractUserLiqRecords(CONTRACT_ID, 3826296406)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/v1/ifcontract/positionFee
func TestGetContractPositionFee(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractPositionFee(CONTRACT_ID, 3826296406)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/contract/v1/ifcontract/marginOper
func TestPostContractMarginOper(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractMarginOper(CONTRACT_ID, 3826296406, 100, 1)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}