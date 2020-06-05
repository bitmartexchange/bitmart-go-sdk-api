package bitmart


func (cloudClient *CloudClient) getSystemTime() (*CloudResponse, error) {
	var cloudResponse CloudResponse

	if _, err := cloudClient.Request(GET, API_SYSTEM_TIME_URL, nil, NONE, &cloudResponse); err != nil {
		return nil, err
	}

	return &cloudResponse, nil
}

func (cloudClient *CloudClient) getSystemService() (*CloudResponse, error) {
	var cloudResponse CloudResponse

	if _, err := cloudClient.Request(GET, API_SYSTEM_SERVICE_URL, nil, NONE, &cloudResponse); err != nil {
		return nil, err
	}

	return &cloudResponse, nil
}


