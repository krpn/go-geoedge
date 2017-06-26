package geoedge

import "strings"

// Request
type addAlertRequest struct {
	api *geoedgeApi

	Name string `json:"name"`

	Locations                                   *string `json:"locations,omitempty"`
	Platforms                                   *string `json:"platforms,omitempty"`
	NotificationEmails                          *string `json:"notification_emails,omitempty"`
	NotificationTime                            *int    `json:"notification_time,omitempty"`
	AlertPixelNotification                      *int    `json:"alert_pixel_notification,omitempty"`
	TriggerAdMalware                            *int    `json:"trigger_ad_malware,omitempty"`
	TriggerLpMalware                            *int    `json:"trigger_lp_malware,omitempty"`
	TriggerScanMalware                          *int    `json:"trigger_scan_malware,omitempty"`
	TriggerAdLpPhishing                         *int    `json:"trigger_ad_lp_phishing,omitempty"`
	TriggerLpFileDownload                       *int    `json:"trigger_lp_file_download,omitempty"`
	TriggerAdFileDownload                       *int    `json:"trigger_ad_file_download,omitempty"`
	TriggerScanMalicious                        *int    `json:"trigger_scan_malicious,omitempty"`
	TriggerAdMalicious                          *int    `json:"trigger_ad_malicious,omitempty"`
	TriggerLpMalicious                          *int    `json:"trigger_lp_malicious,omitempty"`
	TriggerAdAutomaticRedirectDomains           *string `json:"trigger_ad_automatic_redirect_domains,omitempty"`
	TriggerAdLpOffensive                        *int    `json:"trigger_ad_lp_offensive,omitempty"`
	TriggerAdAutoPlaySound                      *int    `json:"trigger_ad_auto_play_sound,omitempty"`
	TriggerAdDownloadButton                     *int    `json:"trigger_ad_download_button,omitempty"`
	TriggerAdDeceptive                          *int    `json:"trigger_ad_deceptive,omitempty"`
	TriggerLpCategories                         *string `json:"trigger_lp_categories,omitempty"`
	TriggerLpAdvertiserByDomains                *string `json:"trigger_lp_advertiser_by_domains,omitempty"`
	TriggerLpAdvertiserByNames                  *string `json:"trigger_lp_advertiser_by_names,omitempty"`
	TriggerLpAdvertiserByKeywords               *string `json:"trigger_lp_advertiser_by_keywords,omitempty"`
	TriggerAdChanged                            *int    `json:"trigger_ad_changed,omitempty"`
	TriggerLpChanged                            *int    `json:"trigger_lp_changed,omitempty"`
	TriggerLpMaxSize                            *int    `json:"trigger_lp_max_size,omitempty"`
	TriggerAdPopup                              *int    `json:"trigger_ad_popup,omitempty"`
	TriggerAdPopunder                           *int    `json:"trigger_ad_popunder,omitempty"`
	TriggerAdPopupWithJsAlert                   *int    `json:"trigger_ad_popup_with_js_alert,omitempty"`
	TriggerAdPopupWithJsAlertOnExit             *int    `json:"trigger_ad_popup_with_js_alert_on_exit,omitempty"`
	TriggerAdBrowserLocking                     *int    `json:"trigger_ad_browser_locking,omitempty"`
	TriggerScanIsMultiPop                       *int    `json:"trigger_scan_is_multi_pop,omitempty"`
	TriggerLpAutoPlaySound                      *int    `json:"trigger_lp_auto_play_sound,omitempty"`
	TriggerAdSpecificBanner                     *string `json:"trigger_ad_specific_banner,omitempty"`
	TriggerLpVibrate                            *int    `json:"trigger_lp_vibrate,omitempty"`
	TriggerAdMaxFileSize                        *int    `json:"trigger_ad_max_file_size,omitempty"`
	TriggerAdMaxExpandedAllowedDimensions       *string `json:"trigger_ad_max_expanded_allowed_dimensions,omitempty"`
	TriggerAdMaxExpandedAllowedDimensionsMargin *int    `json:"trigger_ad_max_expanded_allowed_dimensions_margin,omitempty"`
	TriggerAdMaxExpandedAllowedDimensionsType   *int    `json:"trigger_ad_max_expanded_allowed_dimensions_type,omitempty"`
	TriggerAdHtml5ImageMaxFileSize              *int    `json:"trigger_ad_html5_image_max_file_size,omitempty"`
	TriggerAdHtml5JsMaxFileSize                 *int    `json:"trigger_ad_html5_js_max_file_size,omitempty"`
	TriggerAdHtml5InitMaxFileSize               *int    `json:"trigger_ad_html5_init_max_file_size,omitempty"`
	TriggerAdHtml5SubMaxFileSize                *int    `json:"trigger_ad_html5_sub_max_file_size,omitempty"`
	TriggerAdHtml5Zindex                        *string `json:"trigger_ad_html5_zindex,omitempty"`
	TriggerAdHtml5MaxRequests                   *int    `json:"trigger_ad_html5_max_requests,omitempty"`
	TriggerAdHtml5InitRequests                  *int    `json:"trigger_ad_html5_init_requests,omitempty"`
	TriggerAdHtml5SubRequests                   *int    `json:"trigger_ad_html5_sub_requests,omitempty"`
	TriggerAdMaxRequests                        *int    `json:"trigger_ad_max_requests,omitempty"`
	TriggerAdPoliteMaxDownloadSize              *int    `json:"trigger_ad_polite_max_download_size,omitempty"`
	TriggerAdJsPoliteBeforeOnload               *int    `json:"trigger_ad_js_polite_before_onload,omitempty"`
	TriggerAdMaxLoadTime                        *int    `json:"trigger_ad_max_load_time,omitempty"`
	TriggerAdMaxLoadTimeFromStart               *int    `json:"trigger_ad_max_load_time_from_start,omitempty"`
	TriggerAdMaxAllowedDimensions               *string `json:"trigger_ad_max_allowed_dimensions,omitempty"`
	TriggerAdMaxAllowedDimensionsMargin         *int    `json:"trigger_ad_max_allowed_dimensions_margin,omitempty"`
	TriggerAdAllowedRatio                       *string `json:"trigger_ad_allowed_ratio,omitempty"`
	TriggerAd1X1                                *int    `json:"trigger_ad_1x1,omitempty"`
	TriggerLpError                              *int    `json:"trigger_lp_error,omitempty"`
	TriggerAdFlashHighCpu                       *int    `json:"trigger_ad_flash_high_cpu,omitempty"`
	TriggerScanWithHttpOverHttps                *int    `json:"trigger_scan_with_http_over_https,omitempty"`
	TriggerLpCascadingNetworks                  *int    `json:"trigger_lp_cascading_networks,omitempty"`
	TriggerScanRequestsErrors                   *int    `json:"trigger_scan_requests_errors,omitempty"`
	TriggerScanTotalTime                        *int    `json:"trigger_scan_total_time,omitempty"`
	TriggerScanRequestsSize                     *int    `json:"trigger_scan_requests_size,omitempty"`
	TriggerAdAnimationMaxLength                 *int    `json:"trigger_ad_animation_max_length,omitempty"`
	TriggerSslCertViolation                     *int    `json:"trigger_ssl_cert_violation,omitempty"`
	TriggerAdCookiesNetworks                    *string `json:"trigger_ad_cookies_networks,omitempty"`
	TriggerAdCookiesDomains                     *string `json:"trigger_ad_cookies_domains,omitempty"`
	TriggerAdNetworks                           *string `json:"trigger_ad_networks,omitempty"`
	TriggerAdDomains                            *string `json:"trigger_ad_domains,omitempty"`
	TriggerLpDomains                            *string `json:"trigger_lp_domains,omitempty"`
	TriggerAdVideoInDisplay                     *int    `json:"trigger_ad_video_in_display,omitempty"`
	TriggerAdVideoMaxLength                     *int    `json:"trigger_ad_video_max_length,omitempty"`
	TriggerAdVideoSupportedTypes                *string `json:"trigger_ad_video_supported_types,omitempty"`
	TriggerAdVideoVastFromStart                 *int    `json:"trigger_ad_video_vast_from_start,omitempty"`
	TriggerAdVideoAutoSound                     *int    `json:"trigger_ad_video_auto_sound,omitempty"`
	TriggerAdVideoVastMissingMediaFile          *int    `json:"trigger_ad_video_vast_missing_media_file,omitempty"`
	TriggerAdVideoAutoScroll                    *int    `json:"trigger_ad_video_auto_scroll,omitempty"`
	TriggerAdVideoFrameRate                     *int    `json:"trigger_ad_video_frame_rate,omitempty"`
	TriggerAdVideoBitRate                       *int    `json:"trigger_ad_video_bit_rate,omitempty"`
	TriggerAdVideoCompanionAds                  *int    `json:"trigger_ad_video_companion_ads,omitempty"`
	TriggerAdVideoFileSize                      *int    `json:"trigger_ad_video_file_size,omitempty"`
	TriggerVpaidFormat                          *int    `json:"trigger_vpaid_format,omitempty"`
}

