package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/v1/deposit/address
Doc: https://developer-pro.bitmart.com/en/spot/#deposit-address-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Deposit Address (KEYED)
	var ac, err = client.GetAccountDepositAddress("BTC")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
