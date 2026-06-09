package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /spot/v4/query/algo/client-order
Doc: https://developer.bitmart.com/spot/spot-trading/algo-query-client
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

	// Query Algo Order By clientOrderId (v4) (SIGNED)
	var ac, err = client.GetSpotAlgoOrderByClientOrderId("jhjj8h8h8h88h998u9u", "open", 0)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
