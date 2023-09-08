package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/symbols
	Doc: https://developer-pro.bitmart.com/en/spot/#get-list-of-trading-pairs
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get List of Trading Pairs
	var ac, err = client.GetSpotSymbol()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
