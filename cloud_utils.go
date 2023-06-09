package bitmart

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// UTCTime utc time /** Get a UTC-0 timeZ
func UTCTime() string {
	return Int64ToString(time.Now().UnixNano() / 1e6)
}

// PreHashString pre hash string /** timestamp + "#" + memo + "#" + queryString
func PreHashString(timestamp string, memo string, body string) string {
	return fmt.Sprintf("%s#%s#%s", timestamp, memo, body)
}

// HmacSha256Base64Signer hmac sha256 base64 signer
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// JsonBytesToStrut jsonBytes to struct
func JsonBytesToStrut(jsonBytes []byte, result interface{}) error {
	err := json.Unmarshal(jsonBytes, result)
	return err
}

// Int64ToString int64 to string
func Int64ToString(arg int64) string {
	return strconv.FormatInt(arg, 10)
}

// IntToString int to string
func IntToString(arg int) string {
	return strconv.Itoa(arg)
}

// StringToInt string to int
func StringToInt(arg string) int {
	if arg == "" {
		return 0
	}

	i, e := strconv.Atoi(arg)
	if e != nil {
		return 0
	}
	return i
}

func InterfaceToString(inter interface{}) string {

	switch inter.(type) {

	case string:
		return inter.(string)
	case int:
		return IntToString(inter.(int))
	case int64:
		return Int64ToString(inter.(int64))
	}

	return ""
}

// NewParams create params
func NewParams() map[string]interface{} {
	return make(map[string]interface{})
}

// CreateQueryString create query string
func CreateQueryString(params map[string]interface{}) string {
	if params == nil || len(params) == 0 {
		return ""
	}

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	urlParams := url.Values{}
	for k := range params {
		urlParams.Add(k, InterfaceToString(params[k]))
	}
	return "?" + urlParams.Encode()
}

// Headers set headers
func Headers(request *http.Request, apiKey string, timestamp string, sign string) {
	request.Header.Add(ACCEPT, APPLICATION_JSON)
	request.Header.Add(CONTENT_TYPE, APPLICATION_JSON_UTF8)
	request.Header.Add(USER_AGENT, VERSION)

	if apiKey != "" {
		request.Header.Add(X_BM_KEY, apiKey)
	}

	if sign != "" {
		request.Header.Add(X_BM_SIGN, sign)
	}

	if timestamp != "" {
		request.Header.Add(X_BM_TIMESTAMP, timestamp)
	}
}

// PrintRequest print request
func PrintRequest(request *http.Request, body string) {
	fmt.Println("[" + request.Method + "] url:[" + request.URL.String() + "]")
	fmt.Printf("\tHeader: %s\n", request.Header)
	fmt.Println("\tBody: " + body)
}

// PrintResponse print response
func PrintResponse(response *CloudResponse) {
	fmt.Println("\tResponse: ")
	fmt.Println("\t\tHttpStatus: " + IntToString(response.httpStatus))
	fmt.Println("\t\tBody: " + response.response)
	fmt.Println("\t\tRateLimit:")
	fmt.Printf("\t\t\tReset: %d\n", response.limit.reset)
	fmt.Printf("\t\t\tLimit: %d\n", response.limit.limit)
	fmt.Printf("\t\t\tRemaining: %d\n", response.limit.remaining)
}

// GetResponse get response
func GetResponse(response *CloudResponse) string {
	return response.response
}

// GetHttpStatus get http status
func GetHttpStatus(response *CloudResponse) int {
	return response.httpStatus
}

// GetLimit get limit
func GetLimit(response *CloudResponse) RateLimit {
	return response.limit
}

// CreateChannel create channel
func CreateChannel(channel string, symbol string) string {
	return fmt.Sprintf("%s:%s", channel, symbol)
}

// CreateSubscribeParam create subscribe param
func CreateSubscribeParam(channels []string) ([]byte, error) {
	return json.Marshal(OpParam{
		Op:   "subscribe",
		Args: channels,
	})
}
