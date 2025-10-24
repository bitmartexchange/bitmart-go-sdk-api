package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/private/position-v2
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-current-position-v2-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// Get Current Position V2 (KEYED) - All positions
	var ac, err = client.GetContractPositionV2()

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	// Get Current Position V2 (KEYED) - Specific symbol
	var ac2, err2 = client.GetContractPositionV2(map[string]interface{}{
		"symbol": "BTCUSDT",
	})

	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

	// Get Current Position V2 (KEYED) - Specific account
	var ac3, err3 = client.GetContractPositionV2(map[string]interface{}{
		"account": "futures",
	})

	if err3 != nil {
		log.Panic(err3)
	} else {
		log.Println(ac3.Response)
	}

}
