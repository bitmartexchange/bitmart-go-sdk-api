package bitmart

// GetAccountCurrencies /** Get Currencies
func (cloudClient *CloudClient) GetAccountCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_ACCOUNT_CURRENCIES_URL, NONE)
}

// GetSpotAccountWallet /** Get Account Balance (KEYED)
func (cloudClient *CloudClient) GetSpotAccountWallet(currency string) (*CloudResponse, error) {
	params := NewParams()
	if currency != "" {
		params["currency"] = currency
	}
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
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Destination string `json:"destination"` // -To Digital Address
	Address     string `json:"address"`
	AddressMemo string `json:"address_memo"`
}

// PostAccountWithdrawApply /** Withdraw (SIGNED)
func (cloudClient *CloudClient) PostAccountWithdrawApply(apply WithdrawApply) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = apply.Currency
	params["amount"] = apply.Amount
	params["destination"] = apply.Destination
	params["address"] = apply.Address
	params["address_memo"] = apply.AddressMemo
	return cloudClient.requestWithParams(POST, API_ACCOUNT_WITHDRAW_APPLY_URL, params, SIGNED)
}

// HistoryApply Query Withdraw/Deposit History Parameters
type HistoryApply struct {
	Currency      string `json:"currency"`
	OperationType string `json:"operation_type"` // type -deposit=deposit -withdraw=withdraw
	N             int    `json:"N"`
}

// GetDepositWithdrawHistory /** Get Deposit And Withdraw History (KEYED)
func (cloudClient *CloudClient) GetDepositWithdrawHistory(history HistoryApply) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = history.Currency
	params["operation_type"] = history.OperationType
	params["N"] = history.N
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_WITHDRAW_HISTORY_URL, params, KEYED)
}

// GetDepositWithdrawDetail /** Get A Deposit Or Withdraw Detail (KEYED)
func (cloudClient *CloudClient) GetDepositWithdrawDetail(id string) (*CloudResponse, error) {
	params := NewParams()
	params["id"] = id
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_WITHDRAW_DETAIL_URL, params, KEYED)
}

// GetMarginAccountDetailsIsolated /** Get Margin Account Details(Isolated) (KEYED)
func (cloudClient *CloudClient) GetMarginAccountDetailsIsolated(symbol string) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}
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
