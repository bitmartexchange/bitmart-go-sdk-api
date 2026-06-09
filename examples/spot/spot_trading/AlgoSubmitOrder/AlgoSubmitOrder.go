package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/algo/submit_order
Doc: https://developer.bitmart.com/spot/spot-trading/algo-submit
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

	// Submit Algo Order(v4) (SIGNED)
	// type: "trigger"(plan order) or "tp/sl"(one-way take-profit/stop-loss)
	var ac, err = client.PostSpotAlgoSubmitOrder("BTC_USDT", "buy", "trigger", map[string]interface{}{
		"client_order_id": "jhjj8h8h8h88h998u9u",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
