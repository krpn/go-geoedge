package geoedge

import "strings"

// Request
type alertProjectsBindRequest struct {
	api *geoedgeApi

	AlertId    string `json:"-"`
	ProjectIds string `json:"projects"`

	ActionType *int `json:"action_type,omitempty"`
}

func (api *geoedgeApi) AlertProjectsBindRequest(
	alertId string,
	projectIds []string,
	options ...alertProjectsBindRequestOption) *alertProjectsBindRequest {

	request := alertProjectsBindRequest{
		api:        api,
		AlertId:    alertId,
		ProjectIds: strings.Join(projectIds, ","),
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Do
func (request *alertProjectsBindRequest) Do() (err error) {
	reqBody, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	extraHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	return request.api.makeRequest("Alert Projects Bind", request.AlertId, reqBody, extraHeaders, nil)
}

// Options
type alertProjectsBindRequestOption func(*alertProjectsBindRequest)

func AlertProjectsBindRequestActionType(actionType int) alertProjectsBindRequestOption {
	return func(request *alertProjectsBindRequest) {
		request.ActionType = &actionType
	}
}
