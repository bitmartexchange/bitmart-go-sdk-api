package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
GET /contract/private/order-history
Doc: https://developer-pro.bitmart.com/en/futures/#get-order-history-keyed
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
		},
	)
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2)
	}

}
