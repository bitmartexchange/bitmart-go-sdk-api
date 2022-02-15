package bitmart

import (
	"log"
	"testing"
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
