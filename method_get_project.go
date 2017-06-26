package geoedge

// Request
type getProjectRequest struct {
	api *geoedgeApi

	projectId string
}

func (api *geoedgeApi) GetProjectRequest(projectId string) *getProjectRequest {
	return &getProjectRequest{
		api:       api,
		projectId: projectId,
	}
}

// Responce
type getProjectResponce struct {
	Project project `json:"project"`
}

// Do
func (request *getProjectRequest) Do() (project project, err error) {
	responce := getProjectResponce{}
	err = request.api.makeRequest("Get Project", request.projectId, nil, nil, &responce)
	if err != nil {
		return
	}

	project = responce.Project
	return
}
