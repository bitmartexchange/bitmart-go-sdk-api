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


Release Notes
=========================

###### 2022-11-03
- New endpoints for API Broker
  - <code>/spot/v1/broker/rebate</code>Applicable to query API Broker's rebate records

###### 2022-10-20
- Upgrade endpoints for Spot
  - <code>/spot/v1/ticker</code> has been upgraded to <code>/spot/v2/ticker</code> and <code>/spot/v1/ticker_detail</code>
  - <code>/spot/v1/submit_order</code> has been upgraded to <code>/spot/v2/submit_order</code>
  - <code>/spot/v1/batch_orders</code> has been upgraded to <code>/spot/v2/batch_orders</code>
  - <code>/spot/v2/cancel_order</code> has been upgraded to <code>/spot/v3/cancel_order</code>
  - <code>/spot/v1/order_detail</code> has been upgraded to <code>/spot/v2/order_detail</code>
  - <code>/spot/v2/orders</code> has been upgraded to <code>/spot/v3/orders</code>
  - <code>/spot/v1/trades</code> has been upgraded to <code>/spot/v2/trades</code>
- New endpoints for Spot & Margin
  - <code>/spot/v1/margin/isolated/account</code>Applicable for isolated margin account inquiries
  - <code>/spot/v1/margin/isolated/transfer</code>For fund transfers between a margin account and spot account
  - <code>/spot/v1/user_fee</code>For querying the base rate of the current user
  - <code>/spot/v1/trade_fee</code>For the actual fee rate of the trading pairs
  - <code>/spot/v1/margin/submit_order</code>Applicable for margin order placement
  - <code>/spot/v1/margin/isolated/borrow</code>Applicable to isolated margin account borrowing operations
  - <code>/spot/v1/margin/isolated/repay</code>Applicable to isolated margin account repayment operations
  - <code>/spot/v1/margin/isolated/borrow_record</code>Applicable to the inquiry of borrowing records of an isolated margin account
  - <code>/spot/v1/margin/isolated/repay_record</code>Applicable to the inquiry of repayment records of isolated margin account
  - <code>/spot/v1/margin/isolated/pairs</code>Applicable for checking the borrowing rate and borrowing amount of trading pairs
  
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


License
=========================
