package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/assets-detail
Doc: https://developer-pro.bitmart.com/en/futures/#futures-account-data
*/
func main() {

	var yourApiKey = "Your API KEY"

	client := bitmart.NewClient(bitmart.Config{
		Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		TimeoutSecond: 5,
	})

	// Get Contract Assets (KEYED)
	var ac, err = client.GetContractAssetsDetail()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
