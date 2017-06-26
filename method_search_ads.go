package geoedge

import (
	"strings"
	"time"
)

// Request
type searchAdsRequest struct {
	api *geoedgeApi

	MinDatetime string `json:"min_datetime"`

	MaxDatetime *string `json:"max_datetime,omitempty"`
	Locations   *string `json:"locations,omitempty"`
	ProjectId   *string `json:"project_id,omitempty"`
	ScanId      *string `json:"scan_id,omitempty"`
	FullRaw     *int    `json:"full_raw,omitempty"`
	Offset      *int    `json:"offset,omitempty"`
	Limit       *int    `json:"limit,omitempty"`
}

func (api *geoedgeApi) SearchAdsRequest(minDatetime time.Time, options ...searchAdsRequestOption) *searchAdsRequest {
	request := searchAdsRequest{
		api:         api,
		MinDatetime: minDatetime.Format("2006-01-02 15:04:05"),
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type searchAdsResponce struct {
	PrevPage string              `json:"prev_page"`
	NextPage string              `json:"next_page"`
	AdIds    []map[string]string `json:"Ads"`
}

// Do
func (request *searchAdsRequest) DoExtra() (adIds []string, prevPage, nextPage string, err error) {
	urlParams, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	responce := searchAdsResponce{}
	err = request.api.makeRequest("Search Ads", string(urlParams), nil, nil, &responce)
	if err != nil {
		return
	}

	for _, val := range responce.AdIds {
		adIds = append(adIds, val["id"])
	}
	prevPage, nextPage = responce.PrevPage, responce.NextPage
	return
}

func (request *searchAdsRequest) Do() (adIds []string, err error) {
	adIds, _, _, err = request.DoExtra()
	return
}

// Options
type searchAdsRequestOption func(*searchAdsRequest)

func SearchAdsRequestMaxDatetime(maxDatetime time.Time) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		maxDatetimeStr := maxDatetime.Format("2006-01-02 15:04:05")
		request.MaxDatetime = &maxDatetimeStr
	}
}

func SearchAdsRequestLocations(locations []string) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		locationsStr := strings.Join(locations, ",")
		request.Locations = &locationsStr
	}
}

func SearchAdsRequestProjectId(projectId string) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		request.ProjectId = &projectId
	}
}

func SearchAdsRequestScanId(scanId string) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		request.ScanId = &scanId
	}
}

func SearchAdsRequestFullRaw(fullRaw int) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		request.FullRaw = &fullRaw
	}
}

func SearchAdsRequestOffset(offset int) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		request.Offset = &offset
	}
}

func SearchAdsRequestLimit(limit int) searchAdsRequestOption {
	return func(request *searchAdsRequest) {
		request.Limit = &limit
	}
}
