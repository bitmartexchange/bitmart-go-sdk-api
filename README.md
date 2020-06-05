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
	"log"
	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

func main() {

	client := NewClient(Config{
		Url:"https://api-cloud.bitmart.com",
		ApiKey:"",
		SecretKey:"",
		Memo:"",
		TimeoutSecond:30,
		IsPrint:true,
	})

	ac, err := client.postSpotSubmitLimitBuyOrder(LimitBuyOrder{Symbol:"BTC_USDT", Size:"8800", Price:"0.01"})
	if err != nil {
		log.Panic(err)
	} else {
		printResponse(ac)
	}

}


```


License
=========================
