package bitmart

import (
	"log"
	"testing"
	"time"
)

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
	ab, err := c.GetSpotTickerDetail("BTC_USDT")
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
	to := time.Now().Unix()
	from := to - 60*60
	ac, err := c.GetSpotSymbolKline("BTC_USDT", from, to, 15)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/book
func TestGetSpotSymbolBook(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbolBook("BTC_USDT", 2, 10) // find by default
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetSpotSymbolBook("BTC_USDT", 2, 5) // find by precision is 8
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/symbols/trades
func TestGetSpotSymbolTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSymbolTrade("BTC_USDT")
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
	ac, err := c.PostSpotSubmitOrder(Order{
		Symbol:        "BTC_USDT",
		Side:          "buy",
		Type:          "limit",
		ClientOrderId: "",
		Size:          "0.1",
		Price:         "8800",
		Notional:      "",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

//POST https://api-cloud.bitmart.com/spot/v1/margin/submit_order
func TestPostMarginSubmitOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostMarginSubmitOrder(MarginOrder{
		Symbol:        "BTC_USDT",
		Side:          "buy",
		Type:          "limit",
		ClientOrderId: "",
		Size:          "0.1",
		Price:         "8800",
		Notional:      "",
	})
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
	orderParams[0] = Order{Symbol: "BTC_USDT", Side: "buy", Type: "limit", ClientOrderId: "", Size: "0.1", Price: "8800", Notional: ""}
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
	ac, err := c.PostSpotCancelOrder("BTC_USDT", "2147601610", "") // cancel by order_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.PostSpotCancelOrder("BTC_USDT", "", "myid872923") // cancel by client_order_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// POST https://api-cloud.bitmart.com/spot/v1/cancel_orders
func TestPostSpotCancelOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelOrders("BTC_USDT", "sell")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/query/order
func TestGetSpotOrderByOrderId(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderByOrderId("118100034543076010", "", 0)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/query/client-order
func TestGetSpotOrderByClientOrderId(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderByClientOrderId("118100034543076010", "", 0)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/query/open-orders
func TestGetSpotOpenOrders(t *testing.T) {
	c := NewTestClient()
	now := time.Now().UnixNano() / int64(time.Millisecond)
	ac, err := c.GetSpotOpenOrders("BTC_USDT", "spot", 0, now, 10, 5000)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/query/history-orders
func TestGetSpotAccountOrders(t *testing.T) {
	c := NewTestClient()
	now := time.Now().UnixNano() / int64(time.Millisecond)
	ac, err := c.GetSpotAccountOrders("BTC_USDT", "spot", 0, now, 10, 5000)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/query/trades
func TestGetSpotAccountTradeList(t *testing.T) {
	c := NewTestClient()
	now := time.Now().UnixNano() / int64(time.Millisecond)
	ac, err := c.GetSpotAccountTradeList("BTC_USDT", "spot", 0, now, 10, 5000)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/query/order-trades
func TestGetSpotOrderTradeList(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrderTradeList("118100034543076010", 5000)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
