package geoedge

import "strings"

// Request
type editAlertRequest struct {
	api *api

	AlertId string `json:"-"`

	Name                                        *string `json:"name,omitempty"`
	NotificationEmails                          *string `json:"notification_emails,omitempty"`
	NotificationTime                            *int    `json:"notification_time,omitempty"`
	Locations                                   *string `json:"locations,omitempty"`
	Platforms                                   *string `json:"platforms,omitempty"`
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

func (api *api) EditAlertRequestRequest(alertId string, options ...editAlertRequestOption) *editAlertRequest {
	request := editAlertRequest{
		api:     api,
		AlertId: alertId,
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

// Do
func (request *editAlertRequest) Do() (err error) {
	reqBody, err := structToUrlencodedFormat(request)
	if err != nil {
		return
	}

	extraHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	return request.api.makeRequest("Edit Alert", request.AlertId, reqBody, extraHeaders, nil)
}

// Options
type editAlertRequestOption func(*editAlertRequest)

func EditAlertRequestName(name string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.Name = &name
	}
}

func EditAlertRequestNotificationEmails(notificationEmails []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		notificationEmailsStr := strings.Join(notificationEmails, ",")
		request.NotificationEmails = &notificationEmailsStr
	}
}

func EditAlertRequestNotificationTime(notificationTime int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.NotificationTime = &notificationTime
	}
}

func EditAlertRequestLocations(locations []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		locationsStr := strings.Join(locations, ",")
		request.Locations = &locationsStr
	}
}

func EditAlertRequestPlatforms(platforms []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		platformsStr := strings.Join(platforms, ",")
		request.Platforms = &platformsStr
	}
}

func EditAlertRequestAlertPixelNotification(alertPixelNotification int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.AlertPixelNotification = &alertPixelNotification
	}
}

func EditAlertRequestTriggerAdMalware(triggerAdMalware int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMalware = &triggerAdMalware
	}
}

func EditAlertRequestTriggerLpMalware(triggerLpMalware int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpMalware = &triggerLpMalware
	}
}

func EditAlertRequestTriggerScanMalware(triggerScanMalware int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanMalware = &triggerScanMalware
	}
}

func EditAlertRequestTriggerAdLpPhishing(triggerAdLpPhishing int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdLpPhishing = &triggerAdLpPhishing
	}
}

func EditAlertRequestTriggerLpFileDownload(triggerLpFileDownload int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpFileDownload = &triggerLpFileDownload
	}
}

func EditAlertRequestTriggerAdFileDownload(triggerAdFileDownload int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdFileDownload = &triggerAdFileDownload
	}
}

func EditAlertRequestTriggerScanMalicious(triggerScanMalicious int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanMalicious = &triggerScanMalicious
	}
}

func EditAlertRequestTriggerAdMalicious(triggerAdMalicious int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMalicious = &triggerAdMalicious
	}
}

func EditAlertRequestTriggerLpMalicious(triggerLpMalicious int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpMalicious = &triggerLpMalicious
	}
}

func EditAlertRequestTriggerAdAutomaticRedirectDomains(triggerAdAutomaticRedirectDomains []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdAutomaticRedirectDomainsStr := strings.Join(triggerAdAutomaticRedirectDomains, ",")
		request.TriggerAdAutomaticRedirectDomains = &triggerAdAutomaticRedirectDomainsStr
	}
}

func EditAlertRequestTriggerAdLpOffensive(triggerAdLpOffensive int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdLpOffensive = &triggerAdLpOffensive
	}
}

func EditAlertRequestTriggerAdAutoPlaySound(triggerAdAutoPlaySound int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdAutoPlaySound = &triggerAdAutoPlaySound
	}
}

func EditAlertRequestTriggerAdDownloadButton(triggerAdDownloadButton int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdDownloadButton = &triggerAdDownloadButton
	}
}

func EditAlertRequestTriggerAdDeceptive(triggerAdDeceptive int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdDeceptive = &triggerAdDeceptive
	}
}

func EditAlertRequestTriggerLpCategories(triggerLpCategories []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerLpCategoriesStr := strings.Join(triggerLpCategories, ",")
		request.TriggerLpCategories = &triggerLpCategoriesStr
	}
}

