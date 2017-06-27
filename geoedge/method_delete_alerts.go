package geoedge

import "strings"

// Request
type deleteAlertRequest struct {
	api *api

	AlertIds []string
}

func (api *api) DeleteAlertRequest(alertIds []string) *deleteAlertRequest {
	return &deleteAlertRequest{
		api:      api,
		AlertIds: alertIds,
	}
}

// Do
func (request *deleteAlertRequest) Do() (err error) {
	urlParams := strings.Join(request.AlertIds, ",")
	return request.api.makeRequest("Delete Alerts", urlParams, nil, nil, nil)

}
