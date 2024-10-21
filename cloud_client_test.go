package bitmart

import (
	"os"
)

func GetDefaultConfig(url string) *Config {
	var config Config
	config.Url = url
	config.ApiKey = "Your API KEY"
	config.SecretKey = "Your Secret KEY"
	config.Memo = "Your Memo"
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
	return NewClient(*GetDefaultConfig(API_URL_V2_PRO))
}
