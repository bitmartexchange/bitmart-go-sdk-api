package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/v2/deposit-withdraw/history
Doc: https://developer-pro.bitmart.com/en/spot/#get-deposit-and-withdraw-history-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Deposit And Withdraw History (KEYED)
	var ac, err = client.GetDepositWithdrawHistory("withdraw", 10)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.GetDepositWithdrawHistory("withdraw", 10, map[string]interface{}{
		"currency": "BTC",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
