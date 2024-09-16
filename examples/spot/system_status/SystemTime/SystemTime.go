package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /system/time
Doc: https://developer-pro.bitmart.com/en/spot/#get-system-time
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get system time
	var ac, err = client.GetSystemTime()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}
}
