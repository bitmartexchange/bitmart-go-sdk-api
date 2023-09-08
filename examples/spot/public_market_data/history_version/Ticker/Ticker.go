package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v2/ticker
	Doc: https://developer-pro.bitmart.com/en/spot/#get-ticker-of-all-pairs-v2
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Ticker of All Pairs (V2)
	var ac, err = client.GetSpotTicker()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
