package geoedge

import "strings"

// Request
type deleteProjectRequest struct {
	api *geoedgeApi

	ProjectIds []string
}

func (api *geoedgeApi) DeleteProjectRequest(projectIds []string) *deleteProjectRequest {
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