func (api *geoedgeApi) AddAlertRequestRequest(name string, options ...addAlertRequestOption) *addAlertRequest {
	request := addAlertRequest{
		api:  api,
		Name: name,
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Responce
type addAlertResponce struct {
	AlertId string `json:"alert_id"`
}

// Do
func (request *addAlertRequest) Do() (alertId string, err error) {
	reqBody, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	extraHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	responce := addAlertResponce{}
	err = request.api.makeRequest("Add Alert", "", reqBody, extraHeaders, &responce)
	if err != nil {
		return
	}

	alertId = responce.AlertId
	return
}

// Options
type addAlertRequestOption func(*addAlertRequest)

func AddAlertRequestLocations(locations []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		locationsStr := strings.Join(locations, ",")
		request.Locations = &locationsStr
	}
}

func AddAlertRequestPlatforms(platforms []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		platformsStr := strings.Join(platforms, ",")
		request.Platforms = &platformsStr
	}
}

func AddAlertRequestNotificationEmails(notificationEmails []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		notificationEmailsStr := strings.Join(notificationEmails, ",")
		request.NotificationEmails = &notificationEmailsStr
	}
}

func AddAlertRequestNotificationTime(notificationTime int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.NotificationTime = &notificationTime
	}
}

func AddAlertRequestAlertPixelNotification(alertPixelNotification int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.AlertPixelNotification = &alertPixelNotification
	}
}

