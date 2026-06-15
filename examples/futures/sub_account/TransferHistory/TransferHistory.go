package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/contract/sub-account/v1/transfer-history
Doc: https://developer.bitmart.com/futures/futures-sub/transfer-history
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Account Transfer History (For Main and Sub-Account) (KEYED)
	var ac, err = client.GetContractSubAccountTransferHistory(100)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
