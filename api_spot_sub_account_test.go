package bitmart

import (
	"log"
	"testing"
)

// POST https://api-cloud.bitmart.com/account/sub-account/main/v1/sub-to-main
func TestPostSpotSubAccountMainSubToMain(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubAccountMainSubToMain("8e7c2b5a-uuid", "1", "USDT", "subAccountName")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/account/sub-account/sub/v1/sub-to-main
func TestPostSpotSubAccountSubSubToMain(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubAccountSubSubToMain("8e7c2b5a-uuid", "1", "USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/account/sub-account/main/v1/main-to-sub
func TestPostSpotSubAccountMainMainToSub(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubAccountMainMainToSub("8e7c2b5a-uuid", "1", "USDT", "subAccountName")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/account/sub-account/main/v1/sub-to-sub
func TestPostSpotSubAccountMainSubToSub(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostSpotSubAccountMainSubToSub("8e7c2b5a-uuid", "1", "USDT", "fromSub", "toSub")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/sub-account/main/v1/transfer-list
func TestGetSpotSubAccountMainTransferList(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSubAccountMainTransferList("spot to spot", 100, map[string]interface{}{
		"accountName": "subAccountName",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/sub-account/v1/transfer-history
func TestGetSpotSubAccountTransferHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSubAccountTransferHistory("spot to spot", 100)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/sub-account/main/v1/wallet
func TestGetSpotSubAccountMainWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSubAccountMainWallet("subAccountName", map[string]interface{}{
		"currency": "USDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/account/sub-account/main/v1/subaccount-list
func TestGetSpotSubAccountMainList(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotSubAccountMainList()
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
