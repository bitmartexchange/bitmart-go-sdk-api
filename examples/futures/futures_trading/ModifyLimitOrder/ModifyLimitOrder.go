package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
POST /contract/private/modify-limit-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#modify-limit-order-signed
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

	// Modify Limit Order (SIGNED) - Using order_id
	var ac, err = client.PostContractModifyLimitOrder("BTCUSDT", map[string]interface{}{
		"order_id": 12312312,
		"price":    "100",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	// Modify Limit Order (SIGNED) - Using client_order_id
	var ac2, err2 = client.PostContractModifyLimitOrder("BTCUSDT", map[string]interface{}{
		"client_order_id": "123123123123",
		"size":            100,
	})

	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
