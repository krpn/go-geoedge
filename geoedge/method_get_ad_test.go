package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetAdRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "ads/7iPbOgeHNBkeHSHaUy7RfA",
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
        "ad": {
            "ad_id": "7iPbOgeHNBkeHSHaUy7RfA",
            "project_name": {
                "1a755a2adf6301380b5ed35fb303767c": "Test Project 2"
            },
            "emulator": {
                "67": "iPad iOS 7.1"
            },
            "capture_date": "2015-07-09",
            "location": {
                "AU": "Australia"
            },
            "scan_screenshot_url": "https:\/\/geoedge-analytics.s3.amazonaws.com\/screenshots\/f5\/f1\/screenshot_f5f18ebea41bf42c1f415a5e9104e844.jpg",
            "creative_url": "https:\/\/content-ssl.yieldmanager.com\/atoms\/7a\/20\/88\/e6\/7a2088e6746a3b0b1e65ab19322e0e04.jpg",
            "landing_path": [
                "https:\/\/ads.yahoo.com\/clk?3,eJytjdFugkAURL.GNzRclhXIpg8XKFYDGGVXhTe6VlaU0AYM6tcrqTF-gCeTzGRybwYIcxx9Bxa1trlJDfltMzB.KBCHSNhpOmMMbDLWxxYBoh1k6Hti87XqPP.UCQ97QlduO3whfXiDWODEU1ZSFv9NmQ45vgU3qZfNI0.7HSx8G..8hgphus-zybKKqtk-LsUl5NNzxNM24sExTkBl608acrWPr8EhWy8gNQSdL56fH5qm2va3GRAcGMFd-Wl0yVVdj2Rd3QC8gVg2,",
                "http:\/\/www.rmhcnswballraffle.org\/"
            ],
            "creative_screenshot_url": "https:\/\/analytics-fs7.geoedge.com\/snapshots\/objects\/7a\/20\/object_7a2088e6746a3b0b1e65ab19322e0e04.jpg",
            "landing_page_url": "http:\/\/www.rmhcnswballraffle.org\/",
            "landing_page_screenshot_url": "https:\/\/analytics-fs5.geoedge.com\/snapshots\/screenshots\/15\/f5\/landing_15f5e3bb38aee0fefb96fda967c3ff64.jpg",
            "more_info_link": "http:\/\/www.geoedge.com\/analyticsv2\/ad\/1009777086\/750e1f8860a93d9373b6ea1a3d4f94a1",
            "ad_urls": [
                "https:\/\/ads.yahoo.com\/clk?3,eJytjdFugkAURL.GNzRclhXIpg8XKFYDGGVXhTe6VlaU0AYM6tcrqTF-gCeTzGRybwYIcxx9Bxa1trlJDfltMzB.KBCHSNhpOmMMbDLWxxYBoh1k6Hti87XqPP.UCQ97QlduO3whfXiDWODEU1ZSFv9NmQ45vgU3qZfNI0.7HSx8G..8hgphus-zybKKqtk-LsUl5NNzxNM24sExTkBl608acrWPr8EhWy8gNQSdL56fH5qm2va3GRAcGMFd-Wl0yVVdj2Rd3QC8gVg2,",
                "http:\/\/www.rmhcnswballraffle.org\/"
            ],
            "type": "728x90 - Image",
            "network": {
                "network_id": "138",
                "network_name": "Internal Serving",
                "network_type": "0",
                "network_url": "http:\/\/no-website.com",
                "network_icon_type": "",
                "network_logo_type": "",
                "network_desc": "Ad serving done by an adserver on the site&#039;s domain",
                "network_tooltip": "Ad serving done by an adserver on the site&#039;s domain",
                "network_privacy_url": "",
                "is_member_iab": "0",
                "is_member_nai": "0",
                "is_adult": "0",
                "is_mobile": "0",
                "is_video": "0",
                "network_internal_name": ""
            }
        }
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.GetAdRequest("7iPbOgeHNBkeHSHaUy7RfA").Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := ad{
		Id: "7iPbOgeHNBkeHSHaUy7RfA",
		ProjectName: map[string]string{
			"1a755a2adf6301380b5ed35fb303767c": "Test Project 2",
		},
		Emulator: map[string]string{
			"67": "iPad iOS 7.1",
		},
		CaptureDate: "2015-07-09",
		Location: map[string]string{
			"AU": "Australia",
		},
		ScanScreenshotUrl: "https://geoedge-analytics.s3.amazonaws.com/screenshots/f5/f1/screenshot_f5f18ebea41bf42c1f415a5e9104e844.jpg",
		CreativeUrl:       "https://content-ssl.yieldmanager.com/atoms/7a/20/88/e6/7a2088e6746a3b0b1e65ab19322e0e04.jpg",
		LandingPath: []string{
			"https://ads.yahoo.com/clk?3,eJytjdFugkAURL.GNzRclhXIpg8XKFYDGGVXhTe6VlaU0AYM6tcrqTF-gCeTzGRybwYIcxx9Bxa1trlJDfltMzB.KBCHSNhpOmMMbDLWxxYBoh1k6Hti87XqPP.UCQ97QlduO3whfXiDWODEU1ZSFv9NmQ45vgU3qZfNI0.7HSx8G..8hgphus-zybKKqtk-LsUl5NNzxNM24sExTkBl608acrWPr8EhWy8gNQSdL56fH5qm2va3GRAcGMFd-Wl0yVVdj2Rd3QC8gVg2,",
			"http://www.rmhcnswballraffle.org/",
		},
		CreativeScreenshotUrl:    "https://analytics-fs7.geoedge.com/snapshots/objects/7a/20/object_7a2088e6746a3b0b1e65ab19322e0e04.jpg",
		LandingPageUrl:           "http://www.rmhcnswballraffle.org/",
		LandingPageScreenshotUrl: "https://analytics-fs5.geoedge.com/snapshots/screenshots/15/f5/landing_15f5e3bb38aee0fefb96fda967c3ff64.jpg",
		MoreInfoLink:             "http://www.geoedge.com/analyticsv2/ad/1009777086/750e1f8860a93d9373b6ea1a3d4f94a1",
		AdUrls: []string{
			"https://ads.yahoo.com/clk?3,eJytjdFugkAURL.GNzRclhXIpg8XKFYDGGVXhTe6VlaU0AYM6tcrqTF-gCeTzGRybwYIcxx9Bxa1trlJDfltMzB.KBCHSNhpOmMMbDLWxxYBoh1k6Hti87XqPP.UCQ97QlduO3whfXiDWODEU1ZSFv9NmQ45vgU3qZfNI0.7HSx8G..8hgphus-zybKKqtk-LsUl5NNzxNM24sExTkBl608acrWPr8EhWy8gNQSdL56fH5qm2va3GRAcGMFd-Wl0yVVdj2Rd3QC8gVg2,",
			"http://www.rmhcnswballraffle.org/",
		},
		Type: "728x90 - Image",
		Network: adNetwork{
			NetworkId:      "138",
			NetworkName:    "Internal Serving",
			NetworkType:    "0",
			NetworkUrl:     "http://no-website.com",
			NetworkDesc:    "Ad serving done by an adserver on the site&#039;s domain",
			NetworkTooltip: "Ad serving done by an adserver on the site&#039;s domain",
			IsMemberIAB:    "0",
			IsMemberNAI:    "0",
			IsAdult:        "0",
			IsMobile:       "0",
			IsVideo:        "0",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
