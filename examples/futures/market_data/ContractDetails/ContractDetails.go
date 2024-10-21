package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/public/details
Doc: https://developer-pro.bitmart.com/en/futures/#get-contract-details
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Get Contract Details
	var ac, err = client.GetContractDetails("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
