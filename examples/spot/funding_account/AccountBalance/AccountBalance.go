package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/v1/wallet
Doc: https://developer-pro.bitmart.com/en/spot/#get-account-balance-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Account Balance (KEYED)
	var ac, err = client.GetSpotAccountWallet()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.GetSpotAccountWallet(map[string]interface{}{
		"currency": "BTC",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
