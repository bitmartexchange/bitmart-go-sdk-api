package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /spot/v1/margin/isolated/transfer
	Doc: https://developer-pro.bitmart.com/en/spot/#margin-asset-transfer-signed
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

	// Margin Asset Transfer (SIGNED)
	var ac, err = client.MarginAssetTransfer(bitmart.MarginAssetTransfer{
		Symbol:   "BTC_USDT",
		Currency: "BTC",
		Amount:   "1",
		Side:     "in",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
