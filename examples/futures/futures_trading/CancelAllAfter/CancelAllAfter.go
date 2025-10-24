package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
POST /contract/private/cancel-all-after
Doc: https://developer-pro.bitmart.com/en/futuresv2/#timed-cancel-all-orders-signed
*/
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// Timed Cancel All Orders (SIGNED)
	var ac, err = client.PostContractCancelAllAfter("BTCUSDT", 60)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
