package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /newearn/cloud/v1/saving/fixed/redeem
Doc: https://developer.bitmart.com/finance/finance-fixed/early-redeem
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

	// Early Redeem Fixed Savings (SIGNED)
	var ac, err = client.PostFinanceFixedRedeem("200001", "20000009000000300003")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
