package bitmart

import (
	"log"
	"testing"
	"time"
)

// GET https://api-cloud-v2.bitmart.com/contract/public/details
func TestGetContractDetails(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractDetails("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/public/depth
func TestGetContractDepth(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractDepth("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/public/open-interest
func TestGetContractOpenInterest(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOpenInterest("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/public/funding-rate
func TestGetContractFundingRate(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractFundingRate("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/public/kline
func TestGetContractKline(t *testing.T) {
	c := NewTestClient()
	now := time.Now().Unix()
	ac, err := c.GetContractKline("BTCUSDT", int(now-3600*24), int(now), 5)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/assets-detail
func TestGetContractAssetsDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractAssetsDetail()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/order
func TestGetContractOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOrder("BTCUSDT", "220609666322019")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/order-history
func TestGetContractOrderHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOrderHistory("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	now := time.Now().Unix()
	ac2, err2 := c.GetContractOrderHistory("BTCUSDT", map[string]interface{}{
		"start_time": int(now - 3600),
		"end_time":   int(now),
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/get-open-orders
func TestGetContractOpenOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOpenOrders()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.GetContractOpenOrders(map[string]interface{}{
		"symbol":      "BTCUSDT",
		"type":        "limit",
		"order_state": "all",
		"limit":       10,
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/current-plan-order
func TestGetContractCurrentPlanOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractCurrentPlanOrders(map[string]interface{}{
		"symbol":    "BTCUSDT",
		"type":      "limit",
		"limit":     10,
		"plan_type": "plan",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/position
func TestGetContractPosition(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractPosition()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.GetContractPosition(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/position-risk
func TestGetContractPositionRisk(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractPositionRisk()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.GetContractPositionRisk(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud-v2.bitmart.com/contract/private/trades
func TestGetContractTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractTrades("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	now := time.Now().Unix()
	ac2, err2 := c.GetContractTrades("BTCUSDT", map[string]interface{}{
		"start_time": int(now - 3600),
		"end_time":   int(now),
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// POST https://api-cloud-v2.bitmart.com/account/v1/transfer-contract-list
func TestGetContractTransferList(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractTransferList(map[string]interface{}{
		"page":     1,
		"limit":    10,
		"currency": "USDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/submit-order
func TestPostContractSubmitOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractSubmitOrder(ContractOrder{
		Symbol:   "ETHUSDT",
		Side:     4,
		Type:     "limit",
		Leverage: "1",
		OpenType: "isolated",
		Size:     10,
		Price:    "2000",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/cancel-order
func TestPostContractCancelOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractCancelOrder("ETHUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.PostContractCancelOrder("ETHUSDT", map[string]interface{}{
		"order_id": "12321313",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}

	ac3, err3 := c.PostContractCancelOrder("ETHUSDT", map[string]interface{}{
		"client_order_id": "12321313",
	})
	if err3 != nil {
		log.Panic(err3)
	} else {
		PrintResponse(ac3)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/cancel-orders
func TestPostContractCancelOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractCancelOrders("ETHUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/submit-plan-order
func TestPostContractPlanOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractPlanOrder(ContractPlanOrder{
		Symbol:         "ETHUSDT",
		Side:           4,
		Type:           "limit",
		Leverage:       "1",
		OpenType:       "isolated",
		Size:           10,
		Mode:           1,
		TriggerPrice:   "2000",
		ExecutivePrice: "1800",
		PriceWay:       1,
		PriceType:      1,
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/cancel-plan-order
func TestPostContractCancelPlanOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractCancelPlanOrder("ETHUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.PostContractCancelPlanOrder("ETHUSDT", map[string]interface{}{
		"order_id": "220906179559421",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}

}

// POST https://api-cloud-v2.bitmart.com/account/v1/transfer-contract
func TestPostContractTransfer(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractTransfer("USDT", "100", "spot_to_contract", 5000)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/submit-leverage
func TestPostContractSubmitLeverage(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractSubmitLeverage("ETHUSDT", "10", "cross")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/submit-tp-sl-order
func TestPostContractSubmitTpSlOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractSubmitTpSlOrder(
		"ETHUSDT",
		"take_profit",
		2,
		"2000",
		"1450",
		1,
		map[string]interface{}{
			"plan_category":   1,
			"client_order_id": "123123123123",
			"category":        "limit",
		},
	)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/modify-plan-order
func TestPostContractModifyPlanOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractModifyPlanOrder(
		"ETHUSDT",
		"limit",
		"2000",
		1,
		map[string]interface{}{
			"client_order_id": "123123123123",
			"executive_price": "1450",
		},
	)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/modify-preset-plan-order
func TestPostContractModifyPresetPlanOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractModifyPresetPlanOrder(
		"ETHUSDT",
		"220609666322019",
		map[string]interface{}{
			"preset_take_profit_price":      "2000",
			"preset_stop_loss_price":        "1900",
			"preset_take_profit_price_type": 1,
			"preset_stop_loss_price_type":   1,
		},
	)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/contract/private/modify-tp-sl-order
func TestPostContractModifyTpSlOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractModifyTpSlOrder(
		"ETHUSDT",
		"2100",
		2,
		map[string]interface{}{
			"executive_price": "2100",
			"order_id":        "q232434234234",
			"plan_category":   2,
			"category":        "limit",
		},
	)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
