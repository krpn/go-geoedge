package geoedge

// Request
type triggerTypesRequest struct {
	api *geoedgeApi
}

func (api *geoedgeApi) TriggerTypesRequest() *triggerTypesRequest {
	return &triggerTypesRequest{api: api}
}

// Responce
type triggerTypesResponce struct {
	TriggerTypes []triggerType `json:"trigger-types"`
}

// Do
func (request *triggerTypesRequest) Do() (triggerTypes []triggerType, err error) {
	responce := triggerTypesResponce{}
	err = request.api.makeRequest("Trigger Types", "", nil, nil, &responce)
	if err != nil {
		return
	}

	triggerTypes = responce.TriggerTypes
	return
}
