package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v1/margin/isolated/repay
Doc: https://developer-pro.bitmart.com/en/spot/#margin-repay-isolated-signed
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

	// Margin Repay (Isolated) (SIGNED)
	var ac, err = client.MarginRepayIsolated("BTC_USDT", "BTC", "1")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
