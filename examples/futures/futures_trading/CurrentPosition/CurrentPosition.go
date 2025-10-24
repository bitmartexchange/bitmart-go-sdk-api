package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/private/position
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-current-position-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Current Position (KEYED)
	var ac, err = client.GetContractPosition()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	ac2, err2 := client.GetContractPosition(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2)
	}

	// Test with account parameter
	ac3, err3 := client.GetContractPosition(map[string]interface{}{
		"symbol":  "BTCUSDT",
		"account": "futures",
	})
	if err3 != nil {
		log.Panic(err3)
	} else {
		log.Println(ac3)
	}
}
