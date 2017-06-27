package geoedge

import "encoding/json"

type responceStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type rawResponse struct {
	Status   responceStatus  `json:"status"`
	Response json.RawMessage `json:"response"`
}

type projectBase struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	AutoScan     int    `json:"auto_scan"`
	CreationTime int    `json:"creation_time"`
}

type project struct {
	projectBase
	LastRunTime   int               `json:"last_run_time"`
	ExtLineitemId string            `json:"ext_lineitem_id"`
	ScanType      int               `json:"scan_type"`
	EmulationType string            `json:"emulation_type"`
	Locations     map[string]string `json:"locations"`
	EmulationSets map[string]string `json:"emulation_sets"`
	Emulators     map[string]string `json:"emulators"`
	Tags          map[string]struct {
		Tag           string `json:"tag"`
		ExtCreativeId string `json:"ext_creative_id"`
	} `json:"tags"`
	TopUrls     string `json:"top_urls"`
	TimesPerDay int    `json:"times_per_day"`
}

type alert struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	Enabled          string            `json:"enabled"`
	GlobalAlert      string            `json:"global_alert"`
	NotificationTime string            `json:"notification_time"`
	Emails           []string          `json:"emails"`
	Locations        map[string]string `json:"locations"`
	Triggers         [][]struct {
		TriggerTypeId string   `json:"trigger_type_id"`
		Params        []string `json:"params"`
	} `json:"triggers"`
	Projects []map[string]string `json:"projects"`
}

type alertCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type triggerType struct {
	Id          int    `json:"id"` // not as in specification (string), specification is wrong
	Key         string `json:"key"`
	Description string `json:"description"`
}

type alertDetailsInfoBase struct {
	TypeStr string `json:"type_str"`
	Url     string `json:"url"`
}

type alertDetailsTree struct {
	InitiatedBy string `json:"initiated_by"`
	Dom         []struct {
		Item struct {
			Tag      string `json:"tag"`
			UrlClass string `json:"url_class"`
			UrlTitle string `json:"url_title"`
			Url      string `json:"url"`
		} `json:"item"`
		Subitems json.RawMessage `json:"subitems"` // unknown structure
	} `json:"dom"`
	Sequence struct {
		Analytics []struct {
			Highlight bool   `json:"highlight"`
			Cause     string `json:"cause"`
			Url       string `json:"url"`
			Referrer  string `json:"referrer"`
		} `json:"analytics"`
	} `json:"sequence"`
}

type alertDetailsFileInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size string `json:"size"`
	Mime string `json:"mime"`
}

type alertDetailsMalware struct {
	Info struct {
		alertDetailsInfoBase
		UrlSrc        string               `json:"url_src"`
		MaliciousSite bool                 `json:"malicious_site"`
		FileInfo      alertDetailsFileInfo `json:"file_info"`
		ReviewTime    string               `json:"review_time"`
		Vendors       []string             `json:"vendors"`
	} `json:"info"`
	Tree alertDetailsTree `json:"tree"`
}

type alertDetailsPhishing struct {
	Info struct {
		alertDetailsInfoBase
		UrlSrc     string   `json:"url_src"`
		ReviewTime string   `json:"review_time"`
		Vendors    []string `json:"vendors"`
	} `json:"info"`
	Tree alertDetailsTree `json:"tree"`
}

type alertDetailsFileDownload struct {
	Info struct {
		alertDetailsInfoBase
		UrlSrc        string               `json:"url_src"`
		MaliciousSite bool                 `json:"malicious_site"`
		FileInfo      alertDetailsFileInfo `json:"file_info"`
	} `json:"info"`
	Tree alertDetailsTree `json:"tree"`
}

type alertDetailsMalicious struct {
	Info struct {
		alertDetailsInfoBase
		UrlSrc     string `json:"url_src"`
		MetaEngine string `json:"meta_engine"`
	} `json:"info"`
	Tree alertDetailsTree `json:"tree"`
}

