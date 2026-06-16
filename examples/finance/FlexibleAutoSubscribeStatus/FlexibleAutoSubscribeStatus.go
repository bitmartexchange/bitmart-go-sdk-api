package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /newearn/cloud/v1/saving/subscribe/status
Doc: https://developer.bitmart.com/finance/finance-auto/flexible-subscribe-status
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

	// Get Flexible Auto-Subscribe Status (KEYED)
	var ac, err = client.GetFinanceFlexibleAutoSubscribeStatus("1001")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
