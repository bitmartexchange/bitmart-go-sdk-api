//go:build ws
// +build ws

package bitmart

import (
	"fmt"
	"sync"
	"testing"
)

func OnMessageContract(message string) {
	fmt.Println("------------------------>")
}

func TestSubscribeContractWithoutLogin(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := NewTestWSContract()

	_ = ws.Connection(OnMessageContract)

	channels := []string{
		// public channel
		WS_PUBLIC_CONTRACT_TICKER,
		CreateChannel(WS_PUBLIC_CONTRACT_DEPTH20, "BTCUSDT"),
		CreateChannel(WS_PUBLIC_CONTRACT_KLINE_1M, "BTCUSDT"),
	}
	ws.SubscribeWithoutLogin(channels)

	// Just test, Please do not use in production.
	wg.Wait()
}

func TestSubscribeContractWithLogin(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := NewTestWSContractWithLogin()
	_ = ws.Connection(OnMessage)

	channels := []string{
		// private channel
		//"asset:USDT",
		WS_USER_CONTRACT_UNICAST,
		WS_USER_CONTRACT_POSITION,
		CreateChannel(WS_USER_CONTRACT_ASSET, "USDT"),
	}
	ws.SubscribeWithLogin(channels)

	// Just test, Please do not use in production.
	wg.Wait()
}
