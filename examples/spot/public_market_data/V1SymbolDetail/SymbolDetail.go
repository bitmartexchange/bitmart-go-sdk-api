package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/symbols/details
	Doc: https://developer-pro.bitmart.com/en/spot/#get-list-of-trading-pair-details
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get List of Trading Pair Details
	var ac, err = client.GetSpotSymbolDetail()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}
}
