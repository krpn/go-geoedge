package geoedge

import (
	"encoding/json"
	"strings"
)

// Request
type addMultipleProjectsRequest struct {
	api *api

	AddProjectRequests []*AddProjectRequest
}

func (api *api) AddMultipleProjectsRequest(addProjectRequests ...*AddProjectRequest) *addMultipleProjectsRequest {
	return &addMultipleProjectsRequest{
		api:                api,
		AddProjectRequests: addProjectRequests,
	}
}

// Responce
type addMultipleProjectsResult struct {
	Status    responceStatus `json:"status"`
	ProjectId string         `json:"project_id"`
}

type addMultipleProjectsResponce struct {
	Projects []addMultipleProjectsResult `json:"projects"`
}

func (request *addMultipleProjectsRequest) Do() (projects []addMultipleProjectsResult, err error) {
	reqBody, err := json.Marshal(request.AddProjectRequests)
	if err != nil {
		return
	}

	reqBody = []byte(strings.Replace(string(reqBody), `/`, `\/`, -1))

	extraHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	responce := addMultipleProjectsResponce{}
	err = request.api.makeRequest("Add Multiple Projects", "", reqBody, extraHeaders, &responce)
	if err != nil {
		return
	}

	projects = responce.Projects
	return
}
