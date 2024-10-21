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
	wg.Add(3)

	ws := bitmart.NewFuturesWSClient(bitmart.Config{
		WsUrl:        bitmart.FUTURES_WS_URL,
		CustomLogger: bitmart.NewCustomLogger(bitmart.INFO, os.Stdout),
	}, Callback)

	ws.Send(`{"action":"subscribe","args":["futures/klineBin1m:BTCUSDT"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}
