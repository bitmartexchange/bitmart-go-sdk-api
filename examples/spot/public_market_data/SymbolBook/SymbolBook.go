package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/symbols/book
	Doc: https://developer-pro.bitmart.com/en/spot/#get-depth
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Depth
	var ac, err = client.GetSpotSymbolBook("BTC_USDT", 2, 5)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
