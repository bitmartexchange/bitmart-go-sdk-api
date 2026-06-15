package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/sub-account/main/v1/wallet
Doc: https://developer.bitmart.com/spot/spot-sub/wallet
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Sub-Account Spot Wallet Balance (For Main Account) (KEYED); currency optional
	var ac, err = client.GetSpotSubAccountMainWallet("subAccountName", map[string]interface{}{
		"currency": "USDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
