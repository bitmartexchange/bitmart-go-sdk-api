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
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order History (KEYED)
	now := time.Now().Unix()
	var ac, err = client.GetContractOrderHistory(
		"BTCUSDT",
		int(now-3600),
		int(now),
	)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
