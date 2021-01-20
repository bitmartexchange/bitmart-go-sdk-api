package bitmart

import (
	"bytes"
	_ "bytes"
	"compress/flate"
	_ "compress/flate"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"io/ioutil"
	_ "io/ioutil"
	"time"

	"log"
	"os"
	"os/signal"
	_ "runtime/debug"
	"sync"
	"syscall"
	_ "time"
)

type CloudWS struct {
	Config  Config
	Conn    *websocket.Conn

	readMessage  Callback
	close        chan interface{}
	signal       chan os.Signal
	reconnectUseChannels  []string
	reconnectUseLogin     bool
	isClose               bool

	processMut sync.Mutex
}

func NewWS(config Config) *CloudWS {
	var ws CloudWS
	ws.Config = config
	return &ws
}

type Callback func(message string)

func (ws *CloudWS) Connection(callback Callback) error {
	log.Printf("Connecting to %s", ws.Config.WsUrl)
	c, _, err := websocket.DefaultDialer.Dial(ws.Config.WsUrl, nil)

	if err != nil {
		log.Fatalf("Connect failed:%+v", err)
		return err
	}

	if ws.Config.IsPrint {
		log.Printf("Connected to %s", ws.Config.WsUrl)
	}
	ws.Conn = c

	ws.isClose = false
	ws.readMessage = callback
	ws.close = make(chan interface{})
	ws.signal = make(chan os.Signal)

	signal.Notify(ws.signal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return nil
}


// Support public channel and private channel
func (ws *CloudWS) SubscribeWithLogin(channels []string) {
	if err := ws.login(); err != nil {
		return
	}

	ws.reconnectUseLogin = true
	ws.reconnectUseChannels = channels
	if err := ws.subscribe(OpParam{
		Op:   "subscribe",
		Args: channels,
	}); err != nil {
		ws.stop()
	}

}

// Only support public channel
func (ws *CloudWS) SubscribeWithoutLogin(channels []string) {
	ws.reconnectUseLogin = false
	ws.reconnectUseChannels = channels
	if err := ws.subscribe(OpParam{
		Op:   "subscribe",
		Args: channels,
	}); err != nil {
		ws.stop()
	}

}



type OpParam struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

func (ws *CloudWS) login() error {
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

	var result map[string]string
	if err := json.Unmarshal(message, &result); err == nil {
		if result["event"] == "login" && result["errorCode"] != "" {
			ws.stop()
			return errors.New("login failed")
		}
	}

	return nil
}


func (ws *CloudWS) subscribe(opParam OpParam) error {
	message, err := json.Marshal(opParam)
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


func (ws *CloudWS) stop() {
	defer func() {
		a := recover()
		log.Printf("Stop End. Recover msg: %+v", a)
	}()

	ws.isClose = true
	close(ws.close)
}

func (ws *CloudWS) gzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)
}


func (ws *CloudWS) reconnection() {
	time.Sleep(time.Duration(2)*time.Second)

	log.Printf("Reconnecting to %s", ws.Config.WsUrl)
	c, _, err := websocket.DefaultDialer.Dial(ws.Config.WsUrl, nil)

	if err != nil {
		log.Fatalf("Connect failed. dial:%+v", err)
	}

	if ws.Config.IsPrint {
		log.Printf("Connected to %s", ws.Config.WsUrl)
	}

	ws.Conn = c
	if ws.reconnectUseLogin {
		ws.SubscribeWithLogin(ws.reconnectUseChannels)
	} else {
		ws.SubscribeWithoutLogin(ws.reconnectUseChannels)
	}

}


func (ws *CloudWS) keepalive() {
	msg := "ping"
	//if ws.Config.IsPrint {
	//	log.Printf("Send Msg: %s", msg)
	//}

	if err := ws.Conn.WriteMessage(websocket.PingMessage, []byte(msg)); err != nil {
		log.Printf("Send Msg Err: %s", err)
	}
}

func (ws *CloudWS) work() {
	defer func() {
		a := recover()
		log.Printf("Work End. Recover msg: %+v", a)
	}()
	defer ws.stop()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if ws.isClose {
				break
			}
			ws.keepalive()
		}
	}
}


func (ws *CloudWS) receive() {
	defer func() {
		a := recover()
		if a != nil {
			log.Printf("Receive End. Recover msg: %+v", a)
		}
	}()

	for {
		if ws.isClose {
			break
		}

		messageType, message, err := ws.Conn.ReadMessage()
		if err != nil {
			log.Printf("Read Err: %s", err)
			ws.reconnection()
		}

		txtMsg := message
		switch messageType {
		case websocket.BinaryMessage:
			txtMsg, err = ws.gzipDecode(message)
		}

		if ws.Config.IsPrint {
			log.Printf("Receive Message: %s", txtMsg)
		}

		if string(txtMsg) == "pong" {
			return
		}

		ws.readMessage(string(txtMsg))
	}

}

func (ws *CloudWS) finalize() {
	defer func() {
		log.Printf("Finalize End. Connection to WebSocket is closed.")
	}()

	select {
	case <-ws.close:
		log.Printf("Call close")
		if ws.Conn != nil {
			close(ws.signal)
			_ = ws.Conn.Close()
			ws.Conn = nil
			break
		}
	case <-ws.signal:
		log.Printf("Call signal")
		break
	}
}
