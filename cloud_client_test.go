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

/*
func NewTestClient() *CloudClient {
	client := NewClient(*GetDefaultConfig())
	return client
}
*/

func NewTestWS() *CloudWS {
	ws := NewWS(*GetDefaultConfig())
	return ws
}

func NewTestWSContract() *CloudWSContract {
	ws := NewWSContract(*GetContractConfig())
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
	var config Config

	config.Url = API_URL_TEST
	config.WsUrl = WS_URL_TEST
	config.ApiKey = "9d82a570c5a72cdaf4311b0f6bddae1d8256d5bd"
	config.SecretKey = "23eb8a741a80dd32dad1d4ef0474035ad10417bbca0e8cdad3c738a64c308edf"
	config.Memo = "leon"
	config.TimeoutSecond = 30
	config.IsPrint = true
	return &config
}

func NewTestClient() *CloudClient {
	client := NewClient(*GetContractConfig())
	return client
}

func NewTestWSContractWithLogin() *CloudWSContract {
	ws := NewWSContract(*GetContractConfig())
	return ws
}
