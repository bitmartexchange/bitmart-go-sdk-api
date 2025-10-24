package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
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

	// Get Currencies with specific currencies parameter
	var ac2, err2 = client.GetAccountCurrencies(map[string]interface{}{
		"currencies": "BMX,ETH,BTC",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
