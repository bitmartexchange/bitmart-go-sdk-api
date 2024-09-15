package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/public/depth
Doc: https://developer-pro.bitmart.com/en/futures/#get-market-depth
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get Market Depth
	var ac, err = client.GetContractDepth("BTCUSDT")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
