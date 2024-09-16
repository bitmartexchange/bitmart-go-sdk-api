package main

import (
	"github.com/bitmartexchange/bitmart-go-sdk-api"
	"log"
)

/*
POST /account/v1/withdraw/apply
Doc: https://developer-pro.bitmart.com/en/spot/#withdraw-signed
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

	// Withdraw (SIGNED)
	// Parameters for Withdraw to the blockchain
	if ac, err := client.PostAccountWithdrawApply(bitmart.WithdrawApply{
		Currency:    "USDT-ERC20",
		Amount:      "50",
		Destination: "To Digital Address",
		Address:     "0xe57b69a8776b37860407965B73cdFFBDF*******",
		AddressMemo: "",
	}); err != nil {
		log.Panic(err)
	} else {
		log.Println(ac.Response)
	}

	// Parameters for Withdraw to BitMart account
	if ac2, err2 := client.PostAccountWithdrawApply(bitmart.WithdrawApply{
		Currency: "USDT-ERC20",
		Amount:   "50",
		Type:     1,
		Value:    "876940329",
		AreaCode: "",
	}); err2 != nil {
		log.Panic(err2)
	} else {
		log.Println(ac2.Response)
	}

}
