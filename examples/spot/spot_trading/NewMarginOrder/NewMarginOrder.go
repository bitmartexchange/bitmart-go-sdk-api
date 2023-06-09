package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /spot/v1/margin/submit_order
	Doc: https://developer-pro.bitmart.com/en/spot/#new-margin-order-v1-signed
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

	// New Margin Order(v1) (SIGNED)
	var ac, err = client.PostMarginSubmitOrder(bitmart.MarginOrder{
		Symbol:        "BTC_USDT",
		Side:          "sell",
		Type:          "limit",
		ClientOrderId: "",
		Size:          "0.1",
		Price:         "880000",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
