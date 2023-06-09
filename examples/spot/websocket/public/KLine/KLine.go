package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"time"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

// https://developer-pro.bitmart.com/en/spot/#public-kline-channel
func main() {
	ws := bitmart.NewWS(bitmart.Config{WsUrl: bitmart.WS_URL})

	_ = ws.Connection(OnMessage)

	// 【Public】KLine Channel
	channels := []string{
		"spot/kline1m:BTC_USDT",
	}

	ws.SubscribeWithoutLogin(channels)

	// Just test, Please do not use in production.
	time.Sleep(60 * time.Second)
}
