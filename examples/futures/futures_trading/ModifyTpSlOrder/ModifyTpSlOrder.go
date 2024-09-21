package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/modify-tp-sl-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#modify-tp-sl-order-signed
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

	// Modify TP/SL Order (SIGNED)
	var ac, err = client.PostContractModifyTpSlOrder(
		"ETHUSDT",
		"2100",
		2,
		map[string]interface{}{
			"executive_price": "2100",
			"order_id":        "q232434234234",
			"plan_category":   2,
			"category":        "limit",
		},
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
