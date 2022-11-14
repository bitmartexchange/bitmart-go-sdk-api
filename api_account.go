package bitmart

// account/v1/currencies
func (cloudClient *CloudClient) GetAccountCurrencies() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_ACCOUNT_CURRENCIES_URL, NONE)
}

// wallet
func (cloudClient *CloudClient) GetAccountWallet(currency string) (*CloudResponse, error) {
	params := NewParams()
	if currency != "" {
		params["currency"] = currency
	}
	return cloudClient.requestWithParams(GET, API_ACCOUNT_WALLET_URL, params, KEYED)
}

// deposit/address
func (cloudClient *CloudClient) GetAccountDepositAddress(currency string) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = currency
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_ADDRESS_URL, params, KEYED)
}

// withdraw/charge
func (cloudClient *CloudClient) GetAccountWithdrawCharge(currency string) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = currency
	return cloudClient.requestWithParams(GET, API_ACCOUNT_WITHDRAW_CHARGE_URL, params, KEYED)
}

type WithdrawApply struct {
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Destination string `json:"destination"`
	Address     string `json:"address"`
	AddressMemo string `json:"address_memo"`
}

// withdraw/apply
func (cloudClient *CloudClient) PostAccountWithdrawApply(apply WithdrawApply) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = apply.Currency
	params["amount"] = apply.Amount
	params["destination"] = apply.Destination
	params["address"] = apply.Address
	params["address_memo"] = apply.AddressMemo
	return cloudClient.requestWithParams(POST, API_ACCOUNT_WITHDRAW_APPLY_URL, params, SIGNED)
}

type HistoryApply struct {
	Currency      string `json:"currency"`
	OperationType string `json:"operation_type"`
	N             int    `json:"N"`
}

// deposit-withdraw/history
func (cloudClient *CloudClient) GetDepositWithdrawHistory(history HistoryApply) (*CloudResponse, error) {
	params := NewParams()
	params["currency"] = history.Currency
	params["operation_type"] = history.OperationType
	params["N"] = history.N
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_WITHDRAW_HISTORY_URL, params, KEYED)
}

// deposit-withdraw/detail
func (cloudClient *CloudClient) GetDepositWithdrawDetail(id int64) (*CloudResponse, error) {
	params := NewParams()
	params["id"] = id
	return cloudClient.requestWithParams(GET, API_ACCOUNT_DEPOSIT_WITHDRAW_DETAIL_URL, params, KEYED)
}

// margin/isolated/account
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

// margin/isolated/transfer
func (cloudClient *CloudClient) MarginAssetTransfer(transfer MarginAssetTransfer) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = transfer.Symbol
	params["currency"] = transfer.Currency
	params["amount"] = transfer.Amount
	params["side"] = transfer.Side
	return cloudClient.requestWithParams(POST, API_SPOT_MARGIN_ASSET_TRANSFER_URL, params, SIGNED)
}

// user_fee
func (cloudClient *CloudClient) GetBasicFeeRate() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_SPOT_USER_FEE_URL, KEYED)
}

// trade_fee
func (cloudClient *CloudClient) GetActualTradeFeeRate(symbol string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	return cloudClient.requestWithParams(GET, API_SPOT_TRADE_FEE_URL, params, KEYED)
}
