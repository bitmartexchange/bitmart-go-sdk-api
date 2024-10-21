package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/submit-leverage
Doc: https://developer-pro.bitmart.com/en/futures/#submit-leverage-signed
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

	// Submit Leverage (SIGNED)
	var ac, err = client.PostContractSubmitLeverage(
		"BTCUSDT",
		"10",
		"cross",
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