func AddAlertRequestTriggerAdMalware(triggerAdMalware int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMalware = &triggerAdMalware
	}
}

func AddAlertRequestTriggerLpMalware(triggerLpMalware int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpMalware = &triggerLpMalware
	}
}

func AddAlertRequestTriggerScanMalware(triggerScanMalware int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanMalware = &triggerScanMalware
	}
}

func AddAlertRequestTriggerAdLpPhishing(triggerAdLpPhishing int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdLpPhishing = &triggerAdLpPhishing
	}
}

func AddAlertRequestTriggerLpFileDownload(triggerLpFileDownload int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpFileDownload = &triggerLpFileDownload
	}
}

func AddAlertRequestTriggerAdFileDownload(triggerAdFileDownload int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdFileDownload = &triggerAdFileDownload
	}
}

func AddAlertRequestTriggerScanMalicious(triggerScanMalicious int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanMalicious = &triggerScanMalicious
	}
}

func AddAlertRequestTriggerAdMalicious(triggerAdMalicious int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMalicious = &triggerAdMalicious
	}
}

func AddAlertRequestTriggerLpMalicious(triggerLpMalicious int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpMalicious = &triggerLpMalicious
	}
}

func AddAlertRequestTriggerAdAutomaticRedirectDomains(triggerAdAutomaticRedirectDomains []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdAutomaticRedirectDomainsStr := strings.Join(triggerAdAutomaticRedirectDomains, ",")
		request.TriggerAdAutomaticRedirectDomains = &triggerAdAutomaticRedirectDomainsStr
	}
}

func AddAlertRequestTriggerAdLpOffensive(triggerAdLpOffensive int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdLpOffensive = &triggerAdLpOffensive
	}
}

func AddAlertRequestTriggerAdAutoPlaySound(triggerAdAutoPlaySound int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdAutoPlaySound = &triggerAdAutoPlaySound
	}
}

func AddAlertRequestTriggerAdDownloadButton(triggerAdDownloadButton int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdDownloadButton = &triggerAdDownloadButton
	}
}

func AddAlertRequestTriggerAdDeceptive(triggerAdDeceptive int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdDeceptive = &triggerAdDeceptive
	}
}

