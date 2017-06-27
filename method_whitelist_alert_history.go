package geoedge

// Request
type whitelistAlertHistoryRequest struct {
	api *api

	AlertHistoryId string
}

func (api *api) WhitelistAlertHistoryRequest(alertHistoryId string) *whitelistAlertHistoryRequest {
	return &whitelistAlertHistoryRequest{
		api:            api,
		AlertHistoryId: alertHistoryId,
	}
}

// Do
func (request *whitelistAlertHistoryRequest) Do() (err error) {
	return request.api.makeRequest("Whitelist Alert History", request.AlertHistoryId, nil, nil, nil)

}
