package geoedge

import "strconv"

// Request
type scanStatusRequest struct {
	api *geoedgeApi

	ScanId string
}

func (api *geoedgeApi) ScanStatusRequest(scanId string) *scanStatusRequest {
	return &scanStatusRequest{
		api:    api,
		ScanId: scanId,
	}
}

// Responce
type scanStatusResponceRaw struct {
	Status string `json:"status"`
	// can be:
	// int: -1
	// string "1"
	// so can't parse into int or string
	NumOfAds interface{} `json:"num_of_ads"`
}

type scanStatusResponce struct {
	Status   string `json:"status"`
	NumOfAds int    `json:"num_of_ads"`
}

// Do
func (request *scanStatusRequest) Do() (result scanStatusResponce, err error) {
	resultRaw := scanStatusResponceRaw{}
	err = request.api.makeRequest("Scan Status", request.ScanId, nil, nil, &resultRaw)

	result.Status = resultRaw.Status
	result.NumOfAds = -1

	switch numOfAdsT := resultRaw.NumOfAds.(type) {
	case string:
		numOfAds, err := strconv.Atoi(numOfAdsT)
		if err != nil {
			break
		}
		result.NumOfAds = numOfAds
	case float64:
		result.NumOfAds = int(numOfAdsT)
	case int:
		result.NumOfAds = numOfAdsT
	}

	return
}