func EditAlertRequestTriggerLpAdvertiserByDomains(triggerLpAdvertiserByDomains []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerLpAdvertiserByDomainsStr := strings.Join(triggerLpAdvertiserByDomains, ",")
		request.TriggerLpAdvertiserByDomains = &triggerLpAdvertiserByDomainsStr
	}
}

func EditAlertRequestTriggerLpAdvertiserByNames(triggerLpAdvertiserByNames []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerLpAdvertiserByNamesStr := strings.Join(triggerLpAdvertiserByNames, ",")
		request.TriggerLpAdvertiserByNames = &triggerLpAdvertiserByNamesStr
	}
}

func EditAlertRequestTriggerLpAdvertiserByKeywords(triggerLpAdvertiserByKeywords []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerLpAdvertiserByKeywordsStr := strings.Join(triggerLpAdvertiserByKeywords, ",")
		request.TriggerLpAdvertiserByKeywords = &triggerLpAdvertiserByKeywordsStr
	}
}

func EditAlertRequestTriggerAdChanged(triggerAdChanged int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdChanged = &triggerAdChanged
	}
}

func EditAlertRequestTriggerLpChanged(triggerLpChanged int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpChanged = &triggerLpChanged
	}
}

func EditAlertRequestTriggerLpMaxSize(triggerLpMaxSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpMaxSize = &triggerLpMaxSize
	}
}

func EditAlertRequestTriggerAdPopup(triggerAdPopup int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdPopup = &triggerAdPopup
	}
}

func EditAlertRequestTriggerAdPopunder(triggerAdPopunder int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdPopunder = &triggerAdPopunder
	}
}

func EditAlertRequestTriggerAdPopupWithJsAlert(triggerAdPopupWithJsAlert int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdPopupWithJsAlert = &triggerAdPopupWithJsAlert
	}
}

func EditAlertRequestTriggerAdPopupWithJsAlertOnExit(triggerAdPopupWithJsAlertOnExit int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdPopupWithJsAlertOnExit = &triggerAdPopupWithJsAlertOnExit
	}
}

func EditAlertRequestTriggerAdBrowserLocking(triggerAdBrowserLocking int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdBrowserLocking = &triggerAdBrowserLocking
	}
}

func EditAlertRequestTriggerScanIsMultiPop(triggerScanIsMultiPop int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanIsMultiPop = &triggerScanIsMultiPop
	}
}

func EditAlertRequestTriggerLpAutoPlaySound(triggerLpAutoPlaySound int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpAutoPlaySound = &triggerLpAutoPlaySound
	}
}

func EditAlertRequestTriggerAdSpecificBanner(triggerAdSpecificBanner []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdSpecificBannerStr := strings.Join(triggerAdSpecificBanner, ",")
		request.TriggerAdSpecificBanner = &triggerAdSpecificBannerStr
	}
}

func EditAlertRequestTriggerLpVibrate(triggerLpVibrate int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpVibrate = &triggerLpVibrate
	}
}

func EditAlertRequestTriggerAdMaxFileSize(triggerAdMaxFileSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxFileSize = &triggerAdMaxFileSize
	}
}

func EditAlertRequestTriggerAdMaxExpandedAllowedDimensions(triggerAdMaxExpandedAllowedDimensions []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdMaxExpandedAllowedDimensionsStr := strings.Join(triggerAdMaxExpandedAllowedDimensions, ",")
		request.TriggerAdMaxExpandedAllowedDimensions = &triggerAdMaxExpandedAllowedDimensionsStr
	}
}

func EditAlertRequestTriggerAdMaxExpandedAllowedDimensionsMargin(triggerAdMaxExpandedAllowedDimensionsMargin int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxExpandedAllowedDimensionsMargin = &triggerAdMaxExpandedAllowedDimensionsMargin
	}
}

func EditAlertRequestTriggerAdMaxExpandedAllowedDimensionsType(triggerAdMaxExpandedAllowedDimensionsType int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxExpandedAllowedDimensionsType = &triggerAdMaxExpandedAllowedDimensionsType
	}
}

func EditAlertRequestTriggerAdHtml5ImageMaxFileSize(triggerAdHtml5ImageMaxFileSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5ImageMaxFileSize = &triggerAdHtml5ImageMaxFileSize
	}
}

