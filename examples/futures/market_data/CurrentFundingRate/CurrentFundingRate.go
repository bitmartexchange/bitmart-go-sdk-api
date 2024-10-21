package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/public/funding-rate
Doc: https://developer-pro.bitmart.com/en/futures/#get-current-funding-rate
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Get Current Funding Rate
	var ac, err = client.GetContractFundingRate("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
