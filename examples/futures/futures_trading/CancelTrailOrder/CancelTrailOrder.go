package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/cancel-trail-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#cancel-trail-order-signed
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

	// Cancel Trail Order (SIGNED)
	var ac, err = client.PostContractCancelTrailOrder(
		"ETHUSDT",
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	// Cancel Trail Order (SIGNED)
	var ac2, err2 = client.PostContractCancelTrailOrder(
		"ETHUSDT", map[string]interface{}{
			"order_id": "220906179559421",
		},
	)
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
