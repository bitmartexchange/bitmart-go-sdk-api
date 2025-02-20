package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/public/funding-rate-history
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-funding-rate-history
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Get Funding Rate History
	var ac, err = client.GetContractFundingRateHistory("BTCUSDT", 10)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
