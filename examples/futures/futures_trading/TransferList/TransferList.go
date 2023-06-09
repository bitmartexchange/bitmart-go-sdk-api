package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
	POST /account/v1/transfer-contract-list
	Doc: https://developer-pro.bitmart.com/en/futures/#get-transfer-list-signed
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

	// Get Transfer List (SIGNED)
	now := time.Now().UnixNano() / int64(time.Millisecond)
	var ac, err = client.GetContractTransferList(
		"BTCUSDT",
		now-1000*60*60,
		now,
		1,
		20,
		5000,
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
