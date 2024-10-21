package bitmart

import (
	_ "bytes"
	_ "compress/flate"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	_ "io/ioutil"
	_ "runtime/debug"
	"time"
	_ "time"
)

type CloudSpotWSClient struct {
	CloudWsClient
}

func NewSpotWSClient(config Config, callback Callback) *CloudSpotWSClient {
	client := &CloudSpotWSClient{}
	client.Init(true, config, callback)
	return client
}

type OpParam struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

type SpotWsRespMsg struct {
	Event     string
	ErrorCode string
}

// SpotLogin /* Send Login Message
func (ws *CloudWsClient) SpotLogin() error {
	timestamp := UTCTime()
	sign, err := HmacSha256Base64Signer(
		PreHashString(timestamp, ws.Config.Memo, "bitmart.WebSocket"), ws.Config.SecretKey)
	if err != nil {
		return err
	}

	loginParam, err := json.Marshal(OpParam{
		Op:   "login",
		Args: []string{ws.Config.ApiKey, timestamp, sign},
	})
	if err != nil {
		return errors.New("login param convert json error")
	}

	ws.Config.CustomLogger.Logf(DEBUG, "Send Msg: %s", loginParam)

	if err := ws.Conn.WriteMessage(websocket.TextMessage, loginParam); err != nil {
		ws.Config.CustomLogger.Logf(ERROR, "Send Msg Err: %s", err)
		return err
	}

	time.Sleep(2 * time.Second)

	return nil
}

// SpotCheckLogin /* check login result
func (ws *CloudWsClient) SpotCheckLogin(message []byte) bool {
	var result SpotWsRespMsg
	// {"errorMessage":"Param sign is wrong","errorCode":"91011","event":"login"}
	if err := json.Unmarshal(message, &result); err == nil {
		if result.Event == "login" && result.ErrorCode != "" {
			ws.Stop()
			return true
		}
	}

	return false
}

// SpotPing /* Send Ping Message
func (ws *CloudWsClient) SpotPing() {
	msg := "ping"
	ws.Config.CustomLogger.Logf(DEBUG, "Send: ping")
	if err := ws.Conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		ws.Config.CustomLogger.Logf(ERROR, "Send 'ping' Err: %s", err)
	}
}
