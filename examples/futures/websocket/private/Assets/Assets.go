package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"sync"
)

func Callback(message string) {
	fmt.Println("------------------------>" + message)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := bitmart.NewFuturesWSClient(bitmart.Config{
		WsUrl:     bitmart.FUTURES_WS_USER,
		ApiKey:    yourApiKey,
		SecretKey: yourSecretKey,
		Memo:      yourMemo,
	}, Callback)

	// login
	ws.Login()

	// 【Private】Assets Channel
	ws.Send(`{"action": "subscribe","args":["futures/asset:USDT", "futures/asset:BTC", "futures/asset:ETH"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}
