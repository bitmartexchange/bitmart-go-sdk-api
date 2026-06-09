package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/query/algo/history-orders
Doc: https://developer.bitmart.com/spot/spot-trading/algo-history
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

	// Query Algo History Orders (v4) (SIGNED)
	var ac, err = client.GetSpotAlgoHistoryOrders("BTC_USDT", "trigger", 0, 0, 200, 0)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
