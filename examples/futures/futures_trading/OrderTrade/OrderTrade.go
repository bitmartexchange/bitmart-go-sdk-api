package main

import (
	"log"
	"time"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

/*
GET /contract/private/trades
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-order-trade-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order Trade (KEYED) - All trades
	var ac, err = client.GetContractTrades()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	// Get Order Trade (KEYED) - With symbol
	ac2, err2 := client.GetContractTrades(map[string]interface{}{
		"symbol": "BTCUSDT",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

	now := time.Now().Unix()
	ac3, err3 := client.GetContractTrades(map[string]interface{}{
		"symbol":     "BTCUSDT",
		"start_time": int(now - 3600),
		"end_time":   int(now),
		"account":    "futures",
	})
	if err3 != nil {
		log.Panic(err3)
	} else {
		log.Println(ac3.Response)
	}

}
