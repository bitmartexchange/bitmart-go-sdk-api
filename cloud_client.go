package bitmart

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type CloudClient struct {
	Config     Config
	HttpClient *http.Client
}

type ApiMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewClient /** Get a http client
func NewClient(config Config) *CloudClient {
	var client CloudClient
	client.Config = config

	url := config.Url
	if len(url) <= 0 {
		client.Config.Url = API_URL_PRO
	}

	timeout := config.TimeoutSecond
	if timeout <= 0 {
		timeout = 10
	}

	client.HttpClient = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return &client
}

// requestWithoutParams /** Send a GET http request without params
func (cloudClient *CloudClient) requestWithoutParams(method string, path string, auth Auth) (*CloudResponse, error) {
	var cloudResponse CloudResponse

	if _, err := cloudClient.Request(method, path, nil, auth, &cloudResponse); err != nil {
		return nil, err
	}

	return &cloudResponse, nil
}

// requestWithParams /** Send a GET http request with params
func (cloudClient *CloudClient) requestWithParams(method string, path string, params map[string]interface{}, auth Auth) (*CloudResponse, error) {
	var cloudResponse CloudResponse

	if _, err := cloudClient.Request(method, path, params, auth, &cloudResponse); err != nil {
		return nil, err
	}

	return &cloudResponse, nil
}

// Request /** Send a http request to remote server and get a response data
func (cloudClient *CloudClient) Request(method string, requestPath string, params map[string]interface{}, auth Auth, cloudResponse *CloudResponse) (response *http.Response, err error) {
	config := cloudClient.Config
	// uri
	endpoint := config.Url
	if strings.HasSuffix(config.Url, "/") {
		endpoint = config.Url[0 : len(config.Url)-1]
	}

	url := endpoint + requestPath

	// set queryString
	var binBody = bytes.NewReader(make([]byte, 0))
	var jsonBody string
	if GET == method || DELETE == method {
		url = url + CreateQueryString(params)
	} else {

		// get json and bin styles request body
		if params != nil {
			jsonBody, binBody, err = ParseRequestParams(params)
			if err != nil {
				return response, err
			}
		}
	}

	// get a http request
	request, err := http.NewRequest(method, url, binBody)
	if err != nil {
		return response, err
	}

	// set header
	if auth == NONE {
		Headers(request, "", "", "", config.Headers)
	} else if auth == KEYED {
		Headers(request, config.ApiKey, "", "", config.Headers)
	} else if auth == SIGNED {
		timestamp := UTCTime()
		sign, err := HmacSha256Base64Signer(
			PreHashString(timestamp, config.Memo, jsonBody), config.SecretKey)
		if err != nil {
			return response, err
		}
		Headers(request, config.ApiKey, timestamp, sign, config.Headers)
	}

	if config.IsPrint {
		fmt.Println("---------------------------------------------")
		PrintRequest(request, jsonBody)
	}

	// send a request to remote server, and get a response
	response, err = cloudClient.HttpClient.Do(request)
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("ReadCloser Err: %s", err)
		}
	}(response.Body)

	// get a response results and parse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return response, err
	}

	cloudResponse.HttpStatus = response.StatusCode
	cloudResponse.Response = string(body)
	cloudResponse.Limit.Limit = StringToInt(response.Header.Get("X-BM-RateLimit-Limit"))
	cloudResponse.Limit.Remaining = StringToInt(response.Header.Get("X-BM-RateLimit-Remaining"))
	cloudResponse.Limit.Reset = StringToInt(response.Header.Get("X-BM-RateLimit-Reset"))
	cloudResponse.Limit.Mode = response.Header.Get("X-BM-RateLimit-Mode")

	return response, nil
}

// ParseRequestParams /** Parse request params to json and bin styles
func ParseRequestParams(params interface{}) (string, *bytes.Reader, error) {
	if params == nil {
		return "", nil, errors.New("illegal parameter")
	}
	data, err := json.Marshal(params)
	if err != nil {
		return "", nil, errors.New("json convert string error")
	}
	jsonBody := string(data)
	binBody := bytes.NewReader(data)
	return jsonBody, binBody, nil
}
