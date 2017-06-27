package geoedge

// Request
type searchProjectsRequest struct {
	api *api

	Tag  *string `json:"tag,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (api *api) SearchProjectsRequest(options ...searchProjectsRequestOption) *searchProjectsRequest {
	request := searchProjectsRequest{api: api}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type searchProjectsResponce struct {
	Projects []projectBase `json:"projects"`
}

// Do
func (request *searchProjectsRequest) Do() (projects []projectBase, err error) {
	urlParams, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	responce := listProjectsResponce{}
	err = request.api.makeRequest("Search Projects", string(urlParams), nil, nil, &responce)
	if err != nil {
		return
	}

	projects = responce.Projects
	return
}

// Options
type searchProjectsRequestOption func(*searchProjectsRequest)

func SearchProjectsRequestTag(tag string) searchProjectsRequestOption {
	return func(request *searchProjectsRequest) {
		request.Tag = &tag
	}
}

func SearchProjectsRequestName(name string) searchProjectsRequestOption {
	return func(request *searchProjectsRequest) {
		request.Name = &name
	}
}
