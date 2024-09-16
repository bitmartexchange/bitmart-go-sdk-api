package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v3/cancel_order
Doc: https://developer-pro.bitmart.com/en/spot/#cancel-order-v3-signed
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

	// Cancel Order(v3) (SIGNED)
	var ac, err = client.PostSpotCancelOrder("BTC_USDT", map[string]interface{}{
		"order_id": "112121212",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.PostSpotCancelOrder("BTC_USDT", map[string]interface{}{
		"client_order_id": "112121212",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
