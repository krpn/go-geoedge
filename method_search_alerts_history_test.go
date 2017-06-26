package geoedge

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestSearchAlertsHistoryRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "alerts/history?alert_id=3aae8182cc771891a0b0323ea4a91fa2&limit=50&location_id=AU%2CCA%2CUA&max_datetime=2016-01-01+00%3A00%3A00&min_datetime=2014-08-20+14%3A00%3A00&offset=100&trigger_type_id=11",
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
        "prev_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `alerts\/history?alert_id=3aae8182cc771891a0b0323ea4a91fa2&trigger_type_id=11&min_datetime=2014-08-20+14%3A00%3A00&max_datetime=2016-01-01+00%3A00%3A00&location_id=AU,CA,UA&offset=50&limit=50",
        "next_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `alerts\/history?alert_id=3aae8182cc771891a0b0323ea4a91fa2&trigger_type_id=11&min_datetime=2014-08-20+14%3A00%3A00&max_datetime=2016-01-01+00%3A00%3A00&location_id=AU,CA,UA&offset=150&limit=50",
        "alerts": [
            {
                "alert_id": "b389fbcbbf3bec78c06ab71750902a9b",
                "history_id": "i3uUcn9keJ9UpsREBFaSaA",
                "trigger_type_id": "29",
                "trigger_metadata": "14.14 secs",
                "alert_name": "Test Alert 2",
                "event_datetime": "2015-07-09 03:03:27",
                "location": {
                    "AU": "Australia"
                },
                "ad_id": "7iPbOgeHNBkeHSHaUy7RfA",
                "alert_details_url": "http:\/\/www.geoedge.com\/analyticsv2\/ad\/1009777086\/750e1f8860a93d9373b6ea1a3d4f94a1\/alert:115906313626049160",
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
                    "1a755a2adf6301380b5ed35fb303767c": "Test Project 2"
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

	minDatetime, _ := time.Parse("2006-01-02 15:04:05", "2014-08-20 14:00:00")
	maxDatetime, _ := time.Parse("2006-01-02 15:04:05", "2016-01-01 00:00:00")

	result, prevPage, nextPage, err := api.SearchAlertsHistoryRequest(
		SearchAlertsHistoryRequestAlertId("3aae8182cc771891a0b0323ea4a91fa2"),
		SearchAlertsHistoryRequestTriggerTypeId("11"),
		SearchAlertsHistoryRequestMinDatetime(minDatetime),
		SearchAlertsHistoryRequestMaxDatetime(maxDatetime),
		SearchAlertsHistoryRequestLocationId([]string{"AU", "CA", "UA"}),
		SearchAlertsHistoryRequestOffset(100),
		SearchAlertsHistoryRequestLimit(50),
	).Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []alertHistory{
		{
			AlertId:         "b389fbcbbf3bec78c06ab71750902a9b",
			HistoryId:       "i3uUcn9keJ9UpsREBFaSaA",
			TriggerTypeId:   "29",
			TriggerMetadata: "14.14 secs",
			AlertName:       "Test Alert 2",
			EventDatetime:   "2015-07-09 03:03:27",
			Location: map[string]string{
				"AU": "Australia",
			},
			AdId:            "7iPbOgeHNBkeHSHaUy7RfA",
			AlertDetailsUrl: "http://www.geoedge.com/analyticsv2/ad/1009777086/750e1f8860a93d9373b6ea1a3d4f94a1/alert:115906313626049160",
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
				"1a755a2adf6301380b5ed35fb303767c": "Test Project 2",
			},
			Tag: map[string]string{
				"l3jCl0fQmzo": "http://www.yahoo.com/",
			},
		},
	}
	expectedPrevPage := apiUrl + "alerts/history?alert_id=3aae8182cc771891a0b0323ea4a91fa2&trigger_type_id=11&min_datetime=2014-08-20+14%3A00%3A00&max_datetime=2016-01-01+00%3A00%3A00&location_id=AU,CA,UA&offset=50&limit=50"
	expectedNextPage := apiUrl + "alerts/history?alert_id=3aae8182cc771891a0b0323ea4a91fa2&trigger_type_id=11&min_datetime=2014-08-20+14%3A00%3A00&max_datetime=2016-01-01+00%3A00%3A00&location_id=AU,CA,UA&offset=150&limit=50"

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
	if prevPage != expectedPrevPage {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedPrevPage, prevPage)
	}
	if nextPage != expectedNextPage {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedNextPage, nextPage)
	}
}
