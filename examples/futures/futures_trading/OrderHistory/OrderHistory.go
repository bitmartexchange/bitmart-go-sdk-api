package main

import (
	"log"
	"time"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/private/order-history
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-order-history-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order History (KEYED)
	var ac, err = client.GetContractOrderHistory(
		"BTCUSDT",
	)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac)
	}

	now := time.Now().Unix()
	var ac2, err2 = client.GetContractOrderHistory(
		"BTCUSDT",
		map[string]interface{}{
			"start_time": int(now - 3600),
			"end_time":   int(now),
			"account":    "futures",
		},
	)
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2)
	}

	// Test with order_id and client_order_id
	var ac3, err3 = client.GetContractOrderHistory(
		"BTCUSDT",
		map[string]interface{}{
			"order_id":        "123456789",
			"client_order_id": "client_order_123",
			"account":         "futures",
		},
	)
	if err3 != nil {
		log.Panic(err3)
	} else {
		log.Println(ac3)
	}

}
