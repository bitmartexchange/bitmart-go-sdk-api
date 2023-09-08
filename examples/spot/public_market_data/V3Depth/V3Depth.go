package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/quotation/v3/books
Doc: https://developer-pro.bitmart.com/en/spot/#get-depth-v3
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Depth (V3)
	var ac, err = client.GetSpotV3Book("BTC_USDT", 10)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
