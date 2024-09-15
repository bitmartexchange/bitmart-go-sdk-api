package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
GET /contract/public/kline
Doc: https://developer-pro.bitmart.com/en/futures/#get-k-line
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get K-line
	now := time.Now().Unix()
	var ac, err = client.GetContractKline("BTCUSDT", int(now-3600), int(now), 15)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
