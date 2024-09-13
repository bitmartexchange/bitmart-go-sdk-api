[![Logo](https://img.bitmart.com/static-file/public/sdk/sdk_logo.png)](https://bitmart.com)


BitMart-Go-SDK-API
=========================
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/bitmartexchange/bitmart-go-sdk-api)
[![Go version](https://shields.io/badge/Go-v1.15-blue)](https://go.dev/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


[BitMart Exchange official](https://bitmart.com) Go client for the BitMart Cloud API.




Feature
=========================
- Provides exchange quick trading API
- Easier withdrawal
- Efficiency, higher speeds, and lower latencies
- Priority in development and maintenance
- Dedicated and responsive technical support
- Provide webSocket apis calls
- Supported APIs:
    - `/spot/*`
    - `/contract/*`
    - `/account/*`
    - Spot WebSocket Market Stream
    - Spot User Data Stream
    - Contract User Data Stream
    - Contract WebSocket Market Stream
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
)

func main() {
  client := bitmart.NewClient(bitmart.Config{})

  // Get List of Trading Pairs
  client.GetSpotSymbol()
  
  // Get Ticker of a Trading Pair (V3)
  client.GetSpotV3Ticker("BTC_USDT")
  
  // Get Ticker of All Pairs (V3)
  client.GetSpotV3Tickers()

  // Get Depth (V3)
  client.GetSpotV3Book("BTC_USDT", 10)

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
		log.Println(bitmart.GetResponse(ac))
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
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

func main() {
	ws := bitmart.NewWS(bitmart.Config{WsUrl: bitmart.WS_URL})

	_ = ws.Connection(OnMessage)

	// 【Public】Ticker Channel
	channels := []string{
		"spot/ticker:BTC_USDT",
	}

	ws.SubscribeWithoutLogin(channels)

}


```


#### Spot Websocket Endpoints : Subscribe Private Channel

```go

package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"time"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := bitmart.NewWS(bitmart.Config{
		WsUrl:     bitmart.WS_URL_USER,
		ApiKey:    yourApiKey,
		SecretKey: yourSecretKey,
		Memo:      yourMemo,
	})

	_ = ws.Connection(OnMessage)

	// 【Private】Order Progress
	channels := []string{
		"spot/user/order:BTC_USDT",
	}

	ws.SubscribeWithLogin(channels)

}

```


Please find `examples/spot/websocket` folder to check for more endpoints.


---

#### Futures Market Endpoints

```go
package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

func main() {
	client := bitmart.NewClient(bitmart.Config{})

	// Get Contract Details
	client.GetContractDetails("BTCUSDT")
    // Get Current Funding Rate
    client.GetContractFundingRate("BTCUSDT")
    // Get Futures Open Interest
    client.GetContractOpenInterest("BTCUSDT")
    // Get Market Depth
    client.GetContractDepth("BTCUSDT")
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
		log.Println(bitmart.GetResponse(ac))
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
	"time"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

func main() {
	ws := bitmart.NewWSContract(bitmart.Config{WsUrl: bitmart.CONTRACT_WS_URL})

	_ = ws.Connection(OnMessage)

	// 【Public】Ticker Channel
	channels := []string{
		"futures/ticker",
	}

	ws.SubscribeWithoutLogin(channels)

	// Just test, Please do not use in production.
	time.Sleep(60 * time.Second)
}



```

#### Futures Websocket Endpoints : Subscribe Private Channel

```go

package main

import (
	"fmt"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

func OnMessage(message string) {
	fmt.Println("------------------------>" + message)
}

func main() {

	var yourApiKey = "Your API KEY"
	var yourSecretKey = "Your Secret KEY"
	var yourMemo = "Your Memo"

	ws := bitmart.NewWSContract(bitmart.Config{
		WsUrl:     bitmart.CONTRACT_WS_PRIVATE_URL,
		ApiKey:    yourApiKey,
		SecretKey: yourSecretKey,
		Memo:      yourMemo,
	})

	_ = ws.Connection(OnMessage)

	// 【Private】Assets Channel
	channels := []string{
		"futures/asset:USDT",
	}
	ws.SubscribeWithLogin(channels)

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


### Debug
If you want to print the request and response information, you can set it to true.

```go
client := bitmart.NewClient(bitmart.Config{
    IsPrint: true,
})
```

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
