package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"time"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

// https://developer-pro.bitmart.com/en/spot/#private-order-progress
func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := bitmart.NewWS(bitmart.Config{
		WsUrl:     bitmart.WS_URL_USER,
		ApiKey:    yourApiKey,
		SecretKey: yourSecretKey,
		Memo:      yourMemo,
	})

	_ = ws.Connection(OnMessage)

	// 【Private】Order Progress
	channels := []string{
		"spot/user/order:BTC_USDT",
	}

	ws.SubscribeWithLogin(channels)

	// Just test, Please do not use in production.
	time.Sleep(60 * time.Second)

}
