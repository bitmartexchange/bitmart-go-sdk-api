package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/cross_collateral/interest_log
Doc: https://developer.bitmart.com/futures/futures-trading/cross-collateral-interest-log
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

	// Get Cross Collateral Interest Log (KEYED)
	var ac, err = client.GetContractCrossCollateralInterestLog(map[string]interface{}{
		"coin_code": "USDT",
		"size":      1,
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
