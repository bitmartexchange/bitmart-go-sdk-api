package bitmart

import "log"

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
