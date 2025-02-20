package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/submit-trail-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#submit-trail-order-signed
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

	// Submit Trail Order (SIGNED)
	var ac, err = client.PostContractTrailOrder(bitmart.ContractTrailOrder{
		Symbol:              "BTCUSDT",
		Side:                1,
		Leverage:            "5",
		OpenType:            "isolated",
		Size:                1,
		ActivationPrice:     "81000",
		CallbackRate:        "2",
		ActivationPriceType: 1,
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
