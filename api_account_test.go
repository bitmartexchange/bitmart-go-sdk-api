package bitmart

import (
	"log"
	"testing"
)


// GET https://api-cloud.bitmart.com/account/v1/currencies
func TestGetAccountCurrencies(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountCurrencies()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}


// GET https://api-cloud.bitmart.com/account/v1/wallet
func TestGetAccountWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWallet("1")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/v1/deposit/address
func TestGetAccountDepositAddress(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountDepositAddress("USDT-ERC20")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/v1/withdraw/charge
func TestGetAccountWithdrawCharge(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWithdrawCharge("USDT-ERC20")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/account/v1/withdraw/apply
func TestPostAccountWithdrawApply(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostAccountWithdrawApply(WithdrawApply{
		Currency:    "USDT-ERC20",
		Amount:      "50",
		Destination: "To Digital Address",
		Address:     "0xe57b69a8776b37860407965B73cdFFBDFe668Bb5",
		AddressMemo: "",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/v1/deposit-withdraw/history
func TestPostAccountWithdrawHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetDepositWithdrawHistory(HistoryApply{
		Currency:      "USDT-ERC20",
		OperationType: "withdraw",
		Offset:        1,
		Limit:         10,
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/v1/deposit-withdraw/detail
func TestPostAccountWithdrawDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetDepositWithdrawDetail(12121212)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

