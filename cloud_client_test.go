package bitmart

func GetDefaultConfig() *Config {
	var config Config

	config.Url = API_URL_PRO
	config.WsUrl = WS_URL
	config.ApiKey = "80618e45710812162b04892c7ee5ead4a3cc3e56"
	config.SecretKey = "6c6c98544461bbe71db2bca4c6d7fd0021e0ba9efc215f9c6ad41852df9d9df9"
	config.Memo = "test001"
	config.TimeoutSecond = 30
	config.IsPrint = true
	return &config
}

func NewTestClient() *CloudClient {
	client := NewClient(*GetDefaultConfig())
	return client
}

func NewTestWS() *CloudWS {
	ws := NewWS(*GetDefaultConfig())
	return ws
}

func GetWSPrivateConfig() *Config {
	var config Config

	config.Url = API_URL_PRO
	config.WsUrl = WS_URL_USER
	config.ApiKey = "80618e45710812162b04892c7ee5ead4a3cc3e56"
	config.SecretKey = "6c6c98544461bbe71db2bca4c6d7fd0021e0ba9efc215f9c6ad41852df9d9df9"
	config.Memo = "test001"
	config.TimeoutSecond = 30
	config.IsPrint = true
	return &config
}

func NewTestWSWithLogin() *CloudWS {
	ws := NewWS(*GetWSPrivateConfig())
	return ws
}

func GetContractConfig() *Config {
	config := GetDefaultConfig()
	config.WsUrl = CONTRACT_WS_URL
	return config
}

func GetContractWSPrivateConfig() *Config {
	config := GetDefaultConfig()
	config.WsUrl = CONTRACT_WS_PRIVATE_URL
	return config
}

func NewTestWSContract() *CloudWSContract {
	ws := NewWSContract(*GetContractConfig())
	return ws
}

func NewTestWSContractWithLogin() *CloudWSContract {
	ws := NewWSContract(*GetContractWSPrivateConfig())
	return ws
}
