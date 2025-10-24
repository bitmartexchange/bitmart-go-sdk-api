package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/public/market-trade
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-market-trade
*/
func main() {

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Query the latest trade data
	var ac, err = client.GetContractMarketTrade("BTCUSDT", 50)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
