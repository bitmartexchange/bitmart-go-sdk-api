package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
GET /contract/private/transaction-history
Doc: https://developer-pro.bitmart.com/en/futuresv2/#get-transaction-history-keyed
*/
func main() {

	var yourApiKey = "your api key"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Transaction History (KEYED)
	var ac, err = client.GetContractTransactionHistory(map[string]interface{}{
		"symbol":  "BTCUSDT",
		"account": "futures",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	now := time.Now().Unix()
	ac2, err2 := client.GetContractTransactionHistory(map[string]interface{}{
		"flow_type":  1,
		"start_time": int(now-3600) * 1000,
		"end_time":   int(now) * 1000,
		"account":    "futures",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
