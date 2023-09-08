package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/get-open-orders
Doc: https://developer-pro.bitmart.com/en/futures/#get-all-open-orders-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order Detail (KEYED)
	var ac, err = client.GetContractOpenOrders(
		"BTCUSDT",
		"limit",
		"all",
		5,
	)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
