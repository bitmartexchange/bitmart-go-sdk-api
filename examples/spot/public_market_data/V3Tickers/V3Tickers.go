package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/quotation/v3/tickers
Doc: https://developer-pro.bitmart.com/en/spot/#get-ticker-of-all-pairs-v3
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Ticker of All Pairs (V3)
	var ac, err = client.GetSpotV3Tickers()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
