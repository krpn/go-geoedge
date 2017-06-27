package geoedge

// Request
type getAlertRequest struct {
	api *api

	AlertId string
}

func (api *api) GetAlertRequest(AlertId string) *getAlertRequest {
	return &getAlertRequest{
		api:     api,
		AlertId: AlertId,
	}
}

// Responce
type getAlertResponce struct {
	Alerts []alert `json:"alerts"`
}

// Do
func (request *getAlertRequest) Do() (alerts []alert, err error) {
	responce := getAlertResponce{}
	err = request.api.makeRequest("Get Alert", request.AlertId, nil, nil, &responce)
	if err != nil {
		return
	}

	alerts = responce.Alerts
	return
}
