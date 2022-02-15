package bitmart

type Config struct {
	Url           string
	WsUrl         string
	ApiKey        string
	SecretKey     string
	Memo          string
	TimeoutSecond int
	IsPrint       bool
}

type CloudResponse struct {
	httpStatus int
	response   string
	limit      RateLimit
}

type RateLimit struct {
	remaining int
	limit     int
	reset     int
}

type Auth int

const (
	NONE   Auth = 0
	KEYED  Auth = 1
	SIGNED Auth = 2
)

const (
	API_URL_PRO = "https://api-cloud.bitmart.com"
	WS_URL      = "wss://ws-manager-compress.bitmart.com/api?protocol=1.1"  // ws-public
	WS_URL_USER = "wss://ws-manager-compress.bitmart.com/user?protocol=1.1" // ws-private

	X_BM_KEY       = "X-BM-KEY"
	X_BM_SIGN      = "X-BM-SIGN"
	X_BM_TIMESTAMP = "X-BM-TIMESTAMP"

	CONTENT_TYPE = "Content-Type"
	ACCEPT       = "Accept"
	USER_AGENT   = "User-Agent"
	VERSION      = "BitMart-GO-SDK/1.0.1"

	APPLICATION_JSON      = "application/json"
	APPLICATION_JSON_UTF8 = "application/json; charset=UTF-8"

	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

	// system url
	API_SYSTEM_TIME_URL    = "/system/time"
	API_SYSTEM_SERVICE_URL = "/system/service"

	// account url
	API_ACCOUNT_CURRENCIES_URL               = "/account/v1/currencies"
	API_ACCOUNT_WALLET_URL                   = "/account/v1/wallet"
	API_ACCOUNT_DEPOSIT_ADDRESS_URL          = "/account/v1/deposit/address"
	API_ACCOUNT_WITHDRAW_CHARGE_URL          = "/account/v1/withdraw/charge"
	API_ACCOUNT_WITHDRAW_APPLY_URL           = "/account/v1/withdraw/apply"
	API_ACCOUNT_DEPOSIT_WITHDRAW_HISTORY_URL = "/account/v2/deposit-withdraw/history"
	API_ACCOUNT_DEPOSIT_WITHDRAW_DETAIL_URL  = "/account/v1/deposit-withdraw/detail"

	// spot url
	API_SPOT_CURRENCIES_URL      = "/spot/v1/currencies"
	API_SPOT_SYMBOLS_URL         = "/spot/v1/symbols"
	API_SPOT_SYMBOLS_DETAILS_URL = "/spot/v1/symbols/details"
	API_SPOT_TICKER_URL          = "/spot/v1/ticker"
	API_SPOT_STEPS_URL           = "/spot/v1/steps"
	API_SPOT_SYMBOLS_KLINE_URL   = "/spot/v1/symbols/kline"
	API_SPOT_SYMBOLS_BOOK_URL    = "/spot/v1/symbols/book"
	API_SPOT_SYMBOLS_TRADES_URL  = "/spot/v1/symbols/trades"
	API_SPOT_WALLET_URL          = "/spot/v1/wallet"
	API_SPOT_SUBMIT_ORDER_URL    = "/spot/v1/submit_order"
	API_SPOT_BATCH_ORDERS_URL    = "/spot/v1/batch_orders"
	API_SPOT_CANCEL_ORDER_URL    = "/spot/v2/cancel_order"
	API_SPOT_CANCEL_ORDERS_URL   = "/spot/v1/cancel_orders"
	API_SPOT_ORDER_DETAIL_URL    = "/spot/v1/order_detail"
	API_SPOT_ORDERS_URL          = "/spot/v2/orders"
	API_SPOT_TRADES_URL          = "/spot/v1/trades"

	// contract url
	API_CONTRACT_TICKERS_URL = "/contract/v1/tickers"

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

	// spot user
	WS_USER_SPOT_ORDER = "spot/user/order"
)
