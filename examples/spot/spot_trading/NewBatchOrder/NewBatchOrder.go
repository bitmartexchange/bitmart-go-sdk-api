package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
POST /spot/v4/batch_orders
Doc: https://developer-pro.bitmart.com/en/spot/#new-batch-order-v4-signed
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

	// New Batch Order(v4) (SIGNED)
	var orderParams []bitmart.BatchOrder
	orderParams = append(orderParams, bitmart.BatchOrder{
		Side:          "sell",
		Type:          "limit",
		ClientOrderId: "pix123120312312312313",
		Size:          "0.1",
		Price:         "880000",
	})

	orderParams = append(orderParams, bitmart.BatchOrder{
		Side:          "sell",
		Type:          "market",
		ClientOrderId: "pix123120312312312314",
		Size:          "0.1",
		Notional:      "880000",
		StpMode:       "cancel_maker",
	})

	ac, err := client.PostSpotBatchOrders("BTC_USDT", orderParams[:])
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
