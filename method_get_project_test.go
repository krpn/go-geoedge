package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetProjectRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "projects/1a755a2adf6301380b5ed35fb303767c",
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
        "project": {
            "id": "1a755a2adf6301380b5ed35fb303767c",
            "name": "Test Project 2",
            "creation_time": 1431984877,
            "last_run_time": 1436868125,
            "ext_lineitem_id": "",
            "auto_scan": 1,
            "scan_type": 1,
            "emulation_type": "1",
            "locations": {
                "US": "United States"
            },
            "emulation_sets": [],
            "emulators": {
                "67": "iPad iOS 7.1"
            },
            "tags": {
                "l3jCl0fQmzo": {
                    "tag": "http:\/\/www.yahoo.com\/",
                    "ext_creative_id": ""
                }
            },
            "top_urls": "http:\/\/www.yahoo.com",
            "times_per_day": 4
        }
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.GetProjectRequest("1a755a2adf6301380b5ed35fb303767c").Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := project{
		projectBase: projectBase{
			Id:           "1a755a2adf6301380b5ed35fb303767c",
			Name:         "Test Project 2",
			AutoScan:     1,
			CreationTime: 1431984877,
		},
		LastRunTime:   1436868125,
		ExtLineitemId: "",
		ScanType:      1,
		EmulationType: "1",
		Locations: map[string]string{
			"US": "United States",
		},
		EmulationSets: map[string]string{},
		Emulators: map[string]string{
			"67": "iPad iOS 7.1",
		},
		Tags: map[string]struct {
			Tag           string `json:"tag"`
			ExtCreativeId string `json:"ext_creative_id"`
		}{
			"l3jCl0fQmzo": {
				Tag:           "http://www.yahoo.com/",
				ExtCreativeId: "",
			},
		},
		TopUrls:     "http://www.yahoo.com",
		TimesPerDay: 4,
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
