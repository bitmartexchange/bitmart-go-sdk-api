package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/submit-tp-sl-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#submit-tp-or-sl-order-signed
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

	// Submit TP or SL Order (SIGNED)
	var ac, err = client.PostContractSubmitTpSlOrder(
		"ETHUSDT",
		"take_profit",
		2,
		"2000",
		"1450",
		1,
		map[string]interface{}{
			"plan_category":   1,
			"client_order_id": "123123123123",
			"category":        "limit",
		},
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
