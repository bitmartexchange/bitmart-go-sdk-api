package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/v1/currencies
Doc: https://developer-pro.bitmart.com/en/spot/#get-currency-list
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Currency List
	var ac, err = client.GetSpotCurrencies()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
