package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /account/v1/deposit-withdraw/detail
	Doc: https://developer-pro.bitmart.com/en/spot/#get-a-deposit-or-withdraw-detail-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get A Deposit Or Withdraw Detail (KEYED)
	var ac, err = client.GetDepositWithdrawDetail("123123123123123")

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
