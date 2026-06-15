package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/sub-account/v1/transfer-history
Doc: https://developer.bitmart.com/spot/spot-sub/transfer-history
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Account Transfer History (For Main and Sub-Account) (KEYED)
	var ac, err = client.GetSpotSubAccountTransferHistory("spot to spot", 100)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
