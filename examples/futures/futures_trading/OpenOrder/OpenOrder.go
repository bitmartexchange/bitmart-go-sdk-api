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
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order Detail (KEYED)
	var ac, err = client.GetContractOpenOrders()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	ac2, err2 := client.GetContractOpenOrders(map[string]interface{}{
		"symbol":      "BTCUSDT",
		"type":        "limit",
		"order_state": "all",
		"limit":       10,
	})

	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}
}
