package geoedge

// Request
type listNetworksRequest struct {
	api *geoedgeApi
}

func (api *geoedgeApi) ListNetworksRequest() *listNetworksRequest {
	return &listNetworksRequest{api: api}
}

// Responce
type listNetworksResponce struct {
	Networks []network `json:"networks"`
}

// Do
func (request *listNetworksRequest) Do() (networks []network, err error) {
	responce := listNetworksResponce{}
	err = request.api.makeRequest("List Networks", "", nil, nil, &responce)
	if err != nil {
		return
	}

	networks = responce.Networks
	return
}
