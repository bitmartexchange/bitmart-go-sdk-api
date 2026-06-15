package bitmart

// Contract (Futures) Sub-Account endpoints.
// Doc: https://developer.bitmart.com/futures (Sub-Account section)
// Note: these require Config.Url = API_URL_V2_PRO.

// PostContractSubAccountMainSubToMain /** Sub-Account to Main-Account (For Main Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency (currently only USDT is supported)
// - subAccount: Sub-account username
func (cloudClient *CloudClient) PostContractSubAccountMainSubToMain(requestNo string, amount string, currency string, subAccount string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	params["subAccount"] = subAccount
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_CONTRACT_SUB_TO_MAIN_MAIN_URL, params, SIGNED)
}

// PostContractSubAccountMainMainToSub /** Main-Account to Sub-Account (For Main Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency (currently only USDT is supported)
// - subAccount: Sub-account username
func (cloudClient *CloudClient) PostContractSubAccountMainMainToSub(requestNo string, amount string, currency string, subAccount string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	params["subAccount"] = subAccount
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_CONTRACT_MAIN_TO_SUB_URL, params, SIGNED)
}

// PostContractSubAccountSubSubToMain /** Sub-Account to Main-Account (For Sub-Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency (currently only USDT is supported)
func (cloudClient *CloudClient) PostContractSubAccountSubSubToMain(requestNo string, amount string, currency string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_CONTRACT_SUB_TO_MAIN_SUB_URL, params, SIGNED)
}

// GetContractSubAccountMainWallet /** Get Sub-Account Contract Wallet Balance (For Main Account) (KEYED)
// Parameters:
// - subAccount: Sub-account username
// - Options.currency: Currency
func (cloudClient *CloudClient) GetContractSubAccountMainWallet(subAccount string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["subAccount"] = subAccount
	return cloudClient.requestWithParams(GET, API_ACCOUNT_SUB_CONTRACT_WALLET_URL, params, KEYED)
}

// GetContractSubAccountMainTransferList /** Get Sub-Account Transfer History (For Main Account) (KEYED)
// Parameters:
// - subAccount: Sub-account username
// - limit: The most recent N records, range [1,100]
func (cloudClient *CloudClient) GetContractSubAccountMainTransferList(subAccount string, limit int) (*CloudResponse, error) {
	params := NewParams()
	params["subAccount"] = subAccount
	params["limit"] = limit
	return cloudClient.requestWithParams(GET, API_ACCOUNT_SUB_CONTRACT_TRANSFER_LIST_URL, params, KEYED)
}

// GetContractSubAccountTransferHistory /** Get Account Transfer History (For Main and Sub-Account) (KEYED)
// Parameters:
// - limit: The most recent N records, range [1,100]
func (cloudClient *CloudClient) GetContractSubAccountTransferHistory(limit int) (*CloudResponse, error) {
	params := NewParams()
	params["limit"] = limit
	return cloudClient.requestWithParams(GET, API_ACCOUNT_SUB_CONTRACT_TRANSFER_HISTORY_URL, params, KEYED)
}
