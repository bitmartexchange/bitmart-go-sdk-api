package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/public/funding-rate-v2
Doc: https://developer.bitmart.com/futures/futures-market/funding-rate-v2
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Get Current Funding Rate V2 (symbol optional)
	var ac, err = client.GetContractFundingRateV2(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
