package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
POST /contract/private/set-position-mode
Doc: https://developer-pro.bitmart.com/en/futuresv2/#set-position-mode-signed
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

	// Set Position Mode (SIGNED)
	// positionMode: hedge_mode, one_way_mode
	var ac, err = client.PostContractSetPositionMode("hedge_mode")

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
