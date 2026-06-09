package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/affiliate/trade-list
Doc: https://developer.bitmart.com/futures/futures-affiliate/trade-list
*/
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// Get Affiliate Trade List (KEYED), type: 1 USDT-M, 2 Coin-M; time range optional
	var ac, err = client.GetContractAffiliateTradeList(123456789, 1, 1, 50)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