func AddAlertRequestTriggerLpCategories(triggerLpCategories []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerLpCategoriesStr := strings.Join(triggerLpCategories, ",")
		request.TriggerLpCategories = &triggerLpCategoriesStr
	}
}

func AddAlertRequestTriggerLpAdvertiserByDomains(triggerLpAdvertiserByDomains []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerLpAdvertiserByDomainsStr := strings.Join(triggerLpAdvertiserByDomains, ",")
		request.TriggerLpAdvertiserByDomains = &triggerLpAdvertiserByDomainsStr
	}
}

func AddAlertRequestTriggerLpAdvertiserByNames(triggerLpAdvertiserByNames []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerLpAdvertiserByNamesStr := strings.Join(triggerLpAdvertiserByNames, ",")
		request.TriggerLpAdvertiserByNames = &triggerLpAdvertiserByNamesStr
	}
}

func AddAlertRequestTriggerLpAdvertiserByKeywords(triggerLpAdvertiserByKeywords []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerLpAdvertiserByKeywordsStr := strings.Join(triggerLpAdvertiserByKeywords, ",")
		request.TriggerLpAdvertiserByKeywords = &triggerLpAdvertiserByKeywordsStr
	}
}

func AddAlertRequestTriggerAdChanged(triggerAdChanged int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdChanged = &triggerAdChanged
	}
}

func AddAlertRequestTriggerLpChanged(triggerLpChanged int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpChanged = &triggerLpChanged
	}
}

func AddAlertRequestTriggerLpMaxSize(triggerLpMaxSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpMaxSize = &triggerLpMaxSize
	}
}

func AddAlertRequestTriggerAdPopup(triggerAdPopup int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdPopup = &triggerAdPopup
	}
}

func AddAlertRequestTriggerAdPopunder(triggerAdPopunder int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdPopunder = &triggerAdPopunder
	}
}

func AddAlertRequestTriggerAdPopupWithJsAlert(triggerAdPopupWithJsAlert int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdPopupWithJsAlert = &triggerAdPopupWithJsAlert
	}
}

func AddAlertRequestTriggerAdPopupWithJsAlertOnExit(triggerAdPopupWithJsAlertOnExit int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdPopupWithJsAlertOnExit = &triggerAdPopupWithJsAlertOnExit
	}
}

func AddAlertRequestTriggerAdBrowserLocking(triggerAdBrowserLocking int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdBrowserLocking = &triggerAdBrowserLocking
	}
}

func AddAlertRequestTriggerScanIsMultiPop(triggerScanIsMultiPop int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanIsMultiPop = &triggerScanIsMultiPop
	}
}

func AddAlertRequestTriggerLpAutoPlaySound(triggerLpAutoPlaySound int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpAutoPlaySound = &triggerLpAutoPlaySound
	}
}

func AddAlertRequestTriggerAdSpecificBanner(triggerAdSpecificBanner []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdSpecificBannerStr := strings.Join(triggerAdSpecificBanner, ",")
		request.TriggerAdSpecificBanner = &triggerAdSpecificBannerStr
	}
}

func AddAlertRequestTriggerLpVibrate(triggerLpVibrate int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpVibrate = &triggerLpVibrate
	}
}

func AddAlertRequestTriggerAdMaxFileSize(triggerAdMaxFileSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxFileSize = &triggerAdMaxFileSize
	}
}

func AddAlertRequestTriggerAdMaxExpandedAllowedDimensions(triggerAdMaxExpandedAllowedDimensions []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdMaxExpandedAllowedDimensionsStr := strings.Join(triggerAdMaxExpandedAllowedDimensions, ",")
		request.TriggerAdMaxExpandedAllowedDimensions = &triggerAdMaxExpandedAllowedDimensionsStr
	}
}

func AddAlertRequestTriggerAdMaxExpandedAllowedDimensionsMargin(triggerAdMaxExpandedAllowedDimensionsMargin int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxExpandedAllowedDimensionsMargin = &triggerAdMaxExpandedAllowedDimensionsMargin
	}
}

func AddAlertRequestTriggerAdMaxExpandedAllowedDimensionsType(triggerAdMaxExpandedAllowedDimensionsType int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxExpandedAllowedDimensionsType = &triggerAdMaxExpandedAllowedDimensionsType
	}
}

