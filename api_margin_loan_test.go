package bitmart

import (
	"log"
	"testing"
)

// POST https://api-cloud.bitmart.com/spot/v1/margin/isolated/borrow
func TestMarginBorrowIsolated(t *testing.T) {
	c := NewTestClient()
	ac, err := c.MarginBorrowIsolated("BTC_USDT", "BTC", "1")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// POST https://api-cloud.bitmart.com/spot/v1/margin/isolated/repay
func TestMarginRepayIsolated(t *testing.T) {
	c := NewTestClient()
	ac, err := c.MarginRepayIsolated("BTC_USDT", "BTC", "1")
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/margin/isolated/borrow_record
func TestGetBorrowRecordIsolated(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetBorrowRecordIsolated("BTC_USDT", "", 0, 0, 50) // find 50 records
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetBorrowRecordIsolated("BTC_USDT", "ES16655123452160qnqR1ce", 0, 0, 0) // find by borrow_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/margin/isolated/repay_record
func TestGetRepaymentRecordIsolated(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetRepaymentRecordIsolated("BTC_USDT", "", "", 0, 0, 50) // find 50 records
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetRepaymentRecordIsolated("BTC_USDT", "be7e0887-5bc9-4775-8e45-567cfa567af7", "", 0, 0, 0) // find by repay_id
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}

// GET https://api-cloud.bitmart.com/spot/v1/margin/isolated/pairs
func TestGetTradingPairBorrowingRateAndAmount(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetTradingPairBorrowingRateAndAmount("") // find all
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ac)
	}

	ab, err := c.GetTradingPairBorrowingRateAndAmount("BTC_USDT") // find a specified trading pair
	if err != nil {
		log.Panic(err)
	} else {
		PrintResponse(ab)
	}
}
