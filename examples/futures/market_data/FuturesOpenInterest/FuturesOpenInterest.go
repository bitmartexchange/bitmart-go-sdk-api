package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/public/open-interest
Doc: https://developer-pro.bitmart.com/en/futures/#get-futures-openinterest
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Get Futures Open Interest
	var ac, err = client.GetContractOpenInterest("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
