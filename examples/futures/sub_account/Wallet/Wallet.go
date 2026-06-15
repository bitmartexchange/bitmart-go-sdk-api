package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/contract/sub-account/main/v1/wallet
Doc: https://developer.bitmart.com/futures/futures-sub/wallet
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Sub-Account Contract Wallet Balance (For Main Account) (KEYED); currency optional
	var ac, err = client.GetContractSubAccountMainWallet("subAccountName", map[string]interface{}{
		"currency": "USDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
