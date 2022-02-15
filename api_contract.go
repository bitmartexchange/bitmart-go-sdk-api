package bitmart

// contracts

func (cloudClient *CloudClient) GetContractTickers(contractSymbol string) (*CloudResponse, error) {
	params := NewParams()
	params["contract_symbol"] = contractSymbol
	return cloudClient.requestWithParams(GET, API_CONTRACT_TICKERS_URL, params, NONE)
}
