package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /account/v1/withdraw/address/list
Doc: https://developer-pro.bitmart.com/en/spot/#withdraw-address-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Query Withdraw Address List (KEYED)
	var ac, err = client.GetAccountWithdrawAddressList()

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
