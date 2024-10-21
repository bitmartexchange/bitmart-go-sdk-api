package bitmart

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

func TestSubscribePublic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	ws := NewSpotWSClient(Config{
		WsUrl:        SPOT_WS_URL,
		CustomLogger: NewCustomLogger(DEBUG, os.Stdout),
	}, OnMessage)

	ws.Send(`{"op": "subscribe", "args": ["spot/ticker:BTC_USDT"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}

func TestSubscribePrivate(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	ws := NewSpotWSClient(Config{
		WsUrl:        SPOT_WS_USER,
		ApiKey:       "your api key",
		SecretKey:    "your secret key",
		Memo:         "your memo",
		CustomLogger: NewCustomLogger(DEBUG, os.Stdout),
	}, OnMessage)

	ws.Login()

	ws.Send(`{"op": "subscribe", "args": ["spot/user/order:BTC_USDT"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}
