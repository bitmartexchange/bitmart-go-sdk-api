package bitmart

import (
	"testing"
)


func TestUTCTime(t *testing.T) {
	println(UTCTime())
}

func TestBuildOrderParams(t *testing.T) {
	params := NewParams()
	params["symbol"] = "BTC_USDT"
	params["price"] = "0.01"
	params["amount"] = "200"
	println(CreateQueryString(params))
}

func TestHmacSha256Base64Signer(t *testing.T) {
	s, _ := HmacSha256Base64Signer("123456", "12")
	println(s)
}
