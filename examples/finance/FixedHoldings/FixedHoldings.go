package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /newearn/cloud/v1/saving/fixed/earn
Doc: https://developer.bitmart.com/finance/finance-fixed/query-holdings
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

	// Get Fixed Savings Holdings (KEYED)
	var ac, err = client.GetFinanceFixedHoldings(1, 10, map[string]interface{}{"coinName": "USDT"})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
