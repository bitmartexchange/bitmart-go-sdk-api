package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/modify-plan-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#modify-plan-order-signed
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

	// Modify Plan Order (SIGNED)
	var ac, err = client.PostContractModifyPlanOrder(
		"ETHUSDT",
		"limit",
		"2000",
		1,
		map[string]interface{}{
			"client_order_id": "123123123123",
			"executive_price": "1450",
		},
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
