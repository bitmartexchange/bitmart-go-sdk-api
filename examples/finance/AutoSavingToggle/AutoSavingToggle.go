package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /newearn/cloud/v1/saving/subscribe/batch/operate
Doc: https://developer.bitmart.com/finance/finance-auto/toggle
*/
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// Enable/Disable Auto Savings (SIGNED)
	var ac, err = client.PostFinanceAutoSavingToggle("open")
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
