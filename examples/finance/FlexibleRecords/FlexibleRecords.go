package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /newearn/cloud/v1/saving/record
Doc: https://developer.bitmart.com/finance/finance-earn/query-history-records
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

	// Get Flexible Savings History Records (KEYED)
	var ac, err = client.GetFinanceFlexibleRecords("subscribe", 1, 10)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
