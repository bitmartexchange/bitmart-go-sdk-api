[![Logo](https://img.bitmart.com/static-file/public/sdk/sdk_logo.png)](https://bitmart.com)


BitMart-Go-SDK-API
=========================
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/bitmartexchange/bitmart-go-sdk-api)
[![Go version](https://shields.io/badge/Go-v1.15-blue)](https://go.dev/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/Telegram-Join%20Us-blue?logo=Telegram)](https://t.me/bitmart_api)

[BitMart Exchange official](https://bitmart.com) Go client for the BitMart Cloud API.




Feature
=========================
- Provides exchange quick trading API
- Easier withdrawal
- Efficiency, higher speeds, and lower latencies
- Priority in development and maintenance
- Dedicated and responsive technical support
- Supported APIs:
    - `/spot/*`
    - `/contract/*`
    - `/account/*`
- Supported websockets:
  - Spot WebSocket Market Stream
  - Spot User Data Stream
  - futures User Data Stream
  - futures WebSocket Market Stream
- Test cases and examples



Installation
=========================
```git
go get -u github.com/bitmartexchange/bitmart-go-sdk-api
```

To reference the package in your code, use the following import statement:

```go
import (
    "github.com/bitmartexchange/bitmart-go-sdk-api"
)
```



Documentation
=========================
[API Documentation](https://developer-pro.bitmart.com/en/spot/#change-log)

Example
=========================

#### Spot Market Endpoints Example

```go
package main

import (
  "github.com/bitmartexchange/bitmart-go-sdk-api"
  "log"
)

func main() {
  client := bitmart.NewClient(bitmart.Config{})

  // Get List of Trading Pairs
  var ac, _ = client.GetSpotSymbol()
  log.Println(ac.Response)
  
  // Get Ticker of a Trading Pair (V3)
  ac, _ = client.GetSpotV3Ticker("BTC_USDT")
  log.Println(ac.Response)
  
  // Get Ticker of All Pairs (V3)
  ac, _ = client.GetSpotV3Tickers()
  log.Println(ac.Response)
  
  // Get Depth (V3)
  ac, _ = client.GetSpotV3Book("BTC_USDT")
  log.Println(ac.Response)

}
```


#### Spot / Margin Trading Endpoints Example

```go

package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// New Order(v2) (SIGNED)
	var ac, err = client.PostSpotSubmitOrder(bitmart.Order{
		Symbol:        "BTC_USDT",
		Side:          "buy",
		Type:          "limit",
		ClientOrderId: "",
		Size:          "0.1",
		Price:         "8800",
		Notional:      "",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}



```

Please find `examples/spot` folder to check for more endpoints.

---


#### Spot Websocket Endpoints : Subscribe Public Channel


```go
package main

import (
  "fmt"
  "github.com/bitmartexchange/bitmart-go-sdk-api"
  "os"
  "sync"
)

func Callback(message string) {
  fmt.Println("------------------------>" + message)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(1)

  ws := bitmart.NewSpotWSClient(bitmart.Config{
    WsUrl:        bitmart.SPOT_WS_URL,
  }, Callback)

  // Ticker
  ws.Send(`{"op": "subscribe", "args": ["spot/ticker:BTC_USDT"]}`)

  // Kline
  ws.Send(`{"op": "subscribe", "args": ["spot/kline1m:BTC_USDT"]}`)

  // Trade
  ws.Send(`{"op": "subscribe", "args": ["spot/trade:BTC_USDT"]}`)

  // Depth
  ws.Send(`{"op": "subscribe", "args": ["spot/depth5:BTC_USDT"]}`)

  // Just test, Please do not use in production.
  wg.Wait()
}

```


#### Spot Websocket Endpoints : Subscribe Private Channel

```go
package main

import (
  "fmt"
  "github.com/bitmartexchange/bitmart-go-sdk-api"
  "sync"
)

func Callback(message string) {
  fmt.Println("------------------------>" + message)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(1)

  var yourApiKey = "Your API KEY"
  var yourSecretKey = "Your Secret KEY"
  var yourMemo = "Your Memo"

  ws := bitmart.NewSpotWSClient(bitmart.Config{
    WsUrl:     bitmart.SPOT_WS_USER,
    ApiKey:    yourApiKey,
    SecretKey: yourSecretKey,
    Memo:      yourMemo,
  }, Callback)

  // login
  ws.Login()

  // 【Private】Balance Change
  ws.Send(`{"op": "subscribe", "args": ["spot/user/balance:BALANCE_UPDATE"]}`)

  // Just test, Please do not use in production.
  wg.Wait()

}
```


Please find `examples/spot/websocket` folder to check for more endpoints.


---

#### Futures Market Endpoints

```go
package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
    "log"
)

func main() {
	client := bitmart.NewClient(bitmart.Config{
      Url: bitmart.API_URL_V2_PRO,
    })

	// Get Contract Details
    var ac, _ = client.GetContractDetails("BTCUSDT")
    log.Println(ac.Response)
    // Get Current Funding Rate
    ac, _ = client.GetContractFundingRate("BTCUSDT")
	log.Println(ac.Response)
    
    // Get Futures Open Interest
    ac, _ = client.GetContractOpenInterest("BTCUSDT")
    log.Println(ac.Response)
  
    // Get Market Depth
    ac, _ = client.GetContractDepth("BTCUSDT")
    log.Println(ac.Response)
}

```

#### Futures Trading Endpoints


```go

package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	client := bitmart.NewClient(bitmart.Config{
        Url:           bitmart.API_URL_V2_PRO,
		ApiKey:        yourApiKey,
		SecretKey:     yourSecretKey,
		Memo:          yourMemo,
		TimeoutSecond: 5,
	})

	// Submit Order (SIGNED)
	var ac, err = client.PostContractSubmitOrder(bitmart.ContractOrder{
		Symbol:   "ETHUSDT",
		Side:     4,
		Type:     "limit",
		Leverage: "1",
		OpenType: "isolated",
		Size:     10,
		Price:    "2000",
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}


```


Please find `examples/futures` folder to check for more endpoints.

---


#### Futures Websocket Endpoints : Subscribe Public Channel

```go
package main

import (
  "fmt"
  "github.com/bitmartexchange/bitmart-go-sdk-api"
  "os"
  "sync"
)

func Callback(message string) {
  fmt.Println("------------------------>" + message)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(3)

  ws := bitmart.NewFuturesWSClient(bitmart.Config{
    WsUrl:        bitmart.FUTURES_WS_URL,
    CustomLogger: bitmart.NewCustomLogger(bitmart.INFO, os.Stdout),
  }, Callback)

  ws.Send(`{"action":"subscribe","args":["futures/ticker"]}`)

  ws.Send(`{"action":"subscribe","args":["futures/depth20:BTCUSDT"]}`)

  ws.Send(`{"action":"subscribe","args":["futures/klineBin1m:BTCUSDT"]}`)

  ws.Send(`{"action":"subscribe","args":["futures/trade:BTCUSDT"]}`)

  // Just test, Please do not use in production.
  wg.Wait()
}

```

#### Futures Websocket Endpoints : Subscribe Private Channel

```go
package main

import (
  "fmt"
  "github.com/bitmartexchange/bitmart-go-sdk-api"
  "sync"
)

func Callback(message string) {
  fmt.Println("------------------------>" + message)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(3)

  var yourApiKey = "Your API KEY"
  var yourSecretKey = "Your Secret KEY"
  var yourMemo = "Your Memo"

  ws := bitmart.NewFuturesWSClient(bitmart.Config{
    WsUrl:     bitmart.FUTURES_WS_USER,
    ApiKey:    yourApiKey,
    SecretKey: yourSecretKey,
    Memo:      yourMemo,
  }, Callback)

  // login
  ws.Login()

  // 【Private】Assets Channel
  ws.Send(`{"action": "subscribe","args":["futures/asset:USDT", "futures/asset:BTC", "futures/asset:ETH"]}`)

  // Just test, Please do not use in production.
  wg.Wait()
}

```


Extra Options
=========================

### Authentication

```go
client := bitmart.NewClient(bitmart.Config{
    ApiKey:    yourApiKey,
    SecretKey: yourSecretKey,
    Memo:      yourMemo,
})
```

### Timeout
Through the `bitmart.Config` configuration class, you can set the timeout period for http requests. If not set, the default is 30 seconds.

```go
client := bitmart.NewClient(bitmart.Config{
    TimeoutSecond: 5,
})
```


### Logging
If you want to print out the request information, you can do so by setting the log level to `DEBUG`.

```go
client := bitmart.NewClient(bitmart.Config{
    CustomLogger: bitmart.NewCustomLogger(bitmart.DEBUG, os.Stdout)
})
```

Can I write the log to a file? Of course you can. Here is an example of writing a file:

```go
// Create a log file output
file, err3 := os.OpenFile("./custom_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err3 != nil {
    log.Fatalf("Failed to open log file: %v", err3)
}
defer file.Close()

client := bitmart.NewClient(bitmart.Config{
    CustomLogger: bitmart.NewCustomLogger(bitmart.DEBUG, file)]
})
```

Other settings explained:
* os.O_CREATE: If the custom_log.log file does not exist, Go will automatically create it.
* os.O_WRONLY: The file can only be used for writing.
* os.O_APPEND: The newly written data will be appended to the end of the file without overwriting the existing data in the file.
* The last parameter 0666:
  1. It specifies the permissions for the newly created file.
  2. 0666 means the file's read and write permissions are: the owner, group, and other users can read and write the file.

### Custom request headers
You can add your own request header information here, but please do not fill in `X-BM-KEY, X-BM-SIGN, X-BM-TIMESTAMP`


```go
client := bitmart.NewClient(bitmart.Config{Headers: map[string]string{
      "X-Custom-Header1": "HeaderValue1",
      "X-Custom-Header2": "HeaderValue2",
}})
client.GetSpotV3Ticker("BTC_USDT")
```


### Response Metadata
The bitmart API server provides the endpoint rate limit usage in the header of each response. 
This information can be obtained from the headers property. 
x-bm-ratelimit-remaining indicates the number of times the current window has been used,
x-bm-ratelimit-limit indicates the maximum number of times the current window can be used, 
and x-bm-ratelimit-reset indicates the current window time.


Example:

```
x-bm-ratelimit-mode: IP
x-bm-ratelimit-remaining: 10
x-bm-ratelimit-limit: 600
x-bm-ratelimit-reset: 60
```

This means that this IP can call the endpoint 600 times within 60 seconds, and has called 10 times so far.


```go
var ac, err = client.GetSpotV3Ticker("BTC_USDT")
if err != nil {
    log.Panic(err)
} else {
    log.Println(ac.Limit.Limit)
    log.Println(ac.Limit.Remaining)
    log.Println(ac.Limit.Reset)
    log.Println(ac.Limit.Mode)
}
```
