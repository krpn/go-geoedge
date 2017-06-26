package geoedge

// Request
type listLocationsRequest struct {
	api *geoedgeApi
}

func (api *geoedgeApi) ListLocationsRequest() *listLocationsRequest {
	return &listLocationsRequest{api: api}
}

// Responce
type listLocationsResponce struct {
	Locations []location `json:"locations"`
}

// Do
func (request *listLocationsRequest) Do() (locations []location, err error) {
	responce := listLocationsResponce{}
	err = request.api.makeRequest("List Locations", "", nil, nil, &responce)
	if err != nil {
		return
	}

	locations = responce.Locations
	return
}
