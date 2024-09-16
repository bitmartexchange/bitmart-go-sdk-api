package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/v1/trade_fee
Doc: https://developer-pro.bitmart.com/en/spot/#get-actual-trade-fee-rate-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Actual Trade Fee Rate (KEYED)
	var ac, err = client.GetActualTradeFeeRate("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
