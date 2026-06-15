package bitmart

import (
	"log"
	"testing"
)

// POST https://api-cloud-v2.bitmart.com/account/contract/sub-account/main/v1/sub-to-main
func TestPostContractSubAccountMainSubToMain(t *testing.T) {
	c := NewTestFuturesClient()
	ac, err := c.PostContractSubAccountMainSubToMain("8e7c2b5a-uuid", "1", "USDT", "subAccountName")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/account/contract/sub-account/main/v1/main-to-sub
func TestPostContractSubAccountMainMainToSub(t *testing.T) {
	c := NewTestFuturesClient()
	ac, err := c.PostContractSubAccountMainMainToSub("8e7c2b5a-uuid", "1", "USDT", "subAccountName")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud-v2.bitmart.com/account/contract/sub-account/sub/v1/sub-to-main
func TestPostContractSubAccountSubSubToMain(t *testing.T) {
	c := NewTestFuturesClient()
	ac, err := c.PostContractSubAccountSubSubToMain("8e7c2b5a-uuid", "1", "USDT")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/account/contract/sub-account/main/v1/wallet
func TestGetContractSubAccountMainWallet(t *testing.T) {
	c := NewTestFuturesClient()
	ac, err := c.GetContractSubAccountMainWallet("subAccountName", map[string]interface{}{
		"currency": "USDT",
	})
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/account/contract/sub-account/main/v1/transfer-list
func TestGetContractSubAccountMainTransferList(t *testing.T) {
	c := NewTestFuturesClient()
	ac, err := c.GetContractSubAccountMainTransferList("subAccountName", 100)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud-v2.bitmart.com/account/contract/sub-account/v1/transfer-history
func TestGetContractSubAccountTransferHistory(t *testing.T) {
	c := NewTestFuturesClient()
	ac, err := c.GetContractSubAccountTransferHistory(100)
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}
