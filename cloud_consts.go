package bitmart

type Config struct {
	Url           string
	WsUrl         string
	ApiKey        string
	SecretKey     string
	Memo          string
	TimeoutSecond int
	IsPrint       bool
	CustomLogger  *CustomLogger
	Headers       map[string]string
}

type CloudResponse struct {
	HttpStatus int
	Response   string
	Limit      RateLimit
}

type RateLimit struct {
	Remaining int
	Limit     int
	Reset     int
	Mode      string
}

type Auth int

const (
	NONE   Auth = 0
	KEYED  Auth = 1
	SIGNED Auth = 2
)

const (
	API_URL_PRO     = "https://api-cloud.bitmart.com"                           // rest api url
	API_URL_V2_PRO  = "https://api-cloud-v2.bitmart.com"                        // rest api url
	SPOT_WS_URL     = "wss://ws-manager-compress.bitmart.com/api?protocol=1.1"  // spot-ws-public
	SPOT_WS_USER    = "wss://ws-manager-compress.bitmart.com/user?protocol=1.1" // spot-ws-private
	FUTURES_WS_URL  = "wss://openapi-ws-v2.bitmart.com/api?protocol=1.1"        // contract-ws-public
	FUTURES_WS_USER = "wss://openapi-ws-v2.bitmart.com/user?protocol=1.1"       // contract-ws-private

	X_BM_KEY       = "X-BM-KEY"
	X_BM_SIGN      = "X-BM-SIGN"
	X_BM_TIMESTAMP = "X-BM-TIMESTAMP"

	CONTENT_TYPE = "Content-Type"
	ACCEPT       = "Accept"
	USER_AGENT   = "User-Agent"
	VERSION      = "bitmart-go-sdk-api/1.7.0"

	APPLICATION_JSON      = "application/json"
	APPLICATION_JSON_UTF8 = "application/json; charset=UTF-8"

	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

	// API_SYSTEM_TIME_URL System Status Endpoints: https://developer-pro.bitmart.com/en/spot/#system-status
	API_SYSTEM_TIME_URL    = "/system/time"
	API_SYSTEM_SERVICE_URL = "/system/service"

	// API_ACCOUNT_CURRENCIES_URL Funding Account Endpoints: https://developer-pro.bitmart.com/en/spot/#funding-account
	API_ACCOUNT_CURRENCIES_URL               = "/account/v1/currencies"
	API_ACCOUNT_WALLET_URL                   = "/account/v1/wallet"
	API_ACCOUNT_DEPOSIT_ADDRESS_URL          = "/account/v1/deposit/address"
	API_ACCOUNT_WITHDRAW_CHARGE_URL          = "/account/v1/withdraw/charge"
	API_ACCOUNT_WITHDRAW_APPLY_URL           = "/account/v1/withdraw/apply"
	API_ACCOUNT_DEPOSIT_WITHDRAW_HISTORY_URL = "/account/v2/deposit-withdraw/history"
	API_ACCOUNT_DEPOSIT_WITHDRAW_DETAIL_URL  = "/account/v1/deposit-withdraw/detail"
	API_SPOT_MARGIN_ACCOUNT_ISOLATED_URL     = "/spot/v1/margin/isolated/account"
	API_SPOT_MARGIN_ASSET_TRANSFER_URL       = "/spot/v1/margin/isolated/transfer"
	API_SPOT_USER_FEE_URL                    = "/spot/v1/user_fee"
	API_SPOT_TRADE_FEE_URL                   = "/spot/v1/trade_fee"

	// API_SPOT_CURRENCIES_URL Public Market Data Endpoints: https://developer-pro.bitmart.com/en/spot/#public-market-data
	API_SPOT_CURRENCIES_URL      = "/spot/v1/currencies"
	API_SPOT_SYMBOLS_URL         = "/spot/v1/symbols"
	API_SPOT_SYMBOLS_DETAILS_URL = "/spot/v1/symbols/details"

	API_SPOT_V3_TICKERS_URL       = "/spot/quotation/v3/tickers"
	API_SPOT_V3_TICKER_URL        = "/spot/quotation/v3/ticker"
	API_SPOT_V3_LATEST_KLINE_URL  = "/spot/quotation/v3/lite-klines"
	API_SPOT_V3_HISTORY_KLINE_URL = "/spot/quotation/v3/klines"
	API_SPOT_V3_BOOKS_URL         = "/spot/quotation/v3/books"
	API_SPOT_V3_TRADES_URL        = "/spot/quotation/v3/trades"

	// API_SPOT_WALLET_URL Spot / Margin Trading Endpoints: https://developer-pro.bitmart.com/en/spot/#spot-margin-trading
	API_SPOT_WALLET_URL              = "/spot/v1/wallet"
	API_SPOT_SUBMIT_ORDER_URL        = "/spot/v2/submit_order"
	API_SPOT_SUBMIT_MARGIN_ORDER_URL = "/spot/v1/margin/submit_order"
	API_SPOT_BATCH_ORDERS_URL        = "/spot/v4/batch_orders"
	API_SPOT_CANCEL_ORDER_URL        = "/spot/v3/cancel_order"
	API_SPOT_CANCEL_ORDERS_URL       = "/spot/v4/cancel_orders"
	API_SPOT_CANCEL_ALL_URL          = "/spot/v4/cancel_all"

	API_SPOT_V4_QUERY_ORDER_BY_ID_URL     = "/spot/v4/query/order"
	API_SPOT_V4_QUERY_ORDER_BY_CLIENT_URL = "/spot/v4/query/client-order"
	API_SPOT_V4_QUERY_OPEN_ORDERS_URL     = "/spot/v4/query/open-orders"
	API_SPOT_V4_QUERY_HISTORY_ORDERS_URL  = "/spot/v4/query/history-orders"
	API_SPOT_V4_QUERY_TRADES_URL          = "/spot/v4/query/trades"
	API_SPOT_V4_QUERY_ORDER_TRADES_URL    = "/spot/v4/query/order-trades"

	// Spot Algo (Strategy) Order Endpoints: https://developer.bitmart.com/spot
	API_SPOT_ALGO_SUBMIT_ORDER_URL            = "/spot/v4/algo/submit_order"
	API_SPOT_ALGO_CANCEL_ORDER_URL            = "/spot/v4/algo/cancel_order"
	API_SPOT_ALGO_CANCEL_ALL_URL              = "/spot/v4/algo/cancel_all"
	API_SPOT_V4_QUERY_ALGO_ORDER_URL          = "/spot/v4/query/algo/order"
	API_SPOT_V4_QUERY_ALGO_CLIENT_ORDER_URL   = "/spot/v4/query/algo/client-order"
	API_SPOT_V4_QUERY_ALGO_OPEN_ORDERS_URL    = "/spot/v4/query/algo/open-orders"
	API_SPOT_V4_QUERY_ALGO_HISTORY_ORDERS_URL = "/spot/v4/query/algo/history-orders"

	// Margin Loan Endpoints: https://developer-pro.bitmart.com/en/spot/#margin-loan
	API_MARGIN_BORROW_ISOLATED_URL             = "/spot/v1/margin/isolated/borrow"
	API_MARGIN_REPAY_ISOLATED_URL              = "/spot/v1/margin/isolated/repay"
	API_BORROW_ROCORD_ISOLATED_URL             = "/spot/v1/margin/isolated/borrow_record"
	API_REPAYMENT_ROCORD_ISOLATED_URL          = "/spot/v1/margin/isolated/repay_record"
	API_TRADING_PAIR_BORROWING_RATE_AND_AMOUNT = "/spot/v1/margin/isolated/pairs"

	// broker url
	API_BROKER_REBATE_URL = "/spot/v1/broker/rebate"

	// Futures Market Data Endpoints: https://developer-pro.bitmart.com/en/futures/#futures-market-data
	API_CONTRACT_DETAILS_URL              = "/contract/public/details"
	API_CONTRACT_DEPTH_URL                = "/contract/public/depth"
	API_CONTRACT_OPEN_INTEREST_URL        = "/contract/public/open-interest"
	API_CONTRACT_FUNDING_RATE_URL         = "/contract/public/funding-rate"
	API_CONTRACT_FUNDING_RATE_HISTORY_URL = "/contract/public/funding-rate-history"
	API_CONTRACT_KLINE_URL                = "/contract/public/kline"
	API_CONTRACT_MARK_PRICE_KLINE_URL     = "/contract/public/markprice-kline"

	// Futures Account Data Endpoints: https://developer-pro.bitmart.com/en/futures/#futures-account-data
	API_CONTRACT_ASSETS_DETAIL_URL = "/contract/private/assets-detail"

	// Futures Trading Endpoints: https://developer-pro.bitmart.com/en/futures/#futures-trading
	API_CONTRACT_TRADE_FEE_RATE_URL      = "/contract/private/trade-fee-rate"
	API_CONTRACT_ORDER_URL               = "/contract/private/order"
	API_CONTRACT_ORDER_HISTORY_URL       = "/contract/private/order-history"
	API_CONTRACT_OPEN_ORDERS_URL         = "/contract/private/get-open-orders"
	API_CONTRACT_CURRENT_PLAN_ORDER_URL  = "/contract/private/current-plan-order"
	API_CONTRACT_POSITION_URL            = "/contract/private/position"
	API_CONTRACT_POSITION_RISK_URL       = "/contract/private/position-risk"
	API_CONTRACT_TRANSACTION_HISTORY_URL = "/contract/private/transaction-history"
	API_CONTRACT_TRADES_URL              = "/contract/private/trades"
	API_CONTRACT_TRANSFER_LIST_URL       = "/account/v1/transfer-contract-list"
	API_CONTRACT_SUBMIT_ORDER_URL        = "/contract/private/submit-order"
	API_CONTRACT_CANCEL_ORDER_URL        = "/contract/private/cancel-order"
	API_CONTRACT_CANCEL_ORDERS_URL       = "/contract/private/cancel-orders"
	API_CONTRACT_SUBMIT_PLAN_ORDER_URL   = "/contract/private/submit-plan-order"
	API_CONTRACT_CANCEL_PLAN_ORDER_URL   = "/contract/private/cancel-plan-order"
	API_CONTRACT_SUBMIT_TRAIL_ORDER_URL  = "/contract/private/submit-trail-order"
	API_CONTRACT_CANCEL_TRAIL_ORDER_URL  = "/contract/private/cancel-trail-order"

	API_CONTRACT_TRANSFER_URL                 = "/account/v1/transfer-contract"
	API_CONTRACT_SUBMIT_LEVERAGE_URL          = "/contract/private/submit-leverage"
	API_CONTRACT_SUBMIT_TP_SL_ORDER_URL       = "/contract/private/submit-tp-sl-order"
	API_CONTRACT_MODIFY_PLAN_ORDER_URL        = "/contract/private/modify-plan-order"
	API_CONTRACT_MODIFY_PRESET_PLAN_ORDER_URL = "/contract/private/modify-preset-plan-order"
	API_CONTRACT_MODIFY_TP_SL_ORDER_URL       = "/contract/private/modify-tp-sl-order"

	// v1.4.0 New APIs
	API_CONTRACT_MODIFY_LIMIT_ORDER_URL = "/contract/private/modify-limit-order"
	API_CONTRACT_CANCEL_ALL_AFTER_URL   = "/contract/private/cancel-all-after"
	API_CONTRACT_SET_POSITION_MODE_URL  = "/contract/private/set-position-mode"
	API_CONTRACT_GET_POSITION_MODE_URL  = "/contract/private/get-position-mode"
	API_CONTRACT_POSITION_V2_URL        = "/contract/private/position-v2"
	API_CONTRACT_LEVERAGE_BRACKET_URL   = "/contract/public/leverage-bracket"
	API_CONTRACT_MARKET_TRADE_URL       = "/contract/public/market-trade"

	// v1.4.0 Additional APIs
	API_ACCOUNT_WITHDRAW_ADDRESS_LIST_URL = "/account/v1/withdraw/address/list"

	// v1.5.0 New APIs
	// Futures Market Data: https://developer.bitmart.com/futures
	API_CONTRACT_FUNDING_RATE_V2_URL = "/contract/public/funding-rate-v2"
	// Futures Trading (cross collateral): https://developer.bitmart.com/futures
	API_CONTRACT_AUTO_REPAYMENT_URL                = "/contract/private/auto_repayment"
	API_CONTRACT_CROSS_COLLATERAL_INTEREST_LOG_URL = "/contract/private/cross_collateral/interest_log"
	// Futures Demo Trading: https://developer.bitmart.com/futures
	API_CONTRACT_CLAIM_URL = "/contract/private/claim"

	// v1.6.0 New APIs
	// Spot Sub-Account Endpoints: https://developer.bitmart.com/spot
	API_ACCOUNT_SUB_SPOT_SUB_TO_MAIN_MAIN_URL = "/account/sub-account/main/v1/sub-to-main"
	API_ACCOUNT_SUB_SPOT_SUB_TO_MAIN_SUB_URL  = "/account/sub-account/sub/v1/sub-to-main"
	API_ACCOUNT_SUB_SPOT_MAIN_TO_SUB_URL      = "/account/sub-account/main/v1/main-to-sub"
	API_ACCOUNT_SUB_SPOT_SUB_TO_SUB_URL       = "/account/sub-account/main/v1/sub-to-sub"
	API_ACCOUNT_SUB_SPOT_TRANSFER_LIST_URL    = "/account/sub-account/main/v1/transfer-list"
	API_ACCOUNT_SUB_SPOT_TRANSFER_HISTORY_URL = "/account/sub-account/v1/transfer-history"
	API_ACCOUNT_SUB_SPOT_WALLET_URL           = "/account/sub-account/main/v1/wallet"
	API_ACCOUNT_SUB_SPOT_LIST_URL             = "/account/sub-account/main/v1/subaccount-list"
	// Contract Sub-Account Endpoints: https://developer.bitmart.com/futures
	API_ACCOUNT_SUB_CONTRACT_SUB_TO_MAIN_MAIN_URL = "/account/contract/sub-account/main/v1/sub-to-main"
	API_ACCOUNT_SUB_CONTRACT_MAIN_TO_SUB_URL      = "/account/contract/sub-account/main/v1/main-to-sub"
	API_ACCOUNT_SUB_CONTRACT_SUB_TO_MAIN_SUB_URL  = "/account/contract/sub-account/sub/v1/sub-to-main"
	API_ACCOUNT_SUB_CONTRACT_WALLET_URL           = "/account/contract/sub-account/main/v1/wallet"
	API_ACCOUNT_SUB_CONTRACT_TRANSFER_LIST_URL    = "/account/contract/sub-account/main/v1/transfer-list"
	API_ACCOUNT_SUB_CONTRACT_TRANSFER_HISTORY_URL = "/account/contract/sub-account/v1/transfer-history"

	// v1.7.0 New APIs
	// Finance (Savings / Earn) Endpoints: https://developer.bitmart.com/finance
	API_FINANCE_SAVINGS_HOLDINGS_URL               = "/newearn/cloud/v1/earn"
	API_FINANCE_FLEXIBLE_PRODUCT_URL               = "/newearn/cloud/v1/saving/product"
	API_FINANCE_FLEXIBLE_SUBSCRIBE_URL             = "/newearn/cloud/v1/saving/subscribe"
	API_FINANCE_FLEXIBLE_REDEEM_URL                = "/newearn/cloud/v1/saving/redeem"
	API_FINANCE_FLEXIBLE_EARN_URL                  = "/newearn/cloud/v1/saving/earn"
	API_FINANCE_FLEXIBLE_RECORD_URL                = "/newearn/cloud/v1/saving/record"
	API_FINANCE_FIXED_PRODUCT_URL                  = "/newearn/cloud/v1/saving/fixed/product"
	API_FINANCE_FIXED_SUBSCRIBE_URL                = "/newearn/cloud/v1/saving/fixed/subscribe"
	API_FINANCE_FIXED_EARN_URL                     = "/newearn/cloud/v1/saving/fixed/earn"
	API_FINANCE_FIXED_RECORD_URL                   = "/newearn/cloud/v1/saving/fixed/record"
	API_FINANCE_FIXED_REDEEM_URL                   = "/newearn/cloud/v1/saving/fixed/redeem"
	API_FINANCE_FIXED_MODIFY_AUTO_RENEWAL_URL      = "/newearn/cloud/v1/saving/fixed/subscribe/operate"
	API_FINANCE_AUTO_SAVING_TOGGLE_URL             = "/newearn/cloud/v1/saving/subscribe/batch/operate"
	API_FINANCE_AUTO_SAVING_STATUS_URL             = "/newearn/cloud/v1/saving/subscribe/batch"
	API_FINANCE_FLEXIBLE_AUTO_SUBSCRIBE_TOGGLE_URL = "/newearn/cloud/v1/saving/subscribe/operate"
	API_FINANCE_FLEXIBLE_AUTO_SUBSCRIBE_STATUS_URL = "/newearn/cloud/v1/saving/subscribe/status"
	// Futures Affiliate (Broker): https://developer.bitmart.com/futures
	API_CONTRACT_AFFILIATE_REBATE_LIST_URL             = "/contract/private/affiliate/rebate-list"
	API_CONTRACT_AFFILIATE_TRADE_LIST_URL              = "/contract/private/affiliate/trade-list"
	API_CONTRACT_AFFILIATE_REBATE_USER_URL             = "/contract/private/affiliate/rebate-user"
	API_CONTRACT_AFFILIATE_REBATE_API_URL              = "/contract/private/affiliate/rebate-api"
	API_CONTRACT_AFFILIATE_INVITE_CHECK_URL            = "/contract/private/affiliate/invite-check"
	API_CONTRACT_AFFILIATE_REBATE_INVITE_USER_URL      = "/contract/private/affiliate/rebate-inviteUser"
	API_CONTRACT_AFFILIATE_DEPOSIT_WITHDRAWAL_LIST_URL = "/contract/private/affiliate/deposit-withdrawal-list"
	API_CONTRACT_AFFILIATE_CUSTOMER_INFO_URL           = "/contract/private/affiliate/aff-customer-info"

	// web socket
	// spot common
	WS_PUBLIC_SPOT_TICKER     = "spot/ticker"
	WS_PUBLIC_SPOT_TRADE      = "spot/trade"
	WS_PUBLIC_SPOT_DEPTH5     = "spot/depth5"
	WS_PUBLIC_SPOT_DEPTH20    = "spot/depth20"
	WS_PUBLIC_SPOT_DEPTH50    = "spot/depth50"
	WS_PUBLIC_SPOT_KLINE_1M   = "spot/kline1m"
	WS_PUBLIC_SPOT_KLINE_3M   = "spot/kline3m"
	WS_PUBLIC_SPOT_KLINE_5M   = "spot/kline5m"
	WS_PUBLIC_SPOT_KLINE_15M  = "spot/kline15m"
	WS_PUBLIC_SPOT_KLINE_30M  = "spot/kline30m"
	WS_PUBLIC_SPOT_KLINE_1H   = "spot/kline1H"
	WS_PUBLIC_SPOT_KLINE_2H   = "spot/kline2H"
	WS_PUBLIC_SPOT_KLINE_4H   = "spot/kline4H"
	WS_PUBLIC_SPOT_KLINE_1D   = "spot/kline1D"
	WS_PUBLIC_SPOT_KLINE_1W   = "spot/kline1W"
	WS_PUBLIC_SPOT_KLINE_1MON = "spot/kline1M"

	WS_PUBLIC_SPOT_BOOK_TICKER = "spot/bookTicker"

	// spot user
	WS_USER_SPOT_ORDER = "spot/user/order"

	// contract common
	WS_PUBLIC_CONTRACT_TICKER    = "futures/ticker"
	WS_PUBLIC_CONTRACT_DEPTH5    = "futures/depth5"
	WS_PUBLIC_CONTRACT_DEPTH20   = "futures/depth20"
	WS_PUBLIC_CONTRACT_DEPTH50   = "futures/depth50"
	WS_PUBLIC_CONTRACT_KLINE_1M  = "futures/klineBin1m"
	WS_PUBLIC_CONTRACT_KLINE_5M  = "futures/klineBin5m"
	WS_PUBLIC_CONTRACT_KLINE_15M = "futures/klineBin15m"
	WS_PUBLIC_CONTRACT_KLINE_30M = "futures/klineBin30m"
	WS_PUBLIC_CONTRACT_KLINE_1H  = "futures/klineBin1H"
	WS_PUBLIC_CONTRACT_KLINE_2H  = "futures/klineBin2H"
	WS_PUBLIC_CONTRACT_KLINE_4H  = "futures/klineBin4H"
	WS_PUBLIC_CONTRACT_KLINE_1D  = "futures/klineBin1D"
	WS_PUBLIC_CONTRACT_KLINE_1W  = "futures/klineBin1W"

	// contract user
	WS_USER_CONTRACT_ASSET    = "futures/asset"
	WS_USER_CONTRACT_POSITION = "futures/position"
	WS_USER_CONTRACT_UNICAST  = "futures/unicast"
)
