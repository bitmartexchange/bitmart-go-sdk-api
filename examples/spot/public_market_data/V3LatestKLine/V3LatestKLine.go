package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
GET /spot/quotation/v3/lite-klines
Doc: https://developer-pro.bitmart.com/en/spot/#get-latest-k-line-v3
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Latest K-Line (V3)
	var ac, err = client.GetSpotV3LatestKline("BTC_USDT")
	if err == nil {
		log.Println(ac.Response)
	}

	before := time.Now().Unix()
	after := before - 60*60
	var ac2, err2 = client.GetSpotV3LatestKline("BTC_USDT", map[string]interface{}{
		"before": before,
		"after":  after,
		"step":   15,
		"limit":  10,
	})
	if err2 == nil {
		log.Println(ac2.Response)
	}
}
