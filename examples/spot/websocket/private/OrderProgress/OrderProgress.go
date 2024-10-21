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
	wg.Add(1)

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := bitmart.NewSpotWSClient(bitmart.Config{
		WsUrl:     bitmart.SPOT_WS_USER,
		ApiKey:    yourApiKey,
		SecretKey: yourSecretKey,
		Memo:      yourMemo,
	}, Callback)

	// login
	ws.Login()

	// 【Private】Order Progress
	ws.Send(`{"op": "subscribe", "args": ["spot/user/order:BTC_USDT"]}`)

	// Just test, Please do not use in production.
	wg.Wait()

}
