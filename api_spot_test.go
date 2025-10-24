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

// GET https://api-cloud.bitmart.com/spot/quotation/v3/tickers
func TestGetSpotV3Tickers(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotV3Tickers()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/quotation/v3/ticker
func TestGetSpotV3Ticker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotV3Ticker("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/quotation/v3/lite-klines
func TestGetSpotV3LatestKline(t *testing.T) {
	c := NewTestClient()
	before := time.Now().Unix()
	after := before - 60*60

	ac, err := c.GetSpotV3LatestKline("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	var ac2, err2 = c.GetSpotV3LatestKline("BTC_USDT", map[string]interface{}{
		"before": before,
		"after":  after,
		"step":   15,
		"limit":  10,
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud.bitmart.com/spot/quotation/v3/klines
func TestGetSpotV3HistoryKline(t *testing.T) {
	c := NewTestClient()
	before := time.Now().Unix()
	after := before - 60*60
	ac, err := c.GetSpotV3HistoryKline("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	var ac2, err2 = c.GetSpotV3HistoryKline("BTC_USDT", map[string]interface{}{
		"before": before,
		"after":  after,
		"step":   15,
		"limit":  10,
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud.bitmart.com/spot/quotation/v3/books
func TestGGetSpotV3Book(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotV3Book("BTC_USDT") // find by default
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	var ac2, err2 = c.GetSpotV3Book("BTC_USDT", map[string]interface{}{
		"limit": 10,
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}

}

// GET https://api-cloud.bitmart.com/spot/quotation/v3/trades
func TestGetSpotV3Trade(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotV3Trade("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.GetSpotV3Trade("BTC_USDT", map[string]interface{}{
		"limit": 10,
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
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

// POST https://api-cloud.bitmart.com/spot/v2/submit_order
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
		StpMode:       "cancel_maker",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v1/margin/submit_order
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

// POST https://api-cloud.bitmart.com/spot/v4/batch_orders
func TestPostSpotBatchOrders(t *testing.T) {
	c := NewTestClient()
	var orderParams []BatchOrder
	orderParams = append(orderParams, BatchOrder{
		Side:          "sell",
		Type:          "limit",
		ClientOrderId: "pix12312031231zz",
		Size:          "0.1",
		Price:         "880000",
	})

	orderParams = append(orderParams, BatchOrder{
		Side:          "sell",
		Type:          "market",
		ClientOrderId: "pix123120312312nn",
		Size:          "0.1",
		Notional:      "880000",
		StpMode:       "cancel_maker",
	})
	ac, err := c.PostSpotBatchOrders("BTC_USDT", orderParams[:])
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

}

// POST https://api-cloud.bitmart.com/spot/v3/cancel_order
func TestPostSpotCancelOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelOrder("BTC_USDT", map[string]interface{}{
		"order_id": "12321321312312",
	}) // cancel by order_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.PostSpotCancelOrder("BTC_USDT", map[string]interface{}{
		"client_order_id": "xdfsfsdfsd",
	}) // cancel by client_order_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/cancel_orders
func TestPostSpotCancelOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelOrders("BTC_USDT", map[string]interface{}{
		"orderIds": []string{"12312312", "12312312312"},
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.PostSpotCancelOrders("BTC_USDT", map[string]interface{}{
		"clientOrderIds": []string{"12312312", "12312312312"},
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// POST https://api-cloud.bitmart.com/spot/v4/cancel_orders
func TestPostSpotCancelAllOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotCancelAllOrder()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.PostSpotCancelAllOrder(map[string]interface{}{
		"symbol": "BTC_USDT",
		"side":   "buy",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
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
