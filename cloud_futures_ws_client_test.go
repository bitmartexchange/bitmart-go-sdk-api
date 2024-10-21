package bitmart

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

func OnMessageFutures(message string) {
	fmt.Println("------------------------>" + message)
}

func TestSubscribeFuturesPublic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := NewFuturesWSClient(Config{
		WsUrl:        FUTURES_WS_URL,
		CustomLogger: NewCustomLogger(DEBUG, os.Stdout),
	}, OnMessageFutures)

	ws.Send(`{"action":"subscribe","args":["futures/ticker"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}

func TestSubscribeFuturesPrivate(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := NewFuturesWSClient(Config{
		WsUrl:        FUTURES_WS_USER,
		ApiKey:       yourApiKey,
		SecretKey:    yourSecretKey,
		Memo:         yourMemo,
		CustomLogger: NewCustomLogger(DEBUG, os.Stdout),
	}, OnMessage)

	ws.Login()

	ws.Send(`{"action":"subscribe","args":["futures/asset:USDT"]}`)

	// Just test, Please do not use in production.
	wg.Wait()
}
