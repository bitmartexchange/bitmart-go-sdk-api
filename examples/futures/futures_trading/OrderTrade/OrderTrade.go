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
	now := time.Now().Unix()
	var ac, err = client.GetContractTrades("BTCUSDT", int(now-3600), int(now))
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
