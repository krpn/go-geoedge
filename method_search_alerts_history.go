package geoedge

import (
	"strings"
	"time"
)

// Request
type searchAlertsHistoryRequest struct {
	api *api

	AlertId       *string `json:"alert_id,omitempty"`
	ProjectId     *string `json:"project_id,omitempty"`
	TriggerTypeId *string `json:"trigger_type_id,omitempty"`
	MinDatetime   *string `json:"min_datetime,omitempty"`
	MaxDatetime   *string `json:"max_datetime,omitempty"`
	LocationId    *string `json:"location_id,omitempty"`
	FullRaw       *int    `json:"full_raw,omitempty"`
	Offset        *int    `json:"offset,omitempty"`
	Limit         *int    `json:"limit,omitempty"`
}

func (api *api) SearchAlertsHistoryRequest(options ...searchAlertsHistoryRequestOption) *searchAlertsHistoryRequest {
	request := searchAlertsHistoryRequest{api: api}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type searchAlertHistoryResponce struct {
	PrevPage string         `json:"prev_page"`
	NextPage string         `json:"next_page"`
	Alerts   []alertHistory `json:"alerts"`
}

// Do
func (request *searchAlertsHistoryRequest) Do() (alerts []alertHistory, prevPage, nextPage string, err error) {
	urlParams, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	responce := searchAlertHistoryResponce{}
	err = request.api.makeRequest("Search Alerts History", string(urlParams), nil, nil, &responce)
	if err != nil {
		return
	}

	alerts = responce.Alerts
	prevPage, nextPage = responce.PrevPage, responce.NextPage
	return
}

// Options
type searchAlertsHistoryRequestOption func(*searchAlertsHistoryRequest)

func SearchAlertsHistoryRequestAlertId(alertId string) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		request.AlertId = &alertId
	}
}

func SearchAlertsHistoryRequestProjectId(projectId string) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		request.ProjectId = &projectId
	}
}

func SearchAlertsHistoryRequestTriggerTypeId(triggerTypeId string) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		request.TriggerTypeId = &triggerTypeId
	}
}

func SearchAlertsHistoryRequestMinDatetime(minDatetime time.Time) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		minDatetimeStr := minDatetime.Format("2006-01-02 15:04:05")
		request.MinDatetime = &minDatetimeStr
	}
}

func SearchAlertsHistoryRequestMaxDatetime(maxDatetime time.Time) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		maxDatetimeStr := maxDatetime.Format("2006-01-02 15:04:05")
		request.MaxDatetime = &maxDatetimeStr
	}
}

func SearchAlertsHistoryRequestLocationId(locationId []string) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		locationIdStr := strings.Join(locationId, ",")
		request.LocationId = &locationIdStr
	}
}

func SearchAlertsHistoryRequestFullRaw(fullRaw int) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		request.FullRaw = &fullRaw
	}
}

func SearchAlertsHistoryRequestOffset(offset int) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		request.Offset = &offset
	}
}

func SearchAlertsHistoryRequestLimit(limit int) searchAlertsHistoryRequestOption {
	return func(request *searchAlertsHistoryRequest) {
		request.Limit = &limit
	}
}
