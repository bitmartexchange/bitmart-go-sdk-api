package bitmart

// Finance (Savings / Earn) endpoints.
// Doc: https://developer.bitmart.com/finance
// All Finance endpoints use the spot host (API_URL_PRO, https://api-cloud.bitmart.com).

// GetFinanceSavingsHoldings earn /** Get Finance Account Savings Holdings (KEYED)
// No request parameters.
func (cloudClient *CloudClient) GetFinanceSavingsHoldings() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_FINANCE_SAVINGS_HOLDINGS_URL, KEYED)
}

// GetFinanceFlexibleProducts saving/product /** Get Flexible Savings Product List (KEYED)
// Parameters:
// - currentPage: Current page number
// - sizePage: Records per page, max 100
// - Options.coinName: Filter by currency name (e.g. "USDT")
func (cloudClient *CloudClient) GetFinanceFlexibleProducts(currentPage int, sizePage int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["currentPage"] = currentPage
	params["sizePage"] = sizePage
	return cloudClient.requestWithParams(GET, API_FINANCE_FLEXIBLE_PRODUCT_URL, params, KEYED)
}

// PostFinanceFlexibleSubscribe saving/subscribe /** Subscribe Flexible Savings (SIGNED)
// Parameters:
// - productId: Product ID
// - amount: Subscribe amount
// - requestNo: Unique request number (20 digits) for idempotency
func (cloudClient *CloudClient) PostFinanceFlexibleSubscribe(productId string, amount string, requestNo string) (*CloudResponse, error) {
	params := NewParams()
	params["productId"] = productId
	params["amount"] = amount
	params["requestNo"] = requestNo
	return cloudClient.requestWithParams(POST, API_FINANCE_FLEXIBLE_SUBSCRIBE_URL, params, SIGNED)
}

// PostFinanceFlexibleRedeem saving/redeem /** Redeem Flexible Savings (SIGNED)
// Parameters:
// - earnId: Savings order ID
// - amount: Redeem amount
// - requestNo: Unique request number (20 digits) for idempotency
func (cloudClient *CloudClient) PostFinanceFlexibleRedeem(earnId string, amount string, requestNo string) (*CloudResponse, error) {
	params := NewParams()
	params["earnId"] = earnId
	params["amount"] = amount
	params["requestNo"] = requestNo
	return cloudClient.requestWithParams(POST, API_FINANCE_FLEXIBLE_REDEEM_URL, params, SIGNED)
}

// GetFinanceFlexibleHoldings saving/earn /** Get Flexible Savings Holdings (KEYED)
// Parameters:
// - currentPage: Current page number
// - sizePage: Records per page, max 100
// - Options.coinName: Filter by currency name
// - Options.productId: Filter by product ID
func (cloudClient *CloudClient) GetFinanceFlexibleHoldings(currentPage int, sizePage int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["currentPage"] = currentPage
	params["sizePage"] = sizePage
	return cloudClient.requestWithParams(GET, API_FINANCE_FLEXIBLE_EARN_URL, params, KEYED)
}

// GetFinanceFlexibleRecords saving/record /** Get Flexible Savings History Records (KEYED)
// Parameters:
// - recordType: Record type -subscribe -redeem -interest
// - currentPage: Current page number
// - sizePage: Records per page, max 100
// - Options.startTime: Start time, timestamp in milliseconds
// - Options.endTime: End time, timestamp in milliseconds
// - Options.coinName: Filter by currency name
func (cloudClient *CloudClient) GetFinanceFlexibleRecords(recordType string, currentPage int, sizePage int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["type"] = recordType
	params["currentPage"] = currentPage
	params["sizePage"] = sizePage
	return cloudClient.requestWithParams(GET, API_FINANCE_FLEXIBLE_RECORD_URL, params, KEYED)
}

// GetFinanceFixedProducts saving/fixed/product /** Get Fixed Savings Product List (KEYED)
// Parameters:
// - currentPage: Current page number
// - sizePage: Records per page, max 100
// - Options.coinName: Filter by currency name
func (cloudClient *CloudClient) GetFinanceFixedProducts(currentPage int, sizePage int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["currentPage"] = currentPage
	params["sizePage"] = sizePage
	return cloudClient.requestWithParams(GET, API_FINANCE_FIXED_PRODUCT_URL, params, KEYED)
}

// PostFinanceFixedSubscribe saving/fixed/subscribe /** Subscribe Fixed Savings (SIGNED)
// Parameters:
// - productId: Product ID
// - amount: Subscribe amount
// - requestNo: Unique request number (20 digits) for idempotency
// - autoSubscribe: Auto-subscribe type -OFF -REINVEST_FLEXIBLE -REINVEST_FIXED
func (cloudClient *CloudClient) PostFinanceFixedSubscribe(productId string, amount string, requestNo string, autoSubscribe string) (*CloudResponse, error) {
	params := NewParams()
	params["productId"] = productId
	params["amount"] = amount
	params["requestNo"] = requestNo
	params["autoSubscribe"] = autoSubscribe
	return cloudClient.requestWithParams(POST, API_FINANCE_FIXED_SUBSCRIBE_URL, params, SIGNED)
}

