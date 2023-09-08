package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
	GET /spot/v1/symbols/kline
	Doc: https://developer-pro.bitmart.com/en/spot/#get-k-line
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get K-Line
	to := time.Now().Unix()
	from := to - 60*60

	var ac, err = client.GetSpotSymbolKline("BMX_ETH", from, to, 15)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
