package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /spot/v2/submit_order
	Doc: https://developer-pro.bitmart.com/en/spot/#new-order-v2-signed
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

	// New Order(v2) (SIGNED)
	var ac, err = client.PostSpotSubmitOrder(bitmart.Order{
		Symbol:        "BTC_USDT",
		Side:          "buy",
		Type:          "limit",
		ClientOrderId: "",
		Size:          "0.1",
		Price:         "8800",
		Notional:      "",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
