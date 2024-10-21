package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /account/v1/transfer-contract
Doc: https://developer-pro.bitmart.com/en/futures/#cancel-plan-order-signed
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

	// Transfer (SIGNED)
	var ac, err = client.PostContractTransfer(
		"USDT",
		"100",
		"spot_to_contract",
		5000,
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
