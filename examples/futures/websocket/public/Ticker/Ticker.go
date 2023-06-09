package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"time"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

// https://developer-pro.bitmart.com/en/futures/#public-ticker-channel
func main() {
	ws := bitmart.NewWSContract(bitmart.Config{WsUrl: bitmart.CONTRACT_WS_URL})

	_ = ws.Connection(OnMessage)

	// 【Public】Ticker Channel
	channels := []string{
		"futures/ticker",
	}

	ws.SubscribeWithoutLogin(channels)

	// Just test, Please do not use in production.
	time.Sleep(60 * time.Second)
}
