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
```go
package gotest

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

func main() {

	client := bitmart.NewClient(bitmart.Config{
		Url:"https://api-cloud.bitmart.com", // Ues Https url
		ApiKey:"",
		SecretKey:"",
		Memo:"",
		TimeoutSecond:30,
		IsPrint:true,
	})

	var ac, err = client.PostSpotSubmitLimitBuyOrder(bitmart.LimitBuyOrder{Symbol: "BTC_USDT", Size: "8800", Price: "0.01"})
	if err != nil {
		log.Panic(err)
	} else {
		bitmart.PrintResponse(ac)
	}

}

```


Release Notes
=========================

###### 2020-07-16 
- Interface Spot API `Cancel Order` update to v2 version that is `POST https://api-cloud.bitmart.com/spot/v2/cancel_order`
- UserAgent set "BitMart-GO-SDK/1.0.1"
                                                    

###### 2020-09-21
- Interface Spot API `/spot/v1/symbols/book` add `size` parameter, which represents the number of depths


License
=========================
