package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
GET /contract/private/affiliate/rebate-inviteUser
Doc: https://developer.bitmart.com/futures/futures-affiliate/rebate-invite-user
*/
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

	// Get Invited Customer Rebate List (KEYED), cid optional
	var ac, err = client.GetContractAffiliateRebateInviteUser(1735660800, 1740844800, 1, 50, map[string]interface{}{
		"cid": 123456789,
	})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

}
