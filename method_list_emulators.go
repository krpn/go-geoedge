package geoedge

// Request
type listEmulatorsRequest struct {
	api *api
}

func (api *api) ListEmulatorsRequest() *listEmulatorsRequest {
	return &listEmulatorsRequest{api: api}
}

// Responce
type listEmulatorsResponce struct {
	Emulators []emulator `json:"emulators"`
}

// Do
func (request *listEmulatorsRequest) Do() (emulators []emulator, err error) {
	responce := listEmulatorsResponce{}
	err = request.api.makeRequest("List Emulators", "", nil, nil, &responce)
	if err != nil {
		return
	}

	emulators = responce.Emulators
	return
}
