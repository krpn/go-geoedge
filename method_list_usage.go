package geoedge

// Request
type listUsageRequest struct {
	api *geoedgeApi
}

func (api *geoedgeApi) ListUsageRequest() *listUsageRequest {
	return &listUsageRequest{api: api}
}

// Responce
type listUsageResponce struct {
	Usage []usage `json:"usage"`
}

// Do
func (request *listUsageRequest) Do() (usage []usage, err error) {
	responce := listUsageResponce{}
	err = request.api.makeRequest("List Usage", "", nil, nil, &responce)
	if err != nil {
		return
	}

	usage = responce.Usage
	return
}
