package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/algo/cancel_all
Doc: https://developer.bitmart.com/spot/spot-trading/algo-cancel-all
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

	// Cancel All Algo Orders(v4) (SIGNED)
	var ac, err = client.PostSpotAlgoCancelAll("trigger", map[string]interface{}{
		"symbol": "BTC_USDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
