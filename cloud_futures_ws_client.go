package bitmart

import (
	_ "bytes"
	_ "compress/flate"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	_ "io/ioutil"
	"log"
	_ "runtime/debug"
	_ "time"
)

type CloudFuturesWSClient struct {
	CloudWsClient
}

func NewFuturesWSClient(config Config, callback Callback) *CloudSpotWSClient {
	client := &CloudSpotWSClient{}
	client.Init(false, config, callback)

	return client
}

type Msg struct {
	Action string
	Args   []string
}

type FuturesWsRespMsg struct {
	Action  string
	Success bool
}

// FuturesLogin /* Send Login Message
func (ws *CloudWsClient) FuturesLogin() error {
	timestamp := UTCTime()
	sign, err := HmacSha256Base64Signer(
		PreHashString(timestamp, ws.Config.Memo, "bitmart.WebSocket"), ws.Config.SecretKey)
	if err != nil {
		return err
	}

	loginParam, err := json.Marshal(Msg{
		Action: "access",
		Args:   []string{ws.Config.ApiKey, timestamp, sign, "web"},
	})
	if err != nil {
		return errors.New("login param convert json error")
	}

	ws.Config.CustomLogger.Logf(DEBUG, "Send Msg: %s", loginParam)

	if err := ws.Conn.WriteMessage(websocket.TextMessage, loginParam); err != nil {
		log.Printf("Send Msg Err: %s", err)
		return err
	}
	return nil
}

// FuturesCheckLogin /* check login result
func (ws *CloudWsClient) FuturesCheckLogin(message []byte) bool {
	var result FuturesWsRespMsg
	// {"errorMessage":"Param sign is wrong","errorCode":"91011","event":"login"}
	if err := json.Unmarshal(message, &result); err == nil {
		if result.Action == "access" && result.Success != true {
			ws.Stop()
			return true
		}
	}

	return false
}

// FuturesPing /* Send Ping Message
func (ws *CloudWsClient) FuturesPing() {
	// {"action":"ping"}
	ws.Config.CustomLogger.Logf(DEBUG, `Send: {"action":"ping"}`)
	if err := ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"action":"ping"}`)); err != nil {
		ws.Config.CustomLogger.Logf(ERROR, "Send 'ping' Err: %s", err)
	}
}
