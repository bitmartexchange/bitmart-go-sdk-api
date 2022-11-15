package bitmart

import (
	"log"
	"testing"
	"time"
)

// GET https://api-cloud.bitmart.com/spot/v1/broker/rebate
func TestGetBrokerRebate(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetBrokerRebate()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

func TestGetBrokerRebateByTimestamp(t *testing.T) {
	c := NewTestClient()
	now := time.Now().Unix()
	ac, err := c.GetBrokerRebateByTimestamp(now-60*60*24*7, now)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
