package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
	GET /spot/v1/steps
	Doc: https://developer-pro.bitmart.com/en/spot/#get-k-line-step
*/
func main() {
	client := bitmart.NewClient(bitmart.Config{TimeoutSecond: 5})

	// Get K-Line Step
	var ac, err = client.GetSpotSteps()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(bitmart.GetResponse(ac))
	}

}
