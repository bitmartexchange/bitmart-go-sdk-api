//go:build api
// +build api

package bitmart

import (
	"log"
	"testing"
	"time"
)

// GET https://api-cloud.bitmart.com/contract/v1/tickers
func TestGetContractTickers(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractTickers("ETHUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/public/details
func TestGetContractDetails(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractDetails("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/public/depth
func TestGetContractDepth(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractDepth("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/public/open-interest
func TestGetContractOpenInterest(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOpenInterest("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/public/funding-rate
func TestGetContractFundingRate(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractFundingRate("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/public/kline
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

// GET https://api-cloud.bitmart.com/contract/private/assets-detail
func TestGetContractAssetsDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractAssetsDetail()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/private/order
func TestGetContractOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOrder("BTCUSDT", "220609666322019")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/private/order-history
func TestGetContractOrderHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractOrderHistory("BTCUSDT", 1662518172, 1662518172)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/private/position
func TestGetContractPosition(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractPosition("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/contract/private/trades
func TestGetContractTrades(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetContractTrades("BTCUSDT", 1662518172, 1662518172)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/contract/private/submit-order
func TestPostContractSubmitOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractSubmitOrder(ContractOrder{Symbol: "ETHUSDT", Side: 4, Type: "limit", Leverage: "1", OpenType: "isolated", Size: 10, Price: "2000"})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/contract/private/cancel-order
func TestPostContractCancelOrder(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractCancelOrder("ETHUSDT", "220906179559421")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/contract/private/cancel-orders
func TestPostContractCancelOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostContractCancelOrders("ETHUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// go test -run Contract 带Contract的测试案例
// go test -run TestGetContractAssetsDetail
// cloud_consts.go中API_URL_PRO已改为dev域名
