package bitmart

import (
	"fmt"
	"sync"
	"testing"
)

func OnMessage(message string) {
	fmt.Println("------------------------>")
}

func TestSubscribeWithoutLogin(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)


	ws := NewTestWS()

	_ = ws.Connection(OnMessage)

	channels := []string{
		// public channel
		CreateChannel(WS_PUBLIC_SPOT_TICKER, "BTC_USDT"),
	}
	ws.SubscribeWithoutLogin(channels)

	// Just test, Please do not use in production.
	wg.Wait()
}


func TestSubscribeWithLogin(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := NewTestWS()
	_ = ws.Connection(OnMessage)

	channels := []string{
		// public channel
		CreateChannel(WS_PUBLIC_SPOT_TICKER, "BTC_USDT"),
		// private channel
		CreateChannel(WS_USER_SPOT_ORDER, "BTC_USDT"),
	}
	ws.SubscribeWithLogin(channels)


	// Just test, Please do not use in production.
	wg.Wait()
}

