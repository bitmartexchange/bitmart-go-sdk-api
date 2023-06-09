package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /account/v1/withdraw/charge
	Doc: https://developer-pro.bitmart.com/en/spot/#withdraw-quota-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Withdraw Quota (KEYED)
	var ac, err = client.GetAccountWithdrawCharge("BTC")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
