package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/query/algo/order
Doc: https://developer.bitmart.com/spot/spot-trading/algo-query-order
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

	// Query Algo Order By orderId (v4) (SIGNED)
	var ac, err = client.GetSpotAlgoOrderByOrderId("137478201134228132", "open", 0)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
