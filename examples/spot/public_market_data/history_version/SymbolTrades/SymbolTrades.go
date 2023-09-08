package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/symbols/trades
	Doc: https://developer-pro.bitmart.com/en/spot/#get-recent-trades
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Recent Trades
	var ac, err = client.GetSpotSymbolTrade("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
