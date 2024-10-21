package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/trade-fee-rate
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-trade-fee-rate-keyed
*/
func main() {

	var yourApiKey = "your api key"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Trade Fee Rate (KEYED)
	var ac, err = client.GetContractTradeFeeRate(
		"BTCUSDT",
	)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
