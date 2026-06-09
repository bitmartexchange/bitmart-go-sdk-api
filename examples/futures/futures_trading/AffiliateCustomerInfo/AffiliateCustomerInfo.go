package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/affiliate/aff-customer-info
Doc: https://developer.bitmart.com/futures/futures-affiliate/aff-customer-info
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

	// Get Invited User Futures Account Info (KEYED)
	var ac, err = client.GetContractAffiliateCustomerInfo(123456789)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
