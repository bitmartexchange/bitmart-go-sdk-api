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

type CloudWSContract struct {
	CloudWS
}

func NewWSContract(config Config) *CloudWSContract {
	var ws CloudWSContract
	ws.Config = config
	return &ws
}

var pingMsg = &struct {
	Subscribe string
}{
	"ping",
}

type Msg struct {
	Action string
	Args   []string
}

type RespMsg struct {
	Action  string
	Success bool
}

// SubscribeWithLogin Support public channel and private channel
func (ws *CloudWSContract) SubscribeWithLogin(channels []string) {
	if err := ws.login(); err != nil {
		log.Fatalf("Login Err: %s", err)
		return
	}

	ws.reconnectUseLogin = true
	ws.reconnectUseChannels = channels
	if err := ws.subscribe(Msg{
		Action: "subscribe",
		Args:   channels,
	}); err != nil {
		ws.stop()
	}

}

// SubscribeWithoutLogin Only support public channel
func (ws *CloudWSContract) SubscribeWithoutLogin(channels []string) {
	ws.reconnectUseLogin = false
	ws.reconnectUseChannels = channels
	if err := ws.subscribe(Msg{
		Action: "subscribe",
		Args:   channels,
	}); err != nil {
		ws.stop()
	}

}

// login
func (ws *CloudWSContract) login() error {
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

	if ws.Config.IsPrint {
		log.Printf("Send Msg: %s", loginParam)
	}

	if err := ws.Conn.WriteMessage(websocket.TextMessage, loginParam); err != nil {
		log.Printf("Send Msg Err: %s", err)
		return err
	}

	_, message, err := ws.Conn.ReadMessage()
	if err != nil {
		log.Printf("Read Err: %s", err)
	}

	var result RespMsg
	if err := json.Unmarshal(message, &result); err == nil {
		if result.Success != true {
			ws.stop()
			return errors.New(string(message))
		}
	}

	return nil
}

// subscribe
func (ws *CloudWSContract) subscribe(msg Msg) error {
	message, err := json.Marshal(msg)
	if err != nil {
		return errors.New("json convert string error")
	}

	if ws.Config.IsPrint {
		log.Printf("Send Msg: %s", message)
	}

	if err := ws.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Printf("Send Msg Err: %s", err)
		return err
	}

	go ws.work()
	go ws.receive()
	go ws.finalize()

	return nil
}

// keepalive
func (ws *CloudWSContract) keepalive() {
	if err := ws.Conn.WriteJSON(pingMsg); err != nil {
		log.Printf("Send Msg Err: %s", err)
	}
}
