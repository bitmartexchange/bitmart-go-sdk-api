[![Logo](./logo.png)](https://bitmart.com)

BitMart-Go-SDK-API
=========================
<p align="left">
    <a href='#'><img src='https://travis-ci.org/meolu/walle-web.svg?branch=master' alt="Build Status"></a>  
</p>

Go client for the [BitMart Cloud API](http://developer-pro.bitmart.com).



Feature
=========================
- Provides exchange quick trading API
- Easier withdrawal
- Efficiency, higher speeds, and lower latencies
- Priority in development and maintenance
- Dedicated and responsive technical support
- Provide webSocket apis calls

Installation
=========================

* 1.Go 1.12.7 support

* 2.Downloads or updates code's dependencies
```git
go get -u github.com/bitmartexchange/bitmart-go-sdk-api
```


Usage
=========================
* An example of a spot trade API
* Replace it with your own API KEY
* Run

#### API Example
```go
package gotest

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

func main() {

	client := bitmart.NewClient(bitmart.Config{
		Url:"https://api-cloud.bitmart.com", // Ues Https url
		ApiKey:"Your API KEY",
		SecretKey:"Your Secret KEY",
		Memo:"Your Memo",
		TimeoutSecond:10,
		IsPrint:true,
	})

	var ac, err = client.PostSpotSubmitOrder(bitmart.Order{Symbol: TEST_SYMBOL, Side: "buy", Type: "limit", Size: "0.1", Price: "8800", Notional: ""})
	if err != nil {
		log.Panic(err)
	} else {
		bitmart.PrintResponse(ac)
	}

}

```

#### WebSocket Public Channel Example
```go
package gotest
import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
    "fmt"
    "sync"
)

func OnMessage(message string) {
	fmt.Println("------------------------>")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := bitmart.NewWS(bitmart.Config{
                        		WsUrl: "wss://ws-manager-compress.bitmart.com/api?protocol=1.1",
                        		ApiKey:"Your API KEY",
                        		SecretKey:"Your Secret KEY",
                        		Memo:"Your Memo",
                        		TimeoutSecond:10,
                        		IsPrint:true,
                        	})
	_ = ws.Connection(OnMessage)

	channels := []string{
		// public channel
		bitmart.CreateChannel(WS_PUBLIC_SPOT_TICKER, "BTC_USDT"),
	}
	ws.SubscribeWithLogin(channels)


	// Just test, Please do not use in production.
	wg.Wait()
}

```

#### WebSocket Private Channel Example
```go
package gotest
import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
    "fmt"
    "sync"
)

func OnMessage(message string) {
	fmt.Println("------------------------>")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := bitmart.NewWS(bitmart.Config{
                        		WsUrl: "wss://ws-manager-compress.bitmart.com/user?protocol=1.1",
                        		ApiKey:"Your API KEY",
                        		SecretKey:"Your Secret KEY",
                        		Memo:"Your Memo",
                        		TimeoutSecond:10,
                        		IsPrint:true,
                        	})
	_ = ws.Connection(OnMessage)

	channels := []string{
		// private channel
		bitmart.CreateChannel(WS_USER_SPOT_ORDER, "BTC_USDT"),
	}
	ws.SubscribeWithLogin(channels)


	// Just test, Please do not use in production.
	wg.Wait()
}

```

#### Contract WebSocket Public Channel Example
```go
package gotest
import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
    "fmt"
    "sync"
)

func OnMessage(message string) {
	fmt.Println("------------------------>")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := bitmart.NewWSContract(bitmart.Config{
                        		WsUrl: "wss://openapi-ws.bitmart.com/api?protocol=1.1",
                        		ApiKey:"Your API KEY",
                        		SecretKey:"Your Secret KEY",
                        		Memo:"Your Memo",
                        		TimeoutSecond:10,
                        		IsPrint:true,
                        	})
	_ = ws.Connection(OnMessage)

	channels := []string{
      // public channel
      WS_PUBLIC_CONTRACT_TICKER,
      CreateChannel(WS_PUBLIC_CONTRACT_DEPTH20, "BTCUSDT"),
      CreateChannel(WS_PUBLIC_CONTRACT_KLINE_1M, "BTCUSDT"),
	}
	ws.SubscribeWithoutLogin(channels)


	// Just test, Please do not use in production.
	wg.Wait()
}

```

#### Contract WebSocket Private Channel Example
```go
package gotest
import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
    "fmt"
    "sync"
)

func OnMessage(message string) {
	fmt.Println("------------------------>")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ws := bitmart.NewWSContract(bitmart.Config{
                        		WsUrl: "wss://openapi-ws.bitmart.com/user?protocol=1.1",
                        		ApiKey:"Your API KEY",
                        		SecretKey:"Your Secret KEY",
                        		Memo:"Your Memo",
                        		TimeoutSecond:10,
                        		IsPrint:true,
                        	})
	_ = ws.Connection(OnMessage)

	channels := []string{
      // private channel
      WS_USER_CONTRACT_UNICAST,
      WS_USER_CONTRACT_POSITION,
      CreateChannel(WS_USER_CONTRACT_ASSET, "USDT"),
	}
	ws.SubscribeWithLogin(channels)
	
	// Just test, Please do not use in production.
	wg.Wait()
}

```

Release Notes
=========================

###### 2022-01-20
- Update endpoints for Spot
    - <code>/spot/v1/symbols/details</code>Add a new respond parameter trade_status, to show the trading status of a trading pair symbol.

###### 2022-01-18
- websocket public channel address<code>wss://ws-manager-compress.bitmart.com?protocol=1.1</code>will be taken down on 2022-02-28 UTC time,The new address is<code>wss://ws-manager-compress.bitmart.com/api?protocol=1.1</code>

###### 2021-11-24
- New endpoints for Spot
    - <code>/spot/v2/orders</code>Get User Order History V2
    - <code>/spot/v1/batch_orders</code>Batch Order
- Update endpoints for Spot
    - <code>/spot/v1/symbols/kline</code>Add new field 'quote_volume'
    - <code>/spot/v1/symbols/trades</code>Add optional parameter N to return the number of items, the default is up to 50 items
    - <code>/spot/v1/order_detail</code>Add new field 'unfilled_volume'
    - <code>/spot/v1/submit_order</code>The request parameter type added limit_maker and ioc order types
- New endpoints for Account
    - <code>/account/v2/deposit-withdraw/history</code>Get Deposit And Withdraw  History V2
- Update endpoints for Account
    - <code>/account/v1/wallet</code>Remove the account_type,Only respond to currency accounts; you can bring currency parameters (optional)

###### 2021-11-06
- Update endpoints for Spot WebSocket
    - Public-Depth Channel:
        - spot/depth50     50 Level Depth Channel
        - spot/depth100    100 Level Depth Channel
    - User-Trade Channel:
        - Eligible pushes add new orders successfully

###### 2021-01-19
- New endpoints for Spot WebSocket
    - Public - ticket channels
    - Public - K channel
    - Public - trading channels
    - Public - depth channels
    - Login
    - User - Trading Channel


###### 2020-07-16 
- Interface Spot API `Cancel Order` update to v2 version that is `POST https://api-cloud.bitmart.com/spot/v2/cancel_order`
- UserAgent set "BitMart-GO-SDK/1.0.1"
                                                    

###### 2020-09-21
- Interface Spot API `/spot/v1/symbols/book` add `size` parameter, which represents the number of depths

###### 2022-11-8
- New endpoints for Contract Market
  - <code>/contract/public/details</code>Get contract details
  - <code>/contract/public/depth</code>Get contract depth
  - <code>/contract/public/open-interest</code>Get contract open interest
  - <code>/contract/public/funding-rate</code>Get contract funding rate
  - <code>/contract/public/kline</code>Get contract kline
- New endpoints for Contract Account
  - <code>/contract/private/assets-detail</code>Get contract user assets detail
- New endpoints for Contract Trade
  - <code>/contract/private/order</code>Get contract order detail
  - <code>/contract/private/order-history</code>Get contract order history
  - <code>/contract/private/position</code>Get contract position
  - <code>/contract/private/trades</code>Get contract trades
  - <code>/contract/private/submit_order</code>Post contract submit order
  - <code>/contract/private/cancel_order</code>Post contract cancel order
  - <code>/contract/private/cancel_orders</code>Post contract batch cancel orders
- New endpoints for Contract WebSocket
  - contract websocket public channel address<code>wss://openapi-ws.bitmart.com/api?protocol=1.1</code>
  - contract websocket private channel address<code>wss://openapi-ws.bitmart.com/user?protocol=1.1</code>

License
=========================
