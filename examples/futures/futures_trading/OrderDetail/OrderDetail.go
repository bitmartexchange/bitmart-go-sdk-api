package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/order
Doc: https://developer-pro.bitmart.com/en/futures/#get-order-detail-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order Detail (KEYED)
	var ac, err = client.GetContractOrder(
		"BTCUSDT",
		"220609666322019",
	)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
