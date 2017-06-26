package geoedge

// Request
type listAlertsRequest struct {
	api *geoedgeApi
}

func (api *geoedgeApi) ListAlertsRequest() *listAlertsRequest {
	return &listAlertsRequest{api: api}
}

// Responce
type listAlertsResponce struct {
	Alerts []alert `json:"alerts"`
}

// Do
func (request *listAlertsRequest) Do() (alerts []alert, err error) {
	responce := listAlertsResponce{}
	err = request.api.makeRequest("List Alerts", "", nil, nil, &responce)
	if err != nil {
		return
	}

	alerts = responce.Alerts
	return
}
