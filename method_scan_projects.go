package geoedge

import "strings"

// Request
type scanProjectRequest struct {
	api *geoedgeApi

	ProjectIds []string `json:"project_id"`
}

func (api *geoedgeApi) ScanProjectRequest(projectIds []string) *scanProjectRequest {
	return &scanProjectRequest{
		api:        api,
		ProjectIds: projectIds,
	}
}

// Responce
type scanProjectsResponceResult struct {
	ProjectId string
	Scans     []string
}

type scanProjectsResponce struct {
	Projects map[string]struct {
		Scans []string `json:"scans"`
	} `json:"projects"`
}

// Do
func (request *scanProjectRequest) Do() (result []scanProjectsResponceResult, err error) {
	urlParams := strings.Join(request.ProjectIds, ",")

	responce := scanProjectsResponce{}
	err = request.api.makeRequest("Scan Projects", urlParams, nil, nil, &responce)
	if err != nil {
		return
	}

	for key, val := range responce.Projects {
		result = append(result,
			scanProjectsResponceResult{
				ProjectId: key,
				Scans:     val.Scans},
		)
	}

	return
}
