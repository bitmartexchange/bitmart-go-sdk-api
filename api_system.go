package bitmart

// GetSystemTime /** Get System Time (NONE)
func (cloudClient *CloudClient) GetSystemTime() (*CloudResponse, error) {
	var cloudResponse CloudResponse

	if _, err := cloudClient.Request(GET, API_SYSTEM_TIME_URL, nil, NONE, &cloudResponse); err != nil {
		return nil, err
	}

	return &cloudResponse, nil
}

// GetSystemService /** Get System Service Status (NONE)
func (cloudClient *CloudClient) GetSystemService() (*CloudResponse, error) {
	var cloudResponse CloudResponse

	if _, err := cloudClient.Request(GET, API_SYSTEM_SERVICE_URL, nil, NONE, &cloudResponse); err != nil {
		return nil, err
	}

	return &cloudResponse, nil
}
