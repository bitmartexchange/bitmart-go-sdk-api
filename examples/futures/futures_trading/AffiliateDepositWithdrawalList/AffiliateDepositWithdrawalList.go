package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/affiliate/deposit-withdrawal-list
Doc: https://developer.bitmart.com/futures/futures-affiliate/deposit-withdrawal-list
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

	// Get Invited Users Deposit/Withdraw List (KEYED), type optional (1 deposit, 2 withdraw)
	var ac, err = client.GetContractAffiliateDepositWithdrawalList(123456789, 1735660800, 1740844800, 1, 50, map[string]interface{}{
		"type": 1,
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
