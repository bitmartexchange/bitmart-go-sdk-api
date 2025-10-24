package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/private/order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-order-detail-keyed
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

	// Get Order Detail with account parameter
	var ac2, err2 = client.GetContractOrder("BTCUSDT", "220609666322019", map[string]interface{}{
		"account": "futures",
	})

	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
