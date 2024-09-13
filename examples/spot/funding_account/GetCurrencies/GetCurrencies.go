package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/v1/currencies
Doc: https://developer-pro.bitmart.com/en/spot/#get-currencies
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Currencies (V1)
	var ac, err = client.GetAccountCurrencies()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
