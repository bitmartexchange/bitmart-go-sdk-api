package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/margin/isolated/repay_record
	Doc: https://developer-pro.bitmart.com/en/spot/#get-repayment-record-isolated-keyed
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Repayment Record(Isolated) (KEYED)
	var ac, err = client.GetRepaymentRecordIsolated(
		"BTC_USDT", "", "", 0, 0, 10)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