type alertDetailsRedirect struct {
	Info struct {
		alertDetailsInfoBase
		Title      string `json:"title"`
		Screenshot string `json:"screenshot"`
	} `json:"info"`
	Tree alertDetailsTree `json:"tree"`
}

type alertDetails struct {
	Malware              alertDetailsMalware      `json:"malware"`
	Phishing             alertDetailsPhishing     `json:"phishing"`
	FileDownload         alertDetailsFileDownload `json:"file_download"`
	Malicious            alertDetailsMalicious    `json:"malicious"`
	Redirect             alertDetailsRedirect     `json:"redirect"`
	MobileRedirect       alertDetailsRedirect     `json:"mobile_redirect"`
	ImpactString         string                   `json:"impact_string"`
	RecommendationString string                   `json:"recommendation_string"`
}

type alertHistory struct {
	AlertId         string            `json:"alert_id"`
	HistoryId       string            `json:"history_id"`
	TriggerTypeId   string            `json:"trigger_type_id"`
	TriggerMetadata string            `json:"trigger_metadata"`
	AlertName       string            `json:"alert_name"`
	EventDatetime   string            `json:"event_datetime"`
	Location        map[string]string `json:"location"`
	AdId            string            `json:"ad_id"`
	AlertDetailsUrl string            `json:"alert_details_url"`
	AlertDetails    alertDetails      `json:"alert_details"`
	ProjectName     map[string]string `json:"project_name"`
	Tag             map[string]string `json:"tag"`
}

type adNetwork struct {
	NetworkId           string `json:"network_id"`
	NetworkName         string `json:"network_name"`
	NetworkType         string `json:"network_type"`
	NetworkUrl          string `json:"network_url"`
	NetworkIconType     string `json:"network_icon_type"`
	NetworkLogoType     string `json:"network_logo_type"`
	NetworkDesc         string `json:"network_desc"`
	NetworkTooltip      string `json:"network_tooltip"`
	NetworkPrivacyUrl   string `json:"network_privacy_url"`
	IsMemberIAB         string `json:"is_member_iab"`
	IsMemberNAI         string `json:"is_member_nai"`
	IsAdult             string `json:"is_adult"`
	IsMobile            string `json:"is_mobile"`
	IsVideo             string `json:"is_video"`
	NetworkInternalName string `json:"network_internal_name"`
}

type ad struct {
	Id                       string            `json:"ad_id"`
	ProjectName              map[string]string `json:"project_name"`
	Emulator                 map[string]string `json:"emulator"`
	CaptureDate              string            `json:"capture_date"`
	Location                 map[string]string `json:"location"`
	ScanScreenshotUrl        string            `json:"scan_screenshot_url"`
	CreativeUrl              string            `json:"creative_url"`
	LandingPath              []string          `json:"landing_path"`
	CreativeScreenshotUrl    string            `json:"creative_screenshot_url"`
	LandingPageUrl           string            `json:"landing_page_url"`
	LandingPageScreenshotUrl string            `json:"landing_page_screenshot_url"`
	MoreInfoLink             string            `json:"more_info_link"`
	AdUrls                   []string          `json:"ad_urls"`
	Type                     string            `json:"type"`
	Network                  adNetwork         `json:"network"`
}

type location struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Region      string `json:"region"`
}

type platrofm struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type emulationSet struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type emulator struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type network struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type usage struct {
	DailyScans                   string `json:"daily_scans"`
	DailyScansAutomatic          string `json:"daily_scans_automatic"`
	DailyScansManual             string `json:"daily_scans_manual"`
	MonthlyScans                 string `json:"monthly_scans"`
	MonthlyScansAutomatic        string `json:"monthly_scans_automatic"`
	MonthlyScansManual           string `json:"monthly_scans_manual"`
	BillingMonthlyScans          string `json:"billing_monthly_scans"`
	BillingMonthlyScansAutomatic string `json:"billing_monthly_scans_automatic"`
	BillingMonthlyScansManual    string `json:"billing_monthly_scans_manual"`
}
