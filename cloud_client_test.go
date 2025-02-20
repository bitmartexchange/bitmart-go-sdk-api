package bitmart

import (
	"os"
)

func GetDefaultConfig(url string) *Config {
	var config Config
	config.Url = url
	config.ApiKey = "a0a053a31cca9346e52844ee3f13eaa9867c4a51"
	config.SecretKey = "0c45bc61f75ab7b22f34593af965322acfacce90709fdea5eb176c7ab7a63061"
	config.Memo = "mytest"
	config.TimeoutSecond = 30
	config.Headers = map[string]string{
		"X-Custom-Header1": "HeaderValue1",
		"X-Custom-Header2": "HeaderValue2",
	}

	// Initialize another logger to output to the console with a log level of DEBUG
	config.CustomLogger = NewCustomLogger(DEBUG, os.Stdout)

	return &config
}

func NewTestClient() *CloudClient {
	return NewClient(*GetDefaultConfig(API_URL_PRO))
}

func NewTestFuturesClient() *CloudClient {
	return NewClient(*GetDefaultConfig("https://api-cloud-v2.bitmartgcp-test.com"))
}
