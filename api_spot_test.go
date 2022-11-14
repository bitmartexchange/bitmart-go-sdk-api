package bitmart

import (
	"log"
	"testing"
	"time"
)

const TEST_SYMBOL = "BTC_USDT"

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

// GET https://api-cloud.bitmart.com/spot/v2/ticker
func TestGetSpotTicker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotTicker()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/ticker_detail
func TestGetSpotTickerDetail(t *testing.T) {
	c := NewTestClient()
	ab, err := c.GetSpotTickerDetail(TEST_SYMBOL)
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

//POST https://api-cloud.bitmart.com/spot/v2/submit_order
func TestPostSpotSubmitOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubmitOrder(Order{Symbol: TEST_SYMBOL, Side: "buy", Type: "limit", ClientOrderId: "", Size: "0.1", Price: "8800", Notional: ""})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

//POST https://api-cloud.bitmart.com/spot/v1/margin/submit_order
func TestPostMarginSubmitOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostMarginSubmitOrder(MarginOrder{Symbol: TEST_SYMBOL, Side: "buy", Type: "limit", ClientOrderId: "", Size: "0.1", Price: "8800", Notional: ""})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

//POST https://api-cloud.bitmart.com/spot/v2/batch_orders
func TestPostSpotBatchOrders(t *testing.T) {
	c := NewTestClient()
	var orderParams [1]Order
	orderParams[0] = Order{Symbol: TEST_SYMBOL, Side: "buy", Type: "limit", ClientOrderId: "", Size: "0.1", Price: "8800", Notional: ""}
	ac, err := c.PostSpotBatchOrders(orderParams)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v3/cancel_order
func TestPostSpotCancelOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelOrder(TEST_SYMBOL, "2147601610", "") // cancel by order_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.PostSpotCancelOrder(TEST_SYMBOL, "", "myid872923") // cancel by client_order_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
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

// GET https://api-cloud.bitmart.com/spot/v2/order_detail
func TestGetSpotOrderDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderDetail("2147601610")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v3/orders
func TestGetSpotOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrders(TEST_SYMBOL, "", 100, "10")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestGetSpotOrdersByTime(t *testing.T) {
	c := NewTestClient()
	now := time.Now().UnixMilli()
	ac, err := c.GetSpotOrdersByTime(TEST_SYMBOL, "", 100, "10",
		now-1000*60*60*24*7, now)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v2/trades
func TestGetSpotHistoryTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotHistoryTrades(TEST_SYMBOL, "", 10)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestGetSpotHistoryTradesByTime(t *testing.T) {
	c := NewTestClient()
	now := time.Now().UnixMilli()
	ac, err := c.GetSpotHistoryTradesByTime(TEST_SYMBOL, "", 10,
		now-1000*60*60*24*7, now)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestGetSpotOrderTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderTrades(TEST_SYMBOL, "", "2147601596")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
