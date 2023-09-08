package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
	"time"
)

/*
GET /spot/quotation/v3/klines
Doc: https://developer-pro.bitmart.com/en/spot/#get-history-k-line-v3
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get History K-Line (V3)
	before := time.Now().Unix()
	after := before - 60*60
	var ac, err = client.GetSpotV3HistoryKline("BTC_USDT", before, after, 15, 10)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
