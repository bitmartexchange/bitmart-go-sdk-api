package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/submit-order
Doc: https://developer-pro.bitmart.com/en/futures/#submit-order-signed
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

	// Submit Order (SIGNED)
	var ac, err = client.PostContractSubmitOrder(bitmart.ContractOrder{
		Symbol:   "ETHUSDT",
		Side:     4,
		Type:     "limit",
		Leverage: "1",
		OpenType: "isolated",
		Size:     10,
		Price:    "2000",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
