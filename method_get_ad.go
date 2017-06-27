package geoedge

import (
	"encoding/json"
	"fmt"
)

// Request
type getAdRequest struct {
	api *api

	AdId            string
	ExtraDataParams *int
}

func (api *api) GetAdRequest(adId string, options ...getAdRequestOption) *getAdRequest {
	request := getAdRequest{
		api:  api,
		AdId: adId,
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type getAdResponce struct {
	Ad ad `json:"Ad"`
	getAdExtraData
}

type getAdExtraData struct {
	CaptureRequestsTotalTime json.RawMessage `json:"capture_requests_total_time"` // Because of unstable structure
	CaptureOnloadTime        string          `json:"capture_onload_time"`
	CaptureRequests          json.RawMessage `json:"capture_requests"`
	// Because of unstable structure
	/*
		CaptureRequests []struct{
			Url string `json:"Url"`
			ResponseStatus int `json:"ResponseStatus"`
			ContentType string `json:"ContentType"`
			ContentLength int `json:"ContentLength"`
			JSCodeUrlId int `json:"JSCodeUrlId"`
			Start string `json:"Start"`
			End string `json:"End"`
		} `json:"capture_requests"`
	*/
	AdCalls         string          `json:"ad_calls"`
	VastCallingUrls json.RawMessage `json:"vast_calling_urls"` // unknown structure
}

// Do
func (request *getAdRequest) DoExtra() (ad ad, extraData getAdExtraData, err error) {
	urlParams := request.AdId
	if request.ExtraDataParams != nil {
		urlParams += fmt.Sprintf("?extra_data_params=%v", *request.ExtraDataParams)
	}

	responce := getAdResponce{}
	err = request.api.makeRequest("Get Ad", urlParams, nil, nil, &responce)
	if err != nil {
		return
	}

	ad = responce.Ad
	extraData = responce.getAdExtraData
	return
}

func (request *getAdRequest) Do() (ad ad, err error) {
	ad, _, err = request.DoExtra()
	return
}

// Options
type getAdRequestOption func(*getAdRequest)

func GetAdRequestExtraDataParams(extraDataParams int) getAdRequestOption {
	return func(request *getAdRequest) {
		request.ExtraDataParams = &extraDataParams
	}
}
