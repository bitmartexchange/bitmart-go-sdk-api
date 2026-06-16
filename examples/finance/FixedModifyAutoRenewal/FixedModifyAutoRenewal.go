package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /newearn/cloud/v1/saving/fixed/subscribe/operate
Doc: https://developer.bitmart.com/finance/finance-fixed/modify-auto-renewal
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

	// Modify Fixed Savings Auto-Renewal (SIGNED)
	var ac, err = client.PostFinanceFixedModifyAutoRenewal("200001", "OFF")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
