package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /account/contract/sub-account/main/v1/sub-to-main
Doc: https://developer.bitmart.com/futures/futures-sub/sub-to-main
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Sub-Account to Main-Account (For Main Account) (SIGNED)
	var ac, err = client.PostContractSubAccountMainSubToMain("8e7c2b5a-uuid", "1", "USDT", "subAccountName")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
