package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/sub-account/main/v1/subaccount-list
Doc: https://developer.bitmart.com/spot/spot-sub/list
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Sub-Account List (For Main Account) (KEYED)
	var ac, err = client.GetSpotSubAccountMainList()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
