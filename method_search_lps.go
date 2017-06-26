package geoedge

import (
	"strings"
	"time"
)

// Request
type searchLpsRequest struct {
	api *geoedgeApi

	MinDatetime string `json:"min_datetime"`

	MaxDatetime *string `json:"max_datetime,omitempty"`
	Locations   *string `json:"locations,omitempty"`
	ProjectId   *string `json:"project_id,omitempty"`
	ScanId      *string `json:"scan_id,omitempty"`
	Type        *int    `json:"type,omitempty"`
	FullRaw     *int    `json:"full_raw,omitempty"`
	Offset      *int    `json:"offset,omitempty"`
	Limit       *int    `json:"limit,omitempty"`
}

func (api *geoedgeApi) SearchLpsRequest(minDatetime time.Time, options ...searchLpsRequestOption) *searchLpsRequest {
	request := searchLpsRequest{
		api:         api,
		MinDatetime: minDatetime.Format("2006-01-02 15:04:05"),
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type searchLpsResponce struct {
	PrevPage string              `json:"prev_page"`
	NextPage string              `json:"next_page"`
	AdIds    []map[string]string `json:"ads"`
}

// Do
func (request *searchLpsRequest) DoExtra() (adIds []string, prevPage, nextPage string, err error) {
	urlParams, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	responce := searchLpsResponce{}
	err = request.api.makeRequest("Search LPs", string(urlParams), nil, nil, &responce)
	if err != nil {
		return
	}

	for _, val := range responce.AdIds {
		adIds = append(adIds, val["id"])
	}
	prevPage, nextPage = responce.PrevPage, responce.NextPage
	return
}

func (request *searchLpsRequest) Do() (adIds []string, err error) {
	adIds, _, _, err = request.DoExtra()
	return
}

// Options
type searchLpsRequestOption func(*searchLpsRequest)

func SearchLpsRequestMaxDatetime(maxDatetime time.Time) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		maxDatetimeStr := maxDatetime.Format("2006-01-02 15:04:05")
		request.MaxDatetime = &maxDatetimeStr
	}
}

func SearchLpsRequestLocations(locations []string) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		locationsStr := strings.Join(locations, ",")
		request.Locations = &locationsStr
	}
}

func SearchLpsRequestProjectId(projectId string) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		request.ProjectId = &projectId
	}
}

func SearchLpsRequestScanId(scanId string) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		request.ScanId = &scanId
	}
}

func SearchLpsRequestType(typ int) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		request.Type = &typ
	}
}

func SearchLpsRequestFullRaw(fullRaw int) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		request.FullRaw = &fullRaw
	}
}

func SearchLpsRequestOffset(offset int) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		request.Offset = &offset
	}
}

func SearchLpsRequestLimit(limit int) searchLpsRequestOption {
	return func(request *searchLpsRequest) {
		request.Limit = &limit
	}
}
