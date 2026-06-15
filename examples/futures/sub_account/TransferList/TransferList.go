package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /account/contract/sub-account/main/v1/transfer-list
Doc: https://developer.bitmart.com/futures/futures-sub/transfer-list
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Get Sub-Account Transfer History (For Main Account) (KEYED)
	var ac, err = client.GetContractSubAccountMainTransferList("subAccountName", 100)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
