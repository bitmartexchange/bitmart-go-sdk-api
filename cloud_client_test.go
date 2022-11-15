package bitmart

func GetDefaultConfig() *Config {
	var config Config

	config.Url = API_URL_PRO
	config.WsUrl = WS_URL
	config.ApiKey = "Your API KEY"
	config.SecretKey = "Your Secret KEY"
	config.Memo = "Your Memo"
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
	config.ApiKey = "Your API KEY"
	config.SecretKey = "Your Secret KEY"
	config.Memo = "Your Memo"
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
