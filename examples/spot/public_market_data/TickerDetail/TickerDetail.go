package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/ticker_detail
	Doc: https://developer-pro.bitmart.com/en/spot/#get-ticker-of-a-trading-pair
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Ticker of a Trading Pair
	var ac, err = client.GetSpotTickerDetail("BMX_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
