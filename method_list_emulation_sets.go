package geoedge

// Request
type listEmulationSetsRequest struct {
	api *api
}

func (api *api) ListEmulationSetsRequest() *listEmulationSetsRequest {
	return &listEmulationSetsRequest{api: api}
}

// Responce
type listEmulationSetsResponce struct {
	EmulationSets []emulationSet `json:"emulation-sets"`
}

// Do
func (request *listEmulationSetsRequest) Do() (emulationSets []emulationSet, err error) {
	responce := listEmulationSetsResponce{}
	err = request.api.makeRequest("List Emulation Sets", "", nil, nil, &responce)
	if err != nil {
		return
	}

	emulationSets = responce.EmulationSets
	return
}
