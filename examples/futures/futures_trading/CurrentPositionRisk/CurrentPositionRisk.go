package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/position-risk
Doc: https://developer-pro.bitmart.com/en/futures/#get-current-position-risk-details-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Current Position Risk Details(KEYED)
	var ac, err = client.GetContractPositionRisk()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	ac2, err2 := client.GetContractPositionRisk(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2)
	}
}
