package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /account/sub-account/main/v1/main-to-sub
Doc: https://developer.bitmart.com/spot/spot-sub/main-to-sub
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        "Your API KEY",
		SecretKey:     "Your Secret KEY",
		Memo:          "Your Memo",
		TimeoutSecond: 5,
	})

	// Main-Account to Sub-Account (For Main Account) (SIGNED)
	var ac, err = client.PostSpotSubAccountMainMainToSub("8e7c2b5a-uuid", "1", "USDT", "subAccountName")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
