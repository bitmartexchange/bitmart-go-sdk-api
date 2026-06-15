package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/sub-account/main/v1/transfer-list
Doc: https://developer.bitmart.com/spot/spot-sub/transfer-list
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Sub-Account Transfer History (For Main Account) (KEYED); accountName optional
	var ac, err = client.GetSpotSubAccountMainTransferList("spot to spot", 100, map[string]interface{}{
		"accountName": "subAccountName",
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
