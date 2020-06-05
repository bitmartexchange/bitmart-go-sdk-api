package bitmart

import (
	"log"
	"testing"
)

const TEST_SYMBOL  = "BTC_USDT"

// GET https://api-cloud.bitmart.com/spot/v1/currencies
func TestGetSpotCurrencies(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotCurrencies()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols
func TestGetSpotSymbol(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotSymbol()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/details
func TestGetSpotSymbolDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotSymbolDetail()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/ticker
func TestGetSpotTicker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotTicker("") // find all
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}

	ab, err := c.getSpotTicker(TEST_SYMBOL) // find by symbol
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/steps
func TestGetSpotSteps(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotSteps()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/kline
func TestGetSpotSymbolKline(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotSymbolKline(TEST_SYMBOL, 1525760116000, 1525769116000, 15)
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/book
func TestGetSpotSymbolBook(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotSymbolBook(TEST_SYMBOL, 0) // find by default
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}

	ab, err := c.getSpotSymbolBook(TEST_SYMBOL, 8) // find by precision is 8
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/trades
func TestGetSpotSymbolTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotSymbolTrade(TEST_SYMBOL)
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/wallet
func TestGetSpotWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotWallet()
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

//POST https://api-cloud.bitmart.com/spot/v1/submit_order
func TestPostSpotSubmitLimitBuyOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.postSpotSubmitLimitBuyOrder(LimitBuyOrder{Symbol:TEST_SYMBOL, Size:"8800", Price:"0.01"})
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

func TestPostSpotSubmitLimitSellOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.postSpotSubmitLimitSellOrder(LimitSellOrder{Symbol:TEST_SYMBOL, Size:"10050", Price:"1"})
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

func TestPostSpotSubmitMarketBuyOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.postSpotSubmitMarketBuyOrder(MarketBuyOrder{Symbol:TEST_SYMBOL, Notional:"1000"})
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}


func TestPostSpotSubmitMarketSellOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.postSpotSubmitMarketSellOrder(MarketSellOrder{Symbol:TEST_SYMBOL, Size:"1"})
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}



// POST https://api-cloud.bitmart.com/spot/v1/cancel_order
func TestPostSpotCancelOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.postSpotCancelOrder(TEST_SYMBOL, 2147601610)
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v1/cancel_orders
func TestPostSpotCancelOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.postSpotCancelOrders(TEST_SYMBOL, "sell")
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/order_detail
func TestGetSpotOrderDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotOrderDetail(TEST_SYMBOL, 2147601610)
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/orders
func TestGetSpotOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotOrders(TEST_SYMBOL, 1, 10, "10")
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}


// GET https://api-cloud.bitmart.com/spot/v1/trades
func TestGetSpotHistoryTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotHistoryTrades(TEST_SYMBOL, 1, 2)
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}

func TestGetSpotOrderTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.getSpotOrderTrades(TEST_SYMBOL, 2147601596)
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}
}