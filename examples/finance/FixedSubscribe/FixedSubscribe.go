package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /newearn/cloud/v1/saving/fixed/subscribe
Doc: https://developer.bitmart.com/finance/finance-fixed/subscribe
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

	// Subscribe Fixed Savings (SIGNED)
	var ac, err = client.PostFinanceFixedSubscribe("1001", "10", "20000009000000300002", "OFF")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
