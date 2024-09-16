package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/quotation/v3/ticker
Doc: https://developer-pro.bitmart.com/en/spot/#get-ticker-of-a-trading-pair-v3
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Ticker of a Trading Pair (V3)
	var ac, err = client.GetSpotV3Ticker("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