func AddAlertRequestTriggerAdHtml5ImageMaxFileSize(triggerAdHtml5ImageMaxFileSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5ImageMaxFileSize = &triggerAdHtml5ImageMaxFileSize
	}
}

func AddAlertRequestTriggerAdHtml5JsMaxFileSize(triggerAdHtml5JsMaxFileSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5JsMaxFileSize = &triggerAdHtml5JsMaxFileSize
	}
}

func AddAlertRequestTriggerAdHtml5InitMaxFileSize(triggerAdHtml5InitMaxFileSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5InitMaxFileSize = &triggerAdHtml5InitMaxFileSize
	}
}

func AddAlertRequestTriggerAdHtml5SubMaxFileSize(triggerAdHtml5SubMaxFileSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5SubMaxFileSize = &triggerAdHtml5SubMaxFileSize
	}
}

func AddAlertRequestTriggerAdHtml5Zindex(triggerAdHtml5Zindex string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5Zindex = &triggerAdHtml5Zindex
	}
}

func AddAlertRequestTriggerAdHtml5MaxRequests(triggerAdHtml5MaxRequests int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5MaxRequests = &triggerAdHtml5MaxRequests
	}
}

func AddAlertRequestTriggerAdHtml5InitRequests(triggerAdHtml5InitRequests int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5InitRequests = &triggerAdHtml5InitRequests
	}
}

func AddAlertRequestTriggerAdHtml5SubRequests(triggerAdHtml5SubRequests int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdHtml5SubRequests = &triggerAdHtml5SubRequests
	}
}

func AddAlertRequestTriggerAdMaxRequests(triggerAdMaxRequests int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxRequests = &triggerAdMaxRequests
	}
}

func AddAlertRequestTriggerAdPoliteMaxDownloadSize(triggerAdPoliteMaxDownloadSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdPoliteMaxDownloadSize = &triggerAdPoliteMaxDownloadSize
	}
}

func AddAlertRequestTriggerAdJsPoliteBeforeOnload(triggerAdJsPoliteBeforeOnload int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdJsPoliteBeforeOnload = &triggerAdJsPoliteBeforeOnload
	}
}

func AddAlertRequestTriggerAdMaxLoadTime(triggerAdMaxLoadTime int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxLoadTime = &triggerAdMaxLoadTime
	}
}

func AddAlertRequestTriggerAdMaxLoadTimeFromStart(triggerAdMaxLoadTimeFromStart int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxLoadTimeFromStart = &triggerAdMaxLoadTimeFromStart
	}
}

func AddAlertRequestTriggerAdMaxAllowedDimensions(triggerAdMaxAllowedDimensions []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdMaxAllowedDimensionsStr := strings.Join(triggerAdMaxAllowedDimensions, ",")
		request.TriggerAdMaxAllowedDimensions = &triggerAdMaxAllowedDimensionsStr
	}
}

func AddAlertRequestTriggerAdMaxAllowedDimensionsMargin(triggerAdMaxAllowedDimensionsMargin int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdMaxAllowedDimensionsMargin = &triggerAdMaxAllowedDimensionsMargin
	}
}

func AddAlertRequestTriggerAdAllowedRatio(triggerAdAllowedRatio []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdAllowedRatioStr := strings.Join(triggerAdAllowedRatio, ",")
		request.TriggerAdAllowedRatio = &triggerAdAllowedRatioStr
	}
}

func AddAlertRequestTriggerAd1X1(triggerAd1X1 int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAd1X1 = &triggerAd1X1
	}
}

func AddAlertRequestTriggerLpError(triggerLpError int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpError = &triggerLpError
	}
}

func AddAlertRequestTriggerAdFlashHighCpu(triggerAdFlashHighCpu int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdFlashHighCpu = &triggerAdFlashHighCpu
	}
}

func AddAlertRequestTriggerScanWithHttpOverHttps(triggerScanWithHttpOverHttps int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanWithHttpOverHttps = &triggerScanWithHttpOverHttps
	}
}

func AddAlertRequestTriggerLpCascadingNetworks(triggerLpCascadingNetworks int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerLpCascadingNetworks = &triggerLpCascadingNetworks
	}
}

