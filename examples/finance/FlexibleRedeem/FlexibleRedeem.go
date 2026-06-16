package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /newearn/cloud/v1/saving/redeem
Doc: https://developer.bitmart.com/finance/finance-earn/redeem
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

	// Redeem Flexible Savings (SIGNED)
	var ac, err = client.PostFinanceFlexibleRedeem("200001", "50", "20000009000000300001")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
