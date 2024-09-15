package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /spot/v1/margin/isolated/account
Doc: https://developer-pro.bitmart.com/en/spot/#get-margin-account-details-isolated-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Margin Account Details(Isolated) (KEYED)
	var ac, err = client.GetMarginAccountDetailsIsolated()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	var ac2, err2 = client.GetMarginAccountDetailsIsolated(map[string]interface{}{
		"symbol": "BTC_USDT",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}
}
