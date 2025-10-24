package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/private/get-position-mode
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-position-mode-keyed
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

	// Get Position Mode (KEYED)
	var ac, err = client.GetContractPositionMode()

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