func AddAlertRequestTriggerScanRequestsErrors(triggerScanRequestsErrors int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanRequestsErrors = &triggerScanRequestsErrors
	}
}

func AddAlertRequestTriggerScanTotalTime(triggerScanTotalTime int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanTotalTime = &triggerScanTotalTime
	}
}

func AddAlertRequestTriggerScanRequestsSize(triggerScanRequestsSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerScanRequestsSize = &triggerScanRequestsSize
	}
}

func AddAlertRequestTriggerAdAnimationMaxLength(triggerAdAnimationMaxLength int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdAnimationMaxLength = &triggerAdAnimationMaxLength
	}
}

func AddAlertRequestTriggerSslCertViolation(triggerSslCertViolation int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerSslCertViolation = &triggerSslCertViolation
	}
}

func AddAlertRequestTriggerAdCookiesNetworks(triggerAdCookiesNetworks []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdCookiesNetworksStr := strings.Join(triggerAdCookiesNetworks, ",")
		request.TriggerAdCookiesNetworks = &triggerAdCookiesNetworksStr
	}
}

func AddAlertRequestTriggerAdCookiesDomains(triggerAdCookiesDomains []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdCookiesDomainsStr := strings.Join(triggerAdCookiesDomains, ",")
		request.TriggerAdCookiesDomains = &triggerAdCookiesDomainsStr
	}
}

func AddAlertRequestTriggerAdNetworks(triggerAdNetworks []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdNetworksStr := strings.Join(triggerAdNetworks, ",")
		request.TriggerAdNetworks = &triggerAdNetworksStr
	}
}

func AddAlertRequestTriggerAdDomains(triggerAdDomains []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerAdDomainsStr := strings.Join(triggerAdDomains, ",")
		request.TriggerAdDomains = &triggerAdDomainsStr
	}
}

func AddAlertRequestTriggerLpDomains(triggerLpDomains []string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		triggerLpDomainsStr := strings.Join(triggerLpDomains, ",")
		request.TriggerLpDomains = &triggerLpDomainsStr
	}
}

func AddAlertRequestTriggerAdVideoInDisplay(triggerAdVideoInDisplay int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoInDisplay = &triggerAdVideoInDisplay
	}
}

func AddAlertRequestTriggerAdVideoMaxLength(triggerAdVideoMaxLength int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoMaxLength = &triggerAdVideoMaxLength
	}
}

func AddAlertRequestTriggerAdVideoSupportedTypes(triggerAdVideoSupportedTypes string) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoSupportedTypes = &triggerAdVideoSupportedTypes
	}
}

func AddAlertRequestTriggerAdVideoVastFromStart(triggerAdVideoVastFromStart int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoVastFromStart = &triggerAdVideoVastFromStart
	}
}

func AddAlertRequestTriggerAdVideoAutoSound(triggerAdVideoAutoSound int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoAutoSound = &triggerAdVideoAutoSound
	}
}

func AddAlertRequestTriggerAdVideoVastMissingMediaFile(triggerAdVideoVastMissingMediaFile int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoVastMissingMediaFile = &triggerAdVideoVastMissingMediaFile
	}
}

func AddAlertRequestTriggerAdVideoAutoScroll(triggerAdVideoAutoScroll int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoAutoScroll = &triggerAdVideoAutoScroll
	}
}

func AddAlertRequestTriggerAdVideoFrameRate(triggerAdVideoFrameRate int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoFrameRate = &triggerAdVideoFrameRate
	}
}

func AddAlertRequestTriggerAdVideoBitRate(triggerAdVideoBitRate int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoBitRate = &triggerAdVideoBitRate
	}
}

func AddAlertRequestTriggerAdVideoCompanionAds(triggerAdVideoCompanionAds int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoCompanionAds = &triggerAdVideoCompanionAds
	}
}

func AddAlertRequestTriggerAdVideoFileSize(triggerAdVideoFileSize int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerAdVideoFileSize = &triggerAdVideoFileSize
	}
}

func AddAlertRequestTriggerVpaidFormat(triggerVpaidFormat int) addAlertRequestOption {
	return func(request *addAlertRequest) {
		request.TriggerVpaidFormat = &triggerVpaidFormat
	}
}
