package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
GET /contract/private/trades
Doc: https://developer-pro.bitmart.com/en/futures/#get-order-trade-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Order Trade (KEYED)
	var ac, err = client.GetContractTrades("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

	now := time.Now().Unix()
	ac2, err2 := client.GetContractTrades("BTCUSDT", map[string]interface{}{
		"start_time": int(now - 3600),
		"end_time":   int(now),
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
