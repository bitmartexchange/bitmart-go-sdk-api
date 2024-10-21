package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /contract/private/modify-preset-plan-order
Doc: https://developer-pro.bitmart.com/en/futuresv2/#modify-preset-plan-order-signed
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

	// Modify Preset Plan Order (SIGNED)
	var ac, err = client.PostContractModifyPresetPlanOrder(
		"ETHUSDT",
		"220609666322019",
		map[string]interface{}{
			"preset_take_profit_price":      "2000",
			"preset_stop_loss_price":        "1900",
			"preset_take_profit_price_type": 1,
			"preset_stop_loss_price_type":   1,
		},
	)

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
