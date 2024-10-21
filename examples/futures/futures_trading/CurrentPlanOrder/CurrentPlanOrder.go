package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/current-plan-order
Doc: https://developer-pro.bitmart.com/en/futures/#get-all-current-plan-orders-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get All Current Plan Orders (KEYED)
	var ac, err = client.GetContractCurrentPlanOrders()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.GetContractCurrentPlanOrders(map[string]interface{}{
		"symbol":    "BTCUSDT",
		"type":      "limit",
		"limit":     10,
		"plan_type": "plan",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
