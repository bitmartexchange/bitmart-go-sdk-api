package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/query/client-order
Doc: https://developer-pro.bitmart.com/en/spot/#query-order-by-clientorderid-v4-signed
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

	// Query Order By clientOrderId(v4) (SIGNED)
	var ac, err = client.GetSpotOrderByClientOrderId("pix123120312312312312", "open", 10000)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
