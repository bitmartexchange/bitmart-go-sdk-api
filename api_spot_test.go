package bitmart

import (
	"log"
	"testing"
)

const TEST_SYMBOL  = "BTC_USDT"

// GET https://api-cloud.bitmart.com/spot/v1/currencies
func TestGetSpotCurrencies(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotCurrencies()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols
func TestGetSpotSymbol(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbol()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/details
func TestGetSpotSymbolDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbolDetail()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/ticker
func TestGetSpotTicker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotTicker("") // find all
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetSpotTicker(TEST_SYMBOL) // find by symbol
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/steps
func TestGetSpotSteps(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSteps()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/kline
func TestGetSpotSymbolKline(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbolKline(TEST_SYMBOL, 1525760116, 1525769116, 15)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/book
func TestGetSpotSymbolBook(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbolBook(TEST_SYMBOL, 0, 0) // find by default
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetSpotSymbolBook(TEST_SYMBOL, 8, 5) // find by precision is 8
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/trades
func TestGetSpotSymbolTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbolTrade(TEST_SYMBOL)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/wallet
func TestGetSpotWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotWallet()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

//POST https://api-cloud.bitmart.com/spot/v1/submit_order
func TestPostSpotSubmitLimitBuyOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubmitLimitBuyOrder(LimitBuyOrder{Symbol:TEST_SYMBOL, Size:"8800", Price:"0.01"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestPostSpotSubmitLimitSellOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubmitLimitSellOrder(LimitSellOrder{Symbol:TEST_SYMBOL, Size:"10050", Price:"1"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestPostSpotSubmitMarketBuyOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubmitMarketBuyOrder(MarketBuyOrder{Symbol:TEST_SYMBOL, Notional:"1000"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}


func TestPostSpotSubmitMarketSellOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubmitMarketSellOrder(MarketSellOrder{Symbol:TEST_SYMBOL, Size:"1"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}



// POST https://api-cloud.bitmart.com/spot/v2/cancel_order
func TestPostSpotCancelOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelOrder(TEST_SYMBOL, 2147601610)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v1/cancel_orders
func TestPostSpotCancelOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelOrders(TEST_SYMBOL, "sell")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/order_detail
func TestGetSpotOrderDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderDetail(TEST_SYMBOL, 2147601610)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/orders
func TestGetSpotOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrders(TEST_SYMBOL, 1, 10, "10")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}


// GET https://api-cloud.bitmart.com/spot/v1/trades
func TestGetSpotHistoryTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotHistoryTrades(TEST_SYMBOL, 1, 2)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestGetSpotOrderTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderTrades(TEST_SYMBOL, 2147601596)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}