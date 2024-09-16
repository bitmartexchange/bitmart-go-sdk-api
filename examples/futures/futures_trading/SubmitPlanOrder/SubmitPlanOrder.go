package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/submit-plan-order
Doc: https://developer-pro.bitmart.com/en/futures/#submit-plan-order-signed
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

	// Submit Plan Order (SIGNED)
	var ac, err = client.PostContractPlanOrder(bitmart.ContractPlanOrder{
		Symbol:         "ETHUSDT",
		Side:           4,
		Type:           "limit",
		Leverage:       "1",
		OpenType:       "isolated",
		Size:           10,
		Mode:           1,
		TriggerPrice:   "2000",
		ExecutivePrice: "1800",
		PriceWay:       1,
		PriceType:      1,
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
