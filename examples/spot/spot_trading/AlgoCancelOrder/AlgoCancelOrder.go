package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/algo/cancel_order
Doc: https://developer.bitmart.com/spot/spot-trading/algo-cancel
*/
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// Cancel Algo Order(v4) (SIGNED)
	var ac, err = client.PostSpotAlgoCancelOrder("BTC_USDT", "137478201134228132", "trigger")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
