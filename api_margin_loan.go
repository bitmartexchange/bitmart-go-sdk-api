package bitmart

// MarginBorrowIsolated /** Margin Borrow (Isolated) (SIGNED)
func (cloudClient *CloudClient) MarginBorrowIsolated(symbol string, currency string, amount string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["currency"] = currency
	params["amount"] = amount
	return cloudClient.requestWithParams(POST, API_MARGIN_BORROW_ISOLATED_URL, params, SIGNED)
}

// MarginRepayIsolated /** Margin Repay (Isolated) (SIGNED)
func (cloudClient *CloudClient) MarginRepayIsolated(symbol string, currency string, amount string) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	params["currency"] = currency
	params["amount"] = amount
	return cloudClient.requestWithParams(POST, API_MARGIN_REPAY_ISOLATED_URL, params, SIGNED)
}

// GetBorrowRecordIsolated /** Get Borrow Record(Isolated) (KEYED)
func (cloudClient *CloudClient) GetBorrowRecordIsolated(symbol string, borrowId string, startTime int64, endTime int64, N int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if borrowId != "" {
		params["borrow_id"] = borrowId
	}
	if startTime != 0 {
		params["start_time"] = startTime
	}
	if endTime != 0 {
		params["end_time"] = endTime
	}
	if N != 0 {
		params["N"] = N
	}
	return cloudClient.requestWithParams(GET, API_BORROW_ROCORD_ISOLATED_URL, params, KEYED)
}

// GetRepaymentRecordIsolated /** Get Repayment Record(Isolated) (KEYED)
func (cloudClient *CloudClient) GetRepaymentRecordIsolated(symbol string, repayId string, currency string, startTime int64, endTime int64, N int) (*CloudResponse, error) {
	params := NewParams()
	params["symbol"] = symbol
	if repayId != "" {
		params["repay_id"] = repayId
	}
	if currency != "" {
		params["currency"] = currency
	}
	if startTime != 0 {
		params["start_time"] = startTime
	}
	if endTime != 0 {
		params["end_time"] = endTime
	}
	if N != 0 {
		params["N"] = N
	}
	return cloudClient.requestWithParams(GET, API_REPAYMENT_ROCORD_ISOLATED_URL, params, KEYED)
}

// GetTradingPairBorrowingRateAndAmount /** Get Trading Pair Borrowing Rate and Amount (KEYED)
func (cloudClient *CloudClient) GetTradingPairBorrowingRateAndAmount(symbol string) (*CloudResponse, error) {
	params := NewParams()
	if symbol != "" {
		params["symbol"] = symbol
	}
	return cloudClient.requestWithParams(GET, API_TRADING_PAIR_BORROWING_RATE_AND_AMOUNT, params, KEYED)
}
