package geoedge

import "strings"

// Request
type deleteProjectRequest struct {
	api *api

	ProjectIds []string
}

func (api *api) DeleteProjectRequest(projectIds []string) *deleteProjectRequest {
	return &deleteProjectRequest{
		api:        api,
		ProjectIds: projectIds,
	}
}

// Do
func (request *deleteProjectRequest) Do() (err error) {
	urlParams := strings.Join(request.ProjectIds, ",")
	return request.api.makeRequest("Delete Projects", urlParams, nil, nil, nil)

}
