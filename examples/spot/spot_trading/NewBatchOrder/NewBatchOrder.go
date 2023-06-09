package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /spot/v2/batch_orders
	Doc: https://developer-pro.bitmart.com/en/spot/#new-batch-order-v2-signed
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

	// New Batch Order(v2) (SIGNED)
	var orderParams [1]bitmart.Order
	orderParams[0] = bitmart.Order{
		Symbol:        "BTC_USDT",
		Side:          "sell",
		Type:          "limit",
		ClientOrderId: "pix123120312312312312",
		Size:          "0.1",
		Price:         "880000",
	}
	ac, err := client.PostSpotBatchOrders(orderParams)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
