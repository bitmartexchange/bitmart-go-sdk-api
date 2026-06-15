package bitmart

// Spot Sub-Account endpoints.
// Doc: https://developer.bitmart.com/spot (Sub-Account section)

// PostSpotSubAccountMainSubToMain /** Sub-Account to Main-Account (For Main Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency
// - subAccount: Sub-account username
func (cloudClient *CloudClient) PostSpotSubAccountMainSubToMain(requestNo string, amount string, currency string, subAccount string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	params["subAccount"] = subAccount
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_SPOT_SUB_TO_MAIN_MAIN_URL, params, SIGNED)
}

// PostSpotSubAccountSubSubToMain /** Sub-Account to Main-Account (For Sub-Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency
func (cloudClient *CloudClient) PostSpotSubAccountSubSubToMain(requestNo string, amount string, currency string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_SPOT_SUB_TO_MAIN_SUB_URL, params, SIGNED)
}

// PostSpotSubAccountMainMainToSub /** Main-Account to Sub-Account (For Main Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency
// - subAccount: Sub-account username
func (cloudClient *CloudClient) PostSpotSubAccountMainMainToSub(requestNo string, amount string, currency string, subAccount string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	params["subAccount"] = subAccount
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_SPOT_MAIN_TO_SUB_URL, params, SIGNED)
}

// PostSpotSubAccountMainSubToSub /** Sub-Account to Sub-Account (For Main Account) (SIGNED)
// Parameters:
// - requestNo: UUID or other universally unique identifier, max length 64
// - amount: Transfer amount
// - currency: Currency
// - fromAccount: Transfer-out sub-account username
// - toAccount: Transfer-in sub-account username
func (cloudClient *CloudClient) PostSpotSubAccountMainSubToSub(requestNo string, amount string, currency string, fromAccount string, toAccount string) (*CloudResponse, error) {
	params := NewParams()
	params["requestNo"] = requestNo
	params["amount"] = amount
	params["currency"] = currency
	params["fromAccount"] = fromAccount
	params["toAccount"] = toAccount
	return cloudClient.requestWithParams(POST, API_ACCOUNT_SUB_SPOT_SUB_TO_SUB_URL, params, SIGNED)
}

// GetSpotSubAccountMainTransferList /** Get Sub-Account Transfer History (For Main Account) (KEYED)
// Parameters:
// - moveType: Transfer type (e.g. "spot to spot")
// - n: The most recent N records, range [1,100]
// - Options.accountName: Sub-account username (default: query all sub-accounts)
func (cloudClient *CloudClient) GetSpotSubAccountMainTransferList(moveType string, n int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["moveType"] = moveType
	params["N"] = n
	return cloudClient.requestWithParams(GET, API_ACCOUNT_SUB_SPOT_TRANSFER_LIST_URL, params, KEYED)
}

// GetSpotSubAccountTransferHistory /** Get Account Transfer History (For Main and Sub-Account) (KEYED)
// Parameters:
// - moveType: Transfer type (e.g. "spot to spot")
// - n: The most recent N records, range [1,100]
func (cloudClient *CloudClient) GetSpotSubAccountTransferHistory(moveType string, n int) (*CloudResponse, error) {
	params := NewParams()
	params["moveType"] = moveType
	params["N"] = n
	return cloudClient.requestWithParams(GET, API_ACCOUNT_SUB_SPOT_TRANSFER_HISTORY_URL, params, KEYED)
}

// GetSpotSubAccountMainWallet /** Get Sub-Account Spot Wallet Balance (For Main Account) (KEYED)
// Parameters:
// - subAccount: Sub-account username
// - Options.currency: Currency
func (cloudClient *CloudClient) GetSpotSubAccountMainWallet(subAccount string, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["subAccount"] = subAccount
	return cloudClient.requestWithParams(GET, API_ACCOUNT_SUB_SPOT_WALLET_URL, params, KEYED)
}

// GetSpotSubAccountMainList /** Get Sub-Account List (For Main Account) (KEYED)
// No request parameters.
func (cloudClient *CloudClient) GetSpotSubAccountMainList() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_ACCOUNT_SUB_SPOT_LIST_URL, KEYED)
}
