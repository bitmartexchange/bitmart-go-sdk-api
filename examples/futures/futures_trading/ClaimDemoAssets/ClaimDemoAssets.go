package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/claim
Doc: https://developer.bitmart.com/futures/futures-sim/claim
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

	// Claim Demo Trading Assets (SIGNED) - no request parameters
	var ac, err = client.PostContractClaimDemoAssets()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
