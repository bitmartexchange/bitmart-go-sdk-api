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
	var ac, err = client.GetSpotV3Book("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.GetSpotV3Book("BTC_USDT", map[string]interface{}{
		"limit": 10,
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
