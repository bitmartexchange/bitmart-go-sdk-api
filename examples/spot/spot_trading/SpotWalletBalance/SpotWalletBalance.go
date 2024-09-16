package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/v1/wallet
Doc: https://developer-pro.bitmart.com/en/spot/#get-spot-wallet-balance-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Spot Wallet Balance (KEYED)
	var ac, err = client.GetSpotWallet()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
