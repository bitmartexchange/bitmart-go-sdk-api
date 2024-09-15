package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/cancel_orders
Doc: https://developer-pro.bitmart.com/en/spot/#cancel-batch-order-v4-signed
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

	// Cancel Batch Order(v4) (SIGNED)
	var ac, err = client.PostSpotCancelOrders("BTC_USDT", map[string]interface{}{
		"orderIds": []string{"12312312", "12312312312"},
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.PostSpotCancelOrders("BTC_USDT", map[string]interface{}{
		"clientOrderIds": []string{"12312312", "12312312312"},
	})

	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
