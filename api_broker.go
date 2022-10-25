package bitmart

// broker rebate
func (cloudClient *CloudClient) GetBrokerRebate() (*CloudResponse, error) {
	return cloudClient.requestWithoutParams(GET, API_BROKER_REBATE_URL, KEYED)
}

func (cloudClient *CloudClient) GetBrokerRebateByTimestamp(startTime int64, endTime int64) (*CloudResponse, error) {
	params := NewParams()
	params["start_time"] = startTime
	params["end_time"] = endTime
	return cloudClient.requestWithParams(GET, API_BROKER_REBATE_URL, params, KEYED)
}
