package bitmart

// GetAccountCurrencies /** Get Currencies
func (cloudClient *CloudClient) GetAccountCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_ACCOUNT_CURRENCIES_URL, NONE)
}

// GetSpotAccountWallet /** Get Account Balance (KEYED)
// Parameters:
// - Options: currency - currency, .e.g. "BTC", "ETH", "USDT"
func (cloudClient *CloudClient) GetSpotAccountWallet(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(GET, API_ACCOUNT_WALLET_URL, params, KEYED)
}

// GetAccountDepositAddress /** Deposit Address (KEYED)
func (cloudClient *CloudClient) GetAccountDepositAddress(currency string) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = currency
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_ADDRESS_URL, params, KEYED)
}

// GetAccountWithdrawCharge /** Withdraw Quota (KEYED)
func (cloudClient *CloudClient) GetAccountWithdrawCharge(currency string) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = currency
	return cloudClient.requestWithParams(GET, API_ACCOUNT_WITHDRAW_CHARGE_URL, params, KEYED)
}

// WithdrawApply Withdraw Parameters
type WithdrawApply struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`

	Destination string `json:"destination"`  // -To Digital Address
	Address     string `json:"address"`      // Withdraw address (only the address added on the official website is supported)
	AddressMemo string `json:"address_memo"` // Address tag(tag Or payment_id Or memo)

	Type     int    `json:"type"`     // Account type 1=CID 2=Email 3=Phone
	Value    string `json:"value"`    // Account
	AreaCode string `json:"areaCode"` // Phone area code, required when account type is phone, e.g.: 61
}

// PostAccountWithdrawApply /** Withdraw (SIGNED)
func (cloudClient *CloudClient) PostAccountWithdrawApply(apply WithdrawApply) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = apply.Currency
	params["amount"] = apply.Amount
	if apply.Destination != "" {
		params["destination"] = apply.Destination
	}
	if apply.Address != "" {
		params["address"] = apply.Address
	}
	if apply.AddressMemo != "" {
		params["address_memo"] = apply.AddressMemo
	}
	if apply.Type != 0 {
		params["type"] = apply.Type
	}
	if apply.Value != "" {
		params["value"] = apply.Value
	}
	if apply.AreaCode != "" {
		params["areaCode"] = apply.AreaCode
	}

	return cloudClient.requestWithParams(POST, API_ACCOUNT_WITHDRAW_APPLY_URL, params, SIGNED)
}

// GetDepositWithdrawHistory /** Get Deposit And Withdraw History (KEYED)
// Parameters:
// - operationType: type -deposit=deposit -withdraw=withdraw
// - n: Recent N records (value range 1-100)
// - Options: currency - Token symbol, e.g., 'BTC'
func (cloudClient *CloudClient) GetDepositWithdrawHistory(operationType string, n int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["operation_type"] = operationType
	params["N"] = n
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_WITHDRAW_HISTORY_URL, params, KEYED)
}

// GetDepositWithdrawDetail /** Get A Deposit Or Withdraw Detail (KEYED)
func (cloudClient *CloudClient) GetDepositWithdrawDetail(id string) (*CloudResponse, error) {
	params := NewParams()
	params["id"] = id
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_WITHDRAW_DETAIL_URL, params, KEYED)
}

// GetMarginAccountDetailsIsolated /** Get Margin Account Details(Isolated) (KEYED)
// Parameters:
// - Options: symbol - Trading pair (e.g. BMX_USDT), no symbol is passed, and all isolated margin assets are returned
func (cloudClient *CloudClient) GetMarginAccountDetailsIsolated(options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	return cloudClient.requestWithParams(GET, API_SPOT_MARGIN_ACCOUNT_ISOLATED_URL, params, KEYED)
}

type MarginAssetTransfer struct {
	Symbol   string `json:"symbol"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Side     string `json:"side"`
}

// MarginAssetTransfer /** Margin Asset Transfer (SIGNED)
func (cloudClient *CloudClient) MarginAssetTransfer(transfer MarginAssetTransfer) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = transfer.Symbol
	params["currency"] = transfer.Currency
	params["amount"] = transfer.Amount
	params["side"] = transfer.Side
	return cloudClient.requestWithParams(POST, API_SPOT_MARGIN_ASSET_TRANSFER_URL, params, SIGNED)
}

// GetBasicFeeRate /** Get Basic Fee Rate (KEYED)
func (cloudClient *CloudClient) GetBasicFeeRate() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_USER_FEE_URL, KEYED)
}

// GetActualTradeFeeRate /** Get Actual Trade Fee Rate (KEYED)
func (cloudClient *CloudClient) GetActualTradeFeeRate(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_TRADE_FEE_URL, params, KEYED)
}
