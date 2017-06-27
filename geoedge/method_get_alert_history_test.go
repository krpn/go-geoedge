package geoedge

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

func TestGetAlertHistoryRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "alertshistory/AuQh8Tnyk0QrDgEq0SzyBw",
				Body:   nil,
				Headers: map[string]string{
					"Host":          parsedUrl.Host,
					"Authorization": "testToken",
				},
				ResultBody: []byte(`{
    "status": {
        "code": "Success",
        "message": "Success"
    },
    "response": {
        "alerts": [
            {
                "alert_id": "2e2d86b1bfda314c2496011d0426e786",
                "history_id": "AuQh8Tnyk0QrDgEq0SzyBw",
                "trigger_type_id": "4",
                "trigger_metadata": "Finance & Insurance",
                "alert_name": "travel",
                "event_datetime": "2015-07-14 03:04:20",
                "location": {
                    "US": "United States"
                },
                "ad_id": "CElTSt-65Ogan3lnzzwjHg",
                "alert_details_url": "http:\/\/www.geoedge.com\/analyticsv2\/ad\/1021141477\/30c43b41bc9d517d9ebb4455ee7be45d\/alert:118084714624395932",
                "alert_details": {
                    "redirect": {
                        "info": {
                            "type_str": "Automatic Redirect",
                            "url": "http:\/\/example.site\/example\/example.html",
                            "title": "Title example",
                            "screenshot": "https:\/\/geoedge-analytics.s3.amazonaws.com\/images\/screenshots\/screenshots\/19\/68\/landing_19680c655da64183ea8dd3bae253cdb1.jpg"
                        },
                        "tree": {
                            "initiated_by": "* Redirect to URL initiated by a javascript code in: http:\/\/example.site\/example\/example.html",
                            "dom": [
                                {
                                    "item": {
                                        "tag": "IFRAME",
                                        "url_class": "malware",
                                        "url_title": "Suspicious Redirect Started Here",
                                        "url": "http:\/\/example.site\/example\/example.html"
                                    },
                                    "subitems": []
                                }
                            ],
                            "sequence": {
                                "analytics": [
                                    {
                                        "highlight": true,
                                        "cause": "JS.redirect",
                                        "url": "http:\/\/example.site\/example\/example.html",
                                        "referrer": "http:\/\/example.site\/"
                                    }
                                ]
                            }
                        }
                    },
                    "mobile_redirect": {
                        "info": {
                            "type_str": "Automatic Hidden Mobile Redirect",
                            "url": "https:\/\/play.google.com\/store\/apps\/details?id=com.example.example",
                            "title": "Example Title",
                            "screenshot": "http:\/\/analytics-fs8.geoedge.com\/images\/screenshots\/screenshots\/9f\/4d\/landing_9f4d99add5c750fd5ff7a40612555555.jpg"
                        },
                        "tree": {
                            "initiated_by": "* Redirect to URL initiated by a javascript code in: http:\/\/example.site\/example\/example.html",
                            "dom": [
                                {
                                    "item": {
                                        "tag": "IFRAME",
                                        "url_class": "malware",
                                        "url_title": "Suspicious Redirect Started Here",
                                        "url": "http:\/\/example.site\/example\/example.html"
                                    },
                                    "subitems": []
                                }
                            ],
                            "sequence": {
                                "analytics": [
                                    {
                                        "highlight": true,
                                        "cause": "JS.redirect",
                                        "url": "http:\/\/example.site\/example\/example.html",
                                        "referrer": "http:\/\/example.site\/"
                                    }
                                ]
                            }
                        }
                    },
                    "malware": {
                        "info": {
                            "type_str": "Automatically Downloaded Virus File",
                            "url": "http:\/\/example.site\/upload\/s2s\/polyvideo_6122.apk",
                            "url_src": "Landing page",
                            "malicious_site": false,
                            "file_info": {
                                "name": "polyvideo_6122.apk",
                                "type": "Android application package file",
                                "size": "634.5KB",
                                "mime": "application\/vnd.android.package-archive"
                            },
                            "review_time": "Apr 5, 2017 08:26",
                            "vendors": [
                                "Ahnlab - Android-Trojan\/Hidap.3cbc2",
                                "BitDefender - Android.Trojan.Dropper.FX",
                                "ESET - a variant of Android\/AdDisplay.Sprovider.X application",
                                "QuickHeal - Android.Sprovider.F (AdWare)"
                            ]
                        },
                        "tree": []
                    },
                    "phishing": {
                        "info": {
                            "type_str": "Phishing Information",
                            "url": "http:\/\/example.site\/?refId=TMPTM-112296&data=TM-Optimiser",
                            "url_src": "Landing page",
                            "review_time": "Apr 21, 2017 00:49",
                            "vendors": [
                                "Google Safe Browsing - googpub-phish-shavar"
                            ]
                        },
                        "tree": []
                    },
                    "malicious": {
                        "info": {
                            "type_str": "Malicious Information",
                            "url": "http:\/\/example.site\/clk?v=u3I2YTUejQzi8%2Fq%",
                            "url_src": "In redirect path to landing page",
                            "meta_engine": "Malicious Domain"
                        },
                        "tree": []
                    },
                    "file_download": {
                        "info": {
                            "type_str": "Automatically Downloaded File",
                            "url": "http:\/\/example.site\/play\/sd\/Hot_xxx_Videos.apk",
                            "url_src": "Landing page",
                            "malicious_site": false,
                            "file_info": {
                                "name": "Hot_xxx_Videos.apk",
                                "type": "Android application package file",
                                "size": "2.6MB",
                                "mime": "application\/vnd.android.package-archive"
                            }
                        },
                        "tree": []
                    },
                    "impact_string": "Bad User Experience - All users will be impacted",
                    "recommendation_string": "Be aware - act according to policy"
                },
                "project_name": {
                    "1a755a2adf6301380b5ed35fb303767c": "Example Project"
                },
                "tag": {
                    "l3jCl0fQmzo": "http:\/\/www.yahoo.com\/"
                }
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.GetAlertHistoryRequest("AuQh8Tnyk0QrDgEq0SzyBw").Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []alertHistory{
		{
			AlertId:         "2e2d86b1bfda314c2496011d0426e786",
			HistoryId:       "AuQh8Tnyk0QrDgEq0SzyBw",
			TriggerTypeId:   "4",
			TriggerMetadata: "Finance & Insurance",
			AlertName:       "travel",
			EventDatetime:   "2015-07-14 03:04:20",
			Location: map[string]string{
				"US": "United States",
			},
			AdId:            "CElTSt-65Ogan3lnzzwjHg",
			AlertDetailsUrl: "http://www.geoedge.com/analyticsv2/ad/1021141477/30c43b41bc9d517d9ebb4455ee7be45d/alert:118084714624395932",
			AlertDetails: alertDetails{
				Redirect: alertDetailsRedirect{
					Info: struct {
						alertDetailsInfoBase
						Title      string `json:"title"`
						Screenshot string `json:"screenshot"`
					}{
						alertDetailsInfoBase: alertDetailsInfoBase{
							TypeStr: "Automatic Redirect",
							Url:     "http://example.site/example/example.html",
						},
						Title:      "Title example",
						Screenshot: "https://geoedge-analytics.s3.amazonaws.com/images/screenshots/screenshots/19/68/landing_19680c655da64183ea8dd3bae253cdb1.jpg",
					},
					Tree: alertDetailsTree{
						InitiatedBy: "* Redirect to URL initiated by a javascript code in: http://example.site/example/example.html",
						Dom: []struct {
							Item struct {
								Tag      string `json:"tag"`
								UrlClass string `json:"url_class"`
								UrlTitle string `json:"url_title"`
								Url      string `json:"url"`
							} `json:"item"`
							Subitems json.RawMessage `json:"subitems"`
						}{
							{
								Item: struct {
									Tag      string `json:"tag"`
									UrlClass string `json:"url_class"`
									UrlTitle string `json:"url_title"`
									Url      string `json:"url"`
								}{
									Tag:      "IFRAME",
									UrlClass: "malware",
									UrlTitle: "Suspicious Redirect Started Here",
									Url:      "http://example.site/example/example.html",
								},
								Subitems: []byte(`[]`),
							},
						},
						Sequence: struct {
							Analytics []struct {
								Highlight bool   `json:"highlight"`
								Cause     string `json:"cause"`
								Url       string `json:"url"`
								Referrer  string `json:"referrer"`
							} `json:"analytics"`
						}{
							Analytics: []struct {
								Highlight bool   `json:"highlight"`
								Cause     string `json:"cause"`
								Url       string `json:"url"`
								Referrer  string `json:"referrer"`
							}{
								{
									Highlight: true,
									Cause:     "JS.redirect",
									Url:       "http://example.site/example/example.html",
									Referrer:  "http://example.site/",
								},
							},
						},
					},
				},
				MobileRedirect: alertDetailsRedirect{
					Info: struct {
						alertDetailsInfoBase
						Title      string `json:"title"`
						Screenshot string `json:"screenshot"`
					}{
						alertDetailsInfoBase: alertDetailsInfoBase{
							TypeStr: "Automatic Hidden Mobile Redirect",
							Url:     "https://play.google.com/store/apps/details?id=com.example.example",
						},
						Title:      "Example Title",
						Screenshot: "http://analytics-fs8.geoedge.com/images/screenshots/screenshots/9f/4d/landing_9f4d99add5c750fd5ff7a40612555555.jpg",
					},
					Tree: alertDetailsTree{
						InitiatedBy: "* Redirect to URL initiated by a javascript code in: http://example.site/example/example.html",
						Dom: []struct {
							Item struct {
								Tag      string `json:"tag"`
								UrlClass string `json:"url_class"`
								UrlTitle string `json:"url_title"`
								Url      string `json:"url"`
							} `json:"item"`
							Subitems json.RawMessage `json:"subitems"`
						}{
							{
								Item: struct {
									Tag      string `json:"tag"`
									UrlClass string `json:"url_class"`
									UrlTitle string `json:"url_title"`
									Url      string `json:"url"`
								}{
									Tag:      "IFRAME",
									UrlClass: "malware",
									UrlTitle: "Suspicious Redirect Started Here",
									Url:      "http://example.site/example/example.html",
								},
								Subitems: []byte(`[]`),
							},
						},
						Sequence: struct {
							Analytics []struct {
								Highlight bool   `json:"highlight"`
								Cause     string `json:"cause"`
								Url       string `json:"url"`
								Referrer  string `json:"referrer"`
							} `json:"analytics"`
						}{
							Analytics: []struct {
								Highlight bool   `json:"highlight"`
								Cause     string `json:"cause"`
								Url       string `json:"url"`
								Referrer  string `json:"referrer"`
							}{
								{
									Highlight: true,
									Cause:     "JS.redirect",
									Url:       "http://example.site/example/example.html",
									Referrer:  "http://example.site/",
								},
							},
						},
					},
				},
				Malware: alertDetailsMalware{
					Info: struct {
						alertDetailsInfoBase
						UrlSrc        string               `json:"url_src"`
						MaliciousSite bool                 `json:"malicious_site"`
						FileInfo      alertDetailsFileInfo `json:"file_info"`
						ReviewTime    string               `json:"review_time"`
						Vendors       []string             `json:"vendors"`
					}{
						alertDetailsInfoBase: alertDetailsInfoBase{
							TypeStr: "Automatically Downloaded Virus File",
							Url:     "http://example.site/upload/s2s/polyvideo_6122.apk",
						},
						UrlSrc:        "Landing page",
						MaliciousSite: false,
						FileInfo: alertDetailsFileInfo{
							Name: "polyvideo_6122.apk",
							Type: "Android application package file",
							Size: "634.5KB",
							Mime: "application/vnd.android.package-archive",
						},
						ReviewTime: "Apr 5, 2017 08:26",
						Vendors: []string{
							"Ahnlab - Android-Trojan/Hidap.3cbc2",
							"BitDefender - Android.Trojan.Dropper.FX",
							"ESET - a variant of Android/AdDisplay.Sprovider.X application",
							"QuickHeal - Android.Sprovider.F (AdWare)",
						},
					},
					Tree: alertDetailsTree{},
				},
				Phishing: alertDetailsPhishing{
					Info: struct {
						alertDetailsInfoBase
						UrlSrc     string   `json:"url_src"`
						ReviewTime string   `json:"review_time"`
						Vendors    []string `json:"vendors"`
					}{
						alertDetailsInfoBase: alertDetailsInfoBase{
							TypeStr: "Phishing Information",
							Url:     "http://example.site/?refId=TMPTM-112296&data=TM-Optimiser",
						},
						UrlSrc:     "Landing page",
						ReviewTime: "Apr 21, 2017 00:49",
						Vendors: []string{
							"Google Safe Browsing - googpub-phish-shavar",
						},
					},
					Tree: alertDetailsTree{},
				},
				Malicious: alertDetailsMalicious{
					Info: struct {
						alertDetailsInfoBase
						UrlSrc     string `json:"url_src"`
						MetaEngine string `json:"meta_engine"`
					}{
						alertDetailsInfoBase: alertDetailsInfoBase{
							TypeStr: "Malicious Information",
							Url:     "http://example.site/clk?v=u3I2YTUejQzi8%2Fq%",
						},
						UrlSrc:     "In redirect path to landing page",
						MetaEngine: "Malicious Domain",
					},
					Tree: alertDetailsTree{},
				},
				FileDownload: alertDetailsFileDownload{
					Info: struct {
						alertDetailsInfoBase
						UrlSrc        string               `json:"url_src"`
						MaliciousSite bool                 `json:"malicious_site"`
						FileInfo      alertDetailsFileInfo `json:"file_info"`
					}{
						alertDetailsInfoBase: alertDetailsInfoBase{
							TypeStr: "Automatically Downloaded File",
							Url:     "http://example.site/play/sd/Hot_xxx_Videos.apk",
						},
						UrlSrc:        "Landing page",
						MaliciousSite: false,
						FileInfo: alertDetailsFileInfo{
							Name: "Hot_xxx_Videos.apk",
							Type: "Android application package file",
							Size: "2.6MB",
							Mime: "application/vnd.android.package-archive",
						},
					},
					Tree: alertDetailsTree{},
				},
				ImpactString:         "Bad User Experience - All users will be impacted",
				RecommendationString: "Be aware - act according to policy",
			},
			ProjectName: map[string]string{
				"1a755a2adf6301380b5ed35fb303767c": "Example Project",
			},
			Tag: map[string]string{
				"l3jCl0fQmzo": "http://www.yahoo.com/",
			},
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