// GetFinanceFixedHoldings saving/fixed/earn /** Get Fixed Savings Holdings (KEYED)
// Parameters:
// - currentPage: Current page number
// - sizePage: Records per page, max 100
// - Options.coinName: Filter by currency name
// - Options.productId: Filter by product ID
func (cloudClient *CloudClient) GetFinanceFixedHoldings(currentPage int, sizePage int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["currentPage"] = currentPage
	params["sizePage"] = sizePage
	return cloudClient.requestWithParams(GET, API_FINANCE_FIXED_EARN_URL, params, KEYED)
}

// GetFinanceFixedRecords saving/fixed/record /** Get Fixed Savings History Records (KEYED)
// Parameters:
// - recordType: Record type -subscribe -redeem -interest
// - currentPage: Current page number
// - sizePage: Records per page, max 100
// - Options.startTime: Start time, timestamp in milliseconds
// - Options.endTime: End time, timestamp in milliseconds
// - Options.coinName: Filter by currency name
func (cloudClient *CloudClient) GetFinanceFixedRecords(recordType string, currentPage int, sizePage int, options ...map[string]interface{}) (*CloudResponse, error) {
	params := CreateParams(options...)
	params["type"] = recordType
	params["currentPage"] = currentPage
	params["sizePage"] = sizePage
	return cloudClient.requestWithParams(GET, API_FINANCE_FIXED_RECORD_URL, params, KEYED)
}

// PostFinanceFixedRedeem saving/fixed/redeem /** Early Redeem Fixed Savings (SIGNED)
// Parameters:
// - earnId: Savings order ID
// - requestNo: Unique request number (20 digits) for idempotency
func (cloudClient *CloudClient) PostFinanceFixedRedeem(earnId string, requestNo string) (*CloudResponse, error) {
	params := NewParams()
	params["earnId"] = earnId
	params["requestNo"] = requestNo
	return cloudClient.requestWithParams(POST, API_FINANCE_FIXED_REDEEM_URL, params, SIGNED)
}

// PostFinanceFixedModifyAutoRenewal saving/fixed/subscribe/operate /** Modify Fixed Savings Auto-Renewal (SIGNED)
// Parameters:
// - earnId: Savings order ID
// - autoSubscribe: Auto-subscribe type -OFF -REINVEST_FLEXIBLE -REINVEST_FIXED
func (cloudClient *CloudClient) PostFinanceFixedModifyAutoRenewal(earnId string, autoSubscribe string) (*CloudResponse, error) {
	params := NewParams()
	params["earnId"] = earnId
	params["autoSubscribe"] = autoSubscribe
	return cloudClient.requestWithParams(POST, API_FINANCE_FIXED_MODIFY_AUTO_RENEWAL_URL, params, SIGNED)
}

// PostFinanceAutoSavingToggle saving/subscribe/batch/operate /** Enable/Disable Auto Savings (SIGNED)
// Parameters:
// - autoSubscribe: -open -close
func (cloudClient *CloudClient) PostFinanceAutoSavingToggle(autoSubscribe string) (*CloudResponse, error) {
	params := NewParams()
	params["autoSubscribe"] = autoSubscribe
	return cloudClient.requestWithParams(POST, API_FINANCE_AUTO_SAVING_TOGGLE_URL, params, SIGNED)
}

// GetFinanceAutoSavingStatus saving/subscribe/batch /** Get Global Auto Savings Status (KEYED)
// No request parameters.
func (cloudClient *CloudClient) GetFinanceAutoSavingStatus() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_FINANCE_AUTO_SAVING_STATUS_URL, KEYED)
}

// PostFinanceFlexibleAutoSubscribeToggle saving/subscribe/operate /** Enable/Disable Flexible Auto-Subscribe for a Product (SIGNED)
// Parameters:
// - productId: Product ID
// - autoSubscribe: -open -close
func (cloudClient *CloudClient) PostFinanceFlexibleAutoSubscribeToggle(productId string, autoSubscribe string) (*CloudResponse, error) {
	params := NewParams()
	params["productId"] = productId
	params["autoSubscribe"] = autoSubscribe
	return cloudClient.requestWithParams(POST, API_FINANCE_FLEXIBLE_AUTO_SUBSCRIBE_TOGGLE_URL, params, SIGNED)
}

// GetFinanceFlexibleAutoSubscribeStatus saving/subscribe/status /** Get Flexible Auto-Subscribe Status for a Product (KEYED)
// Parameters:
// - productId: Product ID
func (cloudClient *CloudClient) GetFinanceFlexibleAutoSubscribeStatus(productId string) (*CloudResponse, error) {
	params := NewParams()
	params["productId"] = productId
	return cloudClient.requestWithParams(GET, API_FINANCE_FLEXIBLE_AUTO_SUBSCRIBE_STATUS_URL, params, KEYED)
}
