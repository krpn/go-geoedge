package geoedge

// Request
type listProjectsRequest struct {
	api *api

	Offset *int `json:"offset,omitempty"`
	Limit  *int `json:"limit,omitempty"`
}

func (api *api) ListProjectsRequest(options ...listProjectsRequestOption) *listProjectsRequest {
	request := listProjectsRequest{api: api}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type listProjectsResponce struct {
	PrevPage      string        `json:"prev_page"`
	NextPage      string        `json:"next_page"`
	ProjectsTotal int           `json:"projects_total"`
	Projects      []projectBase `json:"projects"`
}

// Do
func (request *listProjectsRequest) DoExtra() (projects []projectBase, prevPage, nextPage string, projectsTotal int, err error) {
	urlParams, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	responce := listProjectsResponce{}
	err = request.api.makeRequest("List Projects", string(urlParams), nil, nil, &responce)
	if err != nil {
		return
	}

	projects = responce.Projects
	prevPage, nextPage, projectsTotal = responce.PrevPage, responce.NextPage, responce.ProjectsTotal
	return
}

func (request *listProjectsRequest) Do() (projects []projectBase, err error) {
	projects, _, _, _, err = request.DoExtra()
	return
}

// Options
type listProjectsRequestOption func(*listProjectsRequest)

func ListProjectsRequestOffset(offset int) listProjectsRequestOption {
	return func(request *listProjectsRequest) {
		request.Offset = &offset
	}
}

func ListProjectsRequestLimit(limit int) listProjectsRequestOption {
	return func(request *listProjectsRequest) {
		request.Limit = &limit
	}
}
