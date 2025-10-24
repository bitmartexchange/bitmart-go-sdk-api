package bitmart

import (
	"log"
	"testing"
	"time"
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

	// Test with currencies parameter
	ac2, err2 := c.GetAccountCurrencies(map[string]interface{}{
		"currencies": "BTC,ETH,BMX",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud.bitmart.com/account/v1/wallet
func TestGetSpotAccountWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccountWallet(map[string]interface{}{
		"currency": "USDT",
	}) // find by currency
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetSpotAccountWallet() // find all
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
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
		Address:     "0xe57b69a8776b37860407965B73cdFFBDF*******",
		AddressMemo: "",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.PostAccountWithdrawApply(WithdrawApply{
		Currency: "USDT-ERC20",
		Amount:   "50",
		Type:     1,
		Value:    "876940329",
		AreaCode: "",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}
}

// GET https://api-cloud.bitmart.com/account/v2/deposit-withdraw/history
func TestPostAccountWithdrawHistory(t *testing.T) {
	c := NewTestClient()

	ac, err := c.GetDepositWithdrawHistory("withdraw", 10)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ac2, err2 := c.GetDepositWithdrawHistory("withdraw", 10, map[string]interface{}{
		"currency": "BTC",
	})
	if err2 != nil {
		log.Panic(err2)
	} else {
		PrintResponse(ac2)
	}

	// Test with startTime and endTime parameters
	now := time.Now().Unix()
	ac3, err3 := c.GetDepositWithdrawHistory("withdraw", 10, map[string]interface{}{
		"currency":  "BTC",
		"startTime": int(now-3600) * 1000,
		"endTime":   int(now) * 1000,
	})
	if err3 != nil {
		log.Panic(err3)
	} else {
		PrintResponse(ac3)
	}

}

// GET https://api-cloud.bitmart.com/account/v1/deposit-withdraw/detail
func TestPostAccountWithdrawDetail(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetDepositWithdrawDetail("12312312312")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/margin/isolated/account
func TestGetMarginAccountDetailsIsolated(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountDetailsIsolated() //specified trading pair
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetMarginAccountDetailsIsolated(map[string]interface{}{
		"symbol": "BTC",
	}) //all isolated margin assets
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// POST https://api-cloud.bitmart.com/spot/v1/margin/isolated/transfer
func TestMarginAssetTransfer(t *testing.T) {
	c := NewTestClient()
	ac, err := c.MarginAssetTransfer(MarginAssetTransfer{
		Symbol:   "BTC_USDT",
		Currency: "BTC",
		Amount:   "1",
		Side:     "in",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/user_fee
func TestGetUserFee(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetBasicFeeRate()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/trade_fee
func TestGetActualTradeFeeRate(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetActualTradeFeeRate("BTC_USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/v1/withdraw/address/list
func TestGetAccountWithdrawAddressList(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWithdrawAddressList()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