func EditAlertRequestTriggerAdHtml5JsMaxFileSize(triggerAdHtml5JsMaxFileSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5JsMaxFileSize = &triggerAdHtml5JsMaxFileSize
	}
}

func EditAlertRequestTriggerAdHtml5InitMaxFileSize(triggerAdHtml5InitMaxFileSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5InitMaxFileSize = &triggerAdHtml5InitMaxFileSize
	}
}

func EditAlertRequestTriggerAdHtml5SubMaxFileSize(triggerAdHtml5SubMaxFileSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5SubMaxFileSize = &triggerAdHtml5SubMaxFileSize
	}
}

func EditAlertRequestTriggerAdHtml5Zindex(triggerAdHtml5Zindex string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5Zindex = &triggerAdHtml5Zindex
	}
}

func EditAlertRequestTriggerAdHtml5MaxRequests(triggerAdHtml5MaxRequests int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5MaxRequests = &triggerAdHtml5MaxRequests
	}
}

func EditAlertRequestTriggerAdHtml5InitRequests(triggerAdHtml5InitRequests int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5InitRequests = &triggerAdHtml5InitRequests
	}
}

func EditAlertRequestTriggerAdHtml5SubRequests(triggerAdHtml5SubRequests int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdHtml5SubRequests = &triggerAdHtml5SubRequests
	}
}

func EditAlertRequestTriggerAdMaxRequests(triggerAdMaxRequests int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxRequests = &triggerAdMaxRequests
	}
}

func EditAlertRequestTriggerAdPoliteMaxDownloadSize(triggerAdPoliteMaxDownloadSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdPoliteMaxDownloadSize = &triggerAdPoliteMaxDownloadSize
	}
}

func EditAlertRequestTriggerAdJsPoliteBeforeOnload(triggerAdJsPoliteBeforeOnload int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdJsPoliteBeforeOnload = &triggerAdJsPoliteBeforeOnload
	}
}

func EditAlertRequestTriggerAdMaxLoadTime(triggerAdMaxLoadTime int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxLoadTime = &triggerAdMaxLoadTime
	}
}

func EditAlertRequestTriggerAdMaxLoadTimeFromStart(triggerAdMaxLoadTimeFromStart int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxLoadTimeFromStart = &triggerAdMaxLoadTimeFromStart
	}
}

func EditAlertRequestTriggerAdMaxAllowedDimensions(triggerAdMaxAllowedDimensions []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdMaxAllowedDimensionsStr := strings.Join(triggerAdMaxAllowedDimensions, ",")
		request.TriggerAdMaxAllowedDimensions = &triggerAdMaxAllowedDimensionsStr
	}
}

func EditAlertRequestTriggerAdMaxAllowedDimensionsMargin(triggerAdMaxAllowedDimensionsMargin int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdMaxAllowedDimensionsMargin = &triggerAdMaxAllowedDimensionsMargin
	}
}

func EditAlertRequestTriggerAdAllowedRatio(triggerAdAllowedRatio []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdAllowedRatioStr := strings.Join(triggerAdAllowedRatio, ",")
		request.TriggerAdAllowedRatio = &triggerAdAllowedRatioStr
	}
}

func EditAlertRequestTriggerAd1X1(triggerAd1X1 int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAd1X1 = &triggerAd1X1
	}
}

func EditAlertRequestTriggerLpError(triggerLpError int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpError = &triggerLpError
	}
}

func EditAlertRequestTriggerAdFlashHighCpu(triggerAdFlashHighCpu int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdFlashHighCpu = &triggerAdFlashHighCpu
	}
}

func EditAlertRequestTriggerScanWithHttpOverHttps(triggerScanWithHttpOverHttps int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanWithHttpOverHttps = &triggerScanWithHttpOverHttps
	}
}

func EditAlertRequestTriggerLpCascadingNetworks(triggerLpCascadingNetworks int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerLpCascadingNetworks = &triggerLpCascadingNetworks
	}
}

func EditAlertRequestTriggerScanRequestsErrors(triggerScanRequestsErrors int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanRequestsErrors = &triggerScanRequestsErrors
	}
}

