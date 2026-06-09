package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"os"
	"sync"
)

func Callback(message string) {
	fmt.Println("------------------------>" + message)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ws := bitmart.NewSpotWSClient(bitmart.Config{
		WsUrl:        bitmart.SPOT_WS_URL,
		CustomLogger: bitmart.NewCustomLogger(bitmart.INFO, os.Stdout),
	}, Callback)

	// 【Public】Best Order (Book Ticker) Channel
	ws.Send(`{"op": "subscribe", "args": ["spot/bookTicker:BTC_USDT"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}
