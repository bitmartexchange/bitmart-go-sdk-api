package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/batch_orders
Doc: https://developer-pro.bitmart.com/en/spot/#new-batch-order-v4-signed
*/
func main() {

	var yourApiKey = "a0a053a31cca9346e52844ee3f13eaa9867c4a51"
	var yourSecretKey = "0c45bc61f75ab7b22f34593af965322acfacce90709fdea5eb176c7ab7a63061"
	var yourMemo = "mytest"

	client := bitmart.NewClient(bitmart.Config{
		Url:           "https://api-cloud.bitmartgcp-test.com",
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
	})

	ac, err := client.PostSpotBatchOrders("BTC_USDT", orderParams[:])
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
