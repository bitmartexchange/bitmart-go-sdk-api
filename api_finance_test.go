package bitmart

import (
	"log"
	"testing"
)

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/earn
func TestGetFinanceSavingsHoldings(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceSavingsHoldings()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/product
func TestGetFinanceFlexibleProducts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFlexibleProducts(1, 10, map[string]interface{}{"coinName": "USDT"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/subscribe
func TestPostFinanceFlexibleSubscribe(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceFlexibleSubscribe("1001", "100", "20000009000000300000")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/redeem
func TestPostFinanceFlexibleRedeem(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceFlexibleRedeem("200001", "50", "20000009000000300001")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/earn
func TestGetFinanceFlexibleHoldings(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFlexibleHoldings(1, 10, map[string]interface{}{"coinName": "USDT"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/record
func TestGetFinanceFlexibleRecords(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFlexibleRecords("subscribe", 1, 10)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/fixed/product
func TestGetFinanceFixedProducts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFixedProducts(1, 10, map[string]interface{}{"coinName": "USDT"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/fixed/subscribe
func TestPostFinanceFixedSubscribe(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceFixedSubscribe("1001", "10", "20000009000000300002", "OFF")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/fixed/earn
func TestGetFinanceFixedHoldings(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFixedHoldings(1, 10, map[string]interface{}{"coinName": "USDT"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/fixed/record
func TestGetFinanceFixedRecords(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFixedRecords("subscribe", 1, 10)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/fixed/redeem
func TestPostFinanceFixedRedeem(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceFixedRedeem("200001", "20000009000000300003")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/fixed/subscribe/operate
func TestPostFinanceFixedModifyAutoRenewal(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceFixedModifyAutoRenewal("200001", "OFF")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/subscribe/batch/operate
func TestPostFinanceAutoSavingToggle(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceAutoSavingToggle("open")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/subscribe/batch
func TestGetFinanceAutoSavingStatus(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceAutoSavingStatus()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/newearn/cloud/v1/saving/subscribe/operate
func TestPostFinanceFlexibleAutoSubscribeToggle(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostFinanceFlexibleAutoSubscribeToggle("1001", "open")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/newearn/cloud/v1/saving/subscribe/status
func TestGetFinanceFlexibleAutoSubscribeStatus(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetFinanceFlexibleAutoSubscribeStatus("1001")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
