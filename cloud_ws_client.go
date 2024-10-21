package bitmart

import (
	"bytes"
	_ "bytes"
	"compress/flate"
	_ "compress/flate"
	"github.com/gorilla/websocket"
	"io"
	_ "io/ioutil"
	"time"

	"os"
	"os/signal"
	_ "runtime/debug"
	"syscall"
	_ "time"
)

type CloudWsClient struct {
	Config Config
	Conn   *websocket.Conn

	readMessage Callback
	close       chan interface{}
	signal      chan os.Signal
	isClose     bool
	isSpot      bool

	reconnectUseMessage []string
	reconnectUseLogin   bool
}

func (ws *CloudWsClient) Init(isSpot bool, config Config, callback Callback) {
	ws.Config = config
	ws.isSpot = isSpot
	ws.readMessage = callback
	ws.reconnectUseLogin = false

	if config.CustomLogger == nil {
		ws.Config.CustomLogger = NewCustomLogger(INFO, os.Stdout)
	}

	_ = ws.connection()
}

// Callback function
type Callback func(message string)

// Connection to websocket
func (ws *CloudWsClient) connection() error {
	ws.Config.CustomLogger.Logf(DEBUG, "Connecting to %s", ws.Config.WsUrl)
	c, _, err := websocket.DefaultDialer.Dial(ws.Config.WsUrl, nil)

	if err != nil {
		ws.Config.CustomLogger.Logf(ERROR, "Connect failed:%+v", err)
		return err
	}

	ws.Config.CustomLogger.Logf(DEBUG, "Connected to %s", ws.Config.WsUrl)
	ws.Conn = c

	ws.isClose = false
	ws.close = make(chan interface{})
	ws.signal = make(chan os.Signal)

	signal.Notify(ws.signal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go ws.work()
	go ws.receive()
	go ws.finalize()
	return nil
}

// Send /* Send message to Server
func (ws *CloudWsClient) Send(message string) error {
	// Check if message already exists
	for _, msg := range ws.reconnectUseMessage {
		if msg == message {
			ws.Config.CustomLogger.Logf(INFO, "Message sent: %s", message)
			return nil
		}
	}

	ws.Config.CustomLogger.Logf(DEBUG, "Send Msg: %s", message)

	if err := ws.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		ws.Config.CustomLogger.Logf(ERROR, "Send Msg Err: %s", err)
		return err
	}

	// If the message does not exist, add it to the reconnectUseMessage
	ws.reconnectUseMessage = append(ws.reconnectUseMessage, message)

	return nil
}

// Login /* Send Login message to Server
func (ws *CloudWsClient) Login() error {
	if ws.isSpot {
		_ = ws.SpotLogin()
	} else {
		_ = ws.FuturesLogin()
	}

	ws.reconnectUseLogin = true
	return nil
}

// Stop /* websocket close
func (ws *CloudWsClient) Stop() {
	defer func() {
		_ = recover()
		ws.Config.CustomLogger.Logf(INFO, "Client Call Stop End")
	}()

	ws.isClose = true
	close(ws.close)
}

// decode message
func (ws *CloudWsClient) gzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			ws.Config.CustomLogger.Logf(ERROR, "ReadCloser Err: %s", err)
		}
	}(reader)

	return io.ReadAll(reader)
}

// reconnect to websocket
func (ws *CloudWsClient) reconnection() {
	time.Sleep(time.Duration(2) * time.Second)

	ws.Config.CustomLogger.Logf(INFO, "Reconnecting to %s", ws.Config.WsUrl)
	c, _, err := websocket.DefaultDialer.Dial(ws.Config.WsUrl, nil)

	if err != nil {
		ws.Config.CustomLogger.Logf(ERROR, "Connect failed. dial:%+v", err)
	}

	ws.Config.CustomLogger.Logf(INFO, "Connected to %s", ws.Config.WsUrl)

	ws.Conn = c
	if ws.reconnectUseLogin {
		_ = ws.Login()
	}

	count := 0
	for _, msg := range ws.reconnectUseMessage {
		_ = ws.Send(msg)
		count++
		if count%100 == 0 {
			time.Sleep(10 * time.Second)
		}
	}

}

// keepalive
func (ws *CloudWsClient) keepalive() {
	if ws.isSpot {
		ws.SpotPing()
	} else {
		ws.FuturesPing()
	}
}

// work for receive message
func (ws *CloudWsClient) work() {
	defer func() {
		a := recover()
		ws.Config.CustomLogger.Logf(INFO, "Work End. Recover msg: %+v", a)
	}()
	defer ws.Stop()

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

// receive message
func (ws *CloudWsClient) receive() {
	defer func() {
		a := recover()
		if a != nil {
			ws.Config.CustomLogger.Logf(INFO, "Receive End. Recover msg: %+v", a)
		}
	}()

	for {
		if ws.isClose { // exits
			break
		}

		messageType, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.Config.CustomLogger.Logf(ERROR, "Read Err: %s", err)
			ws.reconnection()
		}

		txtMsg := message
		switch messageType {
		case websocket.BinaryMessage:
			txtMsg, err = ws.gzipDecode(message)
		}

		ws.Config.CustomLogger.Logf(DEBUG, "Receive Message: %s", txtMsg)

		if string(txtMsg) == "pong" {
			continue
		}

		ws.readMessage(string(txtMsg))

		if ws.reconnectUseLogin {
			if ws.isSpot {
				ws.SpotCheckLogin(message)
			} else {
				ws.FuturesCheckLogin(message)
			}
		}

	}

}

// finalize
func (ws *CloudWsClient) finalize() {
	defer func() {
		ws.Config.CustomLogger.Logf(DEBUG, "Finalize End. Connection to WebSocket is closed.")
	}()

	select {
	case <-ws.close:
		ws.Config.CustomLogger.Logf(DEBUG, "Call close")
		if ws.Conn != nil {
			close(ws.signal)
			_ = ws.Conn.Close()
			ws.Conn = nil
			break
		}
	case <-ws.signal:
		ws.Config.CustomLogger.Logf(DEBUG, "Call signal")
		break
	}
}
