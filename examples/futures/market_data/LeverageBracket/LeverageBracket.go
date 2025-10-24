package main

import (
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/public/leverage-bracket
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-current-leverage-risk-limit
*/
func main() {

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		TimeoutSecond: 5,
	})

	// Get Current Leverage Risk Limit
	var ac, err = client.GetContractLeverageBracket(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
