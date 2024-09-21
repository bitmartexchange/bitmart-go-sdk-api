package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/cancel-order
Doc: https://developer-pro.bitmart.com/en/futures/#cancel-order-signed
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

	// Cancel Order (SIGNED)
	var ac, err = client.PostContractCancelOrder(
		"ETHUSDT",
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.PostContractCancelOrder(
		"ETHUSDT",
		map[string]interface{}{
			"order_id": "220906179559421",
		},
	)
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

	var ac3, err3 = client.PostContractCancelOrder(
		"ETHUSDT",
		map[string]interface{}{
			"client_order_id": "220906179559421",
		},
	)
	if err3 != nil {
		log.Panic(err3)
	} else {
		log.Println(ac3.Response)
	}

}
