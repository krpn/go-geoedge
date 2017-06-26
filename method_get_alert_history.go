package geoedge

// Request
type getAlertHistoryRequest struct {
	api *geoedgeApi

	AlertHistoryId string
}

func (api *geoedgeApi) GetAlertHistoryRequest(alertHistoryId string) *getAlertHistoryRequest {
	return &getAlertHistoryRequest{
		api:            api,
		AlertHistoryId: alertHistoryId,
	}
}

// Responce
type getAlertHistoryResponce struct {
	Alerts []alertHistory `json:"alerts"`
}

// Do
func (request *getAlertHistoryRequest) Do() (alerts []alertHistory, err error) {
	responce := getAlertHistoryResponce{}
	err = request.api.makeRequest("Get Alert History", request.AlertHistoryId, nil, nil, &responce)
	if err != nil {
		return
	}

	alerts = responce.Alerts
	return
}
