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
	var ac, err = client.GetMarginAccountDetailsIsolated("BTC_USDT")

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
