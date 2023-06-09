[![Logo](https://img.bitmart.com/static-file/public/sdk/sdk_logo.png)](https://bitmart.com)


BitMart-Go-SDK-API
=========================
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/bitmartexchange/bitmart-go-sdk-api)
[![Go version](https://shields.io/badge/Go-v1.15-blue)](https://pypi.org/project/bitmart-python-sdk-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


[BitMart Exchange official](https://bitmart.com) Go client for the [BitMart Cloud API](http://developer-pro.bitmart.com).




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



Usage
=========================
* An example of a spot trade API
* Replace it with your own API KEY
* Run


### Examples:

#### Spot / Margin Trading Endpoints

<details>

<summary>New Order(v2) (SIGNED)</summary>

```go

package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /spot/v2/submit_order
	Doc: https://developer-pro.bitmart.com/en/spot/#new-order-v2-signed
*/
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

</details>


#### Spot Websocket Endpoints

<details>

<summary>Subscribe Public Channel: Ticker</summary>

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

// https://developer-pro.bitmart.com/en/spot/#public-ticker-channel
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

</details>

<details>

<summary>Subscribe Private Channel: Order Progress</summary>

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

// https://developer-pro.bitmart.com/en/spot/#private-order-progress
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

</details>

#### Futures Trading Endpoints

<details>

<summary>Submit Order (SIGNED)</summary>

```go

package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	POST /contract/private/submit-order
	Doc: https://developer-pro.bitmart.com/en/futures/#submit-order-signed
*/
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

</details>

#### Futures Websocket Endpoints

<details>

<summary>Subscribe Public Channel: Ticker</summary>

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

// https://developer-pro.bitmart.com/en/futures/#public-ticker-channel
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

</details>

<details>

<summary>Subscribe Private Channel: Assets</summary>

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

// https://developer-pro.bitmart.com/en/futures/#private-assets-channel
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

	// Just test, Please do not use in production.
	time.Sleep(60 * time.Second)
}

```

</details>