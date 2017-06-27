package geoedge

// Request
type listPlatformsRequest struct {
	api *api
}

func (api *api) ListPlatformsRequest() *listPlatformsRequest {
	return &listPlatformsRequest{api: api}
}

// Responce
type listPlatformsResponce struct {
	Platforms []platrofm `json:"Platforms"`
}

// Do
func (request *listPlatformsRequest) Do() (platforms []platrofm, err error) {
	responce := listPlatformsResponce{}
	err = request.api.makeRequest("List Platforms", "", nil, nil, &responce)
	if err != nil {
		return
	}

	platforms = responce.Platforms
	return
}
