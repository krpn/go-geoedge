package geoedge

import "strings"

// Request
type editProjectRequest struct {
	api *api

	ProjectId string `json:"-"`

	AutoScan      *int    `json:"auto_scan,omitempty"`
	Name          *string `json:"name,omitempty"`
	Tag           *string `json:"tag,omitempty"`
	ScanType      *int    `json:"scan_type,omitempty"`
	Locations     *string `json:"locations,omitempty"`
	TimesPerDay   *int    `json:"times_per_day,omitempty"`
	Emulators     *string `json:"emulators,omitempty"`
	EmulationSets *string `json:"emulation_sets,omitempty"`
	EmulationType *int    `json:"emulation_type,omitempty"`
	TopUrls       *string `json:"top_urls,omitempty"`
	Useragents    *string `json:"useragents,omitempty"`
}

func (api *api) EditProjectRequest(projectId string, options ...editProjectRequestOption) *editProjectRequest {
	request := editProjectRequest{
		api:       api,
		ProjectId: projectId,
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Do
func (request *editProjectRequest) Do() (err error) {
	reqBody, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	extraHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	return request.api.makeRequest("Edit Project", request.ProjectId, reqBody, extraHeaders, nil)
}

// Options
type editProjectRequestOption func(*editProjectRequest)

func EditProjectRequestAutoScan(autoScan int) editProjectRequestOption {
	return func(request *editProjectRequest) {
		request.AutoScan = &autoScan
	}
}

func EditProjectRequestName(name string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		request.Name = &name
	}
}

func EditProjectRequestTag(tag string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		request.Tag = &tag
	}
}

func EditProjectRequestScanType(scanType int) editProjectRequestOption {
	return func(request *editProjectRequest) {
		request.ScanType = &scanType
	}
}

func EditProjectRequestLocations(locations []string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		locationsStr := strings.Join(locations, ",")
		request.Locations = &locationsStr
	}
}

func EditProjectRequestTimesPerDay(timesPerDay int) editProjectRequestOption {
	return func(request *editProjectRequest) {
		request.TimesPerDay = &timesPerDay
	}
}

func EditProjectRequestEmulators(emulators []string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		emulatorsStr := strings.Join(emulators, ",")
		request.Emulators = &emulatorsStr
	}
}

func EditProjectRequestEmulationSets(emulationSets []string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		emulationSetsStr := strings.Join(emulationSets, ",")
		request.EmulationSets = &emulationSetsStr
	}
}

func EditProjectRequestEmulationType(emulationType int) editProjectRequestOption {
	return func(request *editProjectRequest) {
		request.EmulationType = &emulationType
	}
}

func EditProjectRequestTopUrls(topUrls []string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		topUrlsStr := strings.Join(topUrls, ",")
		request.TopUrls = &topUrlsStr
	}
}

func EditProjectRequestUseragents(useragents []string) editProjectRequestOption {
	return func(request *editProjectRequest) {
		useragentsStr := strings.Join(useragents, ",")
		request.Useragents = &useragentsStr
	}
}
