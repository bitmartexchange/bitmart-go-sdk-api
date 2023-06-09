package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"time"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

// https://developer-pro.bitmart.com/en/futures/#private-assets-channel
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := bitmart.NewWSContract(bitmart.Config{
		WsUrl:     bitmart.CONTRACT_WS_PRIVATE_URL,
		ApiKey:    yourApiKey,
		SecretKey: yourSecretKey,
		Memo:      yourMemo,
	})

	_ = ws.Connection(OnMessage)

	// 【Private】Assets Channel
	channels := []string{
		"futures/asset:USDT",
	}
	ws.SubscribeWithLogin(channels)

	// Just test, Please do not use in production.
	time.Sleep(60 * time.Second)
}
