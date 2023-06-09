package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /spot/v4/query/open-orders
	Doc: https://developer-pro.bitmart.com/en/spot/#current-open-orders-v4-signed
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

	// Current Open Orders(v4) (SIGNED)
	var ac, err = client.GetSpotOpenOrders(
		"BTC_USDT",
		"spot",
		0,
		0,
		10,
		5000)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
