package geoedge

import "strings"

// Request
type addProjectRequest struct {
	api *api

	Name      string `json:"name"`
	Tag       string `json:"tag"`
	AutoScan  int    `json:"auto_scan"`
	ScanType  int    `json:"scan_type"`
	Locations string `json:"locations"`

	TimesPerDay                         *int    `json:"times_per_day,omitempty"`
	Emulators                           *string `json:"emulators,omitempty"`
	EmulationSets                       *string `json:"emulation_sets,omitempty"`
	EmulationType                       *int    `json:"emulation_type,omitempty"`
	ExtLineitemId                       *string `json:"ext_lineitem_id,omitempty"`
	ExtCreativeId                       *string `json:"ext_creative_id,omitempty"`
	TopUrls                             *string `json:"top_urls,omitempty"`
	TriggerAdMaxFileSize                *int    `json:"trigger_ad_max_file_size,omitempty"`
	TriggerAdPoliteMaxDownloadSize      *int    `json:"trigger_ad_polite_max_download_size,omitempty"`
	TriggerAdMaxLoadTime                *int    `json:"trigger_ad_max_load_time,omitempty"`
	TriggerAdMaxLoadTimeFromStart       *int    `json:"trigger_ad_max_load_time_from_start,omitempty"`
	TriggerAdMaxAllowedDimensions       *string `json:"trigger_ad_max_allowed_dimensions,omitempty"`
	TriggerAdMaxAllowedDimensionsMargin *int    `json:"trigger_ad_max_allowed_dimensions_margin,omitempty"`
	TriggerScanWithHttpOverHttps        *int    `json:"trigger_scan_with_http_over_https,omitempty"`
	NotificationEmail                   *string `json:"notification_email,omitempty"`
	NotificationTime                    *int    `json:"notification_time,omitempty"`
	AlertPixelNotification              *int    `json:"alert_pixel_notification,omitempty"`
	Useragents                          *string `json:"useragents,omitempty"`
}

func (api *api) AddProjectRequestRequest(
	name, tag string,
	autoScan, scanType int,
	locations []string,
	options ...addProjectRequestOption) *addProjectRequest {

	request := addProjectRequest{
		api:       api,
		Name:      name,
		Tag:       tag,
		AutoScan:  autoScan,
		ScanType:  scanType,
		Locations: strings.Join(locations, ","),
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type addProjectResponce struct {
	ProjectId string `json:"project_id"`
}

// Do
func (request *addProjectRequest) Do() (projectId string, err error) {
	reqBody, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	extraHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	responce := addProjectResponce{}
	err = request.api.makeRequest("Add Project", "", reqBody, extraHeaders, &responce)
	if err != nil {
		return
	}

	projectId = responce.ProjectId
	return
}

// Options
type addProjectRequestOption func(*addProjectRequest)

func AddProjectRequestTimesPerDay(timesPerDay int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TimesPerDay = &timesPerDay
	}
}

func AddProjectRequestEmulators(emulators []string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		emulatorsStr := strings.Join(emulators, ",")
		request.Emulators = &emulatorsStr
	}
}

func AddProjectRequestEmulationSets(emulationSets []string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		emulationSetsStr := strings.Join(emulationSets, ",")
		request.EmulationSets = &emulationSetsStr
	}
}
func AddProjectRequestEmulationType(emulationType int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.EmulationType = &emulationType
	}
}

func AddProjectRequestExtLineitemId(extLineitemId string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.ExtLineitemId = &extLineitemId
	}
}

func AddProjectRequestExtCreativeId(extCreativeId string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.ExtCreativeId = &extCreativeId
	}
}

func AddProjectRequestTopUrls(topUrls []string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		topUrlsStr := strings.Join(topUrls, ",")
		request.TopUrls = &topUrlsStr
	}
}

func AddProjectRequestTriggerAdMaxFileSize(triggerAdMaxFileSize int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TriggerAdMaxFileSize = &triggerAdMaxFileSize
	}
}

func AddProjectRequestTriggerAdPoliteMaxDownloadSize(triggerAdPoliteMaxDownloadSize int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TriggerAdPoliteMaxDownloadSize = &triggerAdPoliteMaxDownloadSize
	}
}

func AddProjectRequestTriggerAdMaxLoadTime(triggerAdMaxLoadTime int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TriggerAdMaxLoadTime = &triggerAdMaxLoadTime
	}
}

func AddProjectRequestTriggerAdMaxLoadTimeFromStart(triggerAdMaxLoadTimeFromStart int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TriggerAdMaxLoadTimeFromStart = &triggerAdMaxLoadTimeFromStart
	}
}

func AddProjectRequestTriggerAdMaxAllowedDimensions(triggerAdMaxAllowedDimensions []string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		triggerAdMaxAllowedDimensionsStr := strings.Join(triggerAdMaxAllowedDimensions, ",")
		request.TriggerAdMaxAllowedDimensions = &triggerAdMaxAllowedDimensionsStr
	}
}

func AddProjectRequestTriggerAdMaxAllowedDimensionsMargin(triggerAdMaxAllowedDimensionsMargin int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TriggerAdMaxAllowedDimensionsMargin = &triggerAdMaxAllowedDimensionsMargin
	}
}
func AddProjectRequestTriggerScanWithHttpOverHttps(triggerScanWithHttpOverHttps int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.TriggerScanWithHttpOverHttps = &triggerScanWithHttpOverHttps
	}
}

func AddProjectRequestNotificationEmail(notificationEmail []string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		notificationEmailStr := strings.Join(notificationEmail, ",")
		request.NotificationEmail = &notificationEmailStr
	}
}

func AddProjectRequestNotificationTime(notificationTime int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.NotificationTime = &notificationTime
	}
}

func AddProjectRequestAlertPixelNotification(alertPixelNotification int) addProjectRequestOption {
	return func(request *addProjectRequest) {
		request.AlertPixelNotification = &alertPixelNotification
	}
}

func AddProjectRequestUseragents(useragents []string) addProjectRequestOption {
	return func(request *addProjectRequest) {
		useragentsStr := strings.Join(useragents, ",")
		request.Useragents = &useragentsStr
	}
}