func EditAlertRequestTriggerScanTotalTime(triggerScanTotalTime int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanTotalTime = &triggerScanTotalTime
	}
}

func EditAlertRequestTriggerScanRequestsSize(triggerScanRequestsSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerScanRequestsSize = &triggerScanRequestsSize
	}
}

func EditAlertRequestTriggerAdAnimationMaxLength(triggerAdAnimationMaxLength int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdAnimationMaxLength = &triggerAdAnimationMaxLength
	}
}

func EditAlertRequestTriggerSslCertViolation(triggerSslCertViolation int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerSslCertViolation = &triggerSslCertViolation
	}
}

func EditAlertRequestTriggerAdCookiesNetworks(triggerAdCookiesNetworks []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdCookiesNetworksStr := strings.Join(triggerAdCookiesNetworks, ",")
		request.TriggerAdCookiesNetworks = &triggerAdCookiesNetworksStr
	}
}

func EditAlertRequestTriggerAdCookiesDomains(triggerAdCookiesDomains []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdCookiesDomainsStr := strings.Join(triggerAdCookiesDomains, ",")
		request.TriggerAdCookiesDomains = &triggerAdCookiesDomainsStr
	}
}

func EditAlertRequestTriggerAdNetworks(triggerAdNetworks []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdNetworksStr := strings.Join(triggerAdNetworks, ",")
		request.TriggerAdNetworks = &triggerAdNetworksStr
	}
}

func EditAlertRequestTriggerAdDomains(triggerAdDomains []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdDomainsStr := strings.Join(triggerAdDomains, ",")
		request.TriggerAdDomains = &triggerAdDomainsStr
	}
}

func EditAlertRequestTriggerLpDomains(triggerLpDomains []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerLpDomainsStr := strings.Join(triggerLpDomains, ",")
		request.TriggerLpDomains = &triggerLpDomainsStr
	}
}

func EditAlertRequestTriggerAdVideoInDisplay(triggerAdVideoInDisplay int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoInDisplay = &triggerAdVideoInDisplay
	}
}

func EditAlertRequestTriggerAdVideoMaxLength(triggerAdVideoMaxLength int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoMaxLength = &triggerAdVideoMaxLength
	}
}

func EditAlertRequestTriggerAdVideoSupportedTypes(triggerAdVideoSupportedTypes []string) editAlertRequestOption {
	return func(request *editAlertRequest) {
		triggerAdVideoSupportedTypesStr := strings.Join(triggerAdVideoSupportedTypes, ",")
		request.TriggerAdVideoSupportedTypes = &triggerAdVideoSupportedTypesStr
	}
}

func EditAlertRequestTriggerAdVideoVastFromStart(triggerAdVideoVastFromStart int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoVastFromStart = &triggerAdVideoVastFromStart
	}
}

func EditAlertRequestTriggerAdVideoAutoSound(triggerAdVideoAutoSound int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoAutoSound = &triggerAdVideoAutoSound
	}
}

func EditAlertRequestTriggerAdVideoVastMissingMediaFile(triggerAdVideoVastMissingMediaFile int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoVastMissingMediaFile = &triggerAdVideoVastMissingMediaFile
	}
}

func EditAlertRequestTriggerAdVideoAutoScroll(triggerAdVideoAutoScroll int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoAutoScroll = &triggerAdVideoAutoScroll
	}
}

func EditAlertRequestTriggerAdVideoFrameRate(triggerAdVideoFrameRate int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoFrameRate = &triggerAdVideoFrameRate
	}
}

func EditAlertRequestTriggerAdVideoBitRate(triggerAdVideoBitRate int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoBitRate = &triggerAdVideoBitRate
	}
}

func EditAlertRequestTriggerAdVideoCompanionAds(triggerAdVideoCompanionAds int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoCompanionAds = &triggerAdVideoCompanionAds
	}
}

func EditAlertRequestTriggerAdVideoFileSize(triggerAdVideoFileSize int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerAdVideoFileSize = &triggerAdVideoFileSize
	}
}

func EditAlertRequestTriggerVpaidFormat(triggerVpaidFormat int) editAlertRequestOption {
	return func(request *editAlertRequest) {
		request.TriggerVpaidFormat = &triggerVpaidFormat
	}
}
