package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/affiliate/rebate-list
Doc: https://developer.bitmart.com/futures/futures-affiliate/rebate-list
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

	// Get Affiliate Rebate List (KEYED), user_id and time ranges optional
	var ac, err = client.GetContractAffiliateRebateList("USDT", 1, 50, map[string]interface{}{
		"user_id": 123456789,
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
