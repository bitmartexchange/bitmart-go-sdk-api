package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/affiliate/rebate-api
Doc: https://developer.bitmart.com/futures/futures-affiliate/rebate-api
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

	// Get Affiliate API Rebate of a Single User (KEYED)
	var ac, err = client.GetContractAffiliateRebateApi(123456789, 1735660800, 1740844800)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
