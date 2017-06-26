package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestTriggerTypesRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "alerts/trigger-types",
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
        "trigger-types": [
            {
                "id": 1,
                "key": "trigger_ad_specific_banner",
                "description": "Specific Ad"
            },
            {
                "id": 2,
                "key": "trigger_scan_malware",
                "description": "Malware In Page"
            },
            {
                "id": 3,
                "key": "trigger_ad_lp_offensive",
                "description": "Offensive Content"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.TriggerTypesRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []triggerType{
		{
			Id:          1,
			Key:         "trigger_ad_specific_banner",
			Description: "Specific Ad",
		},
		{
			Id:          2,
			Key:         "trigger_scan_malware",
			Description: "Malware In Page",
		},
		{
			Id:          3,
			Key:         "trigger_ad_lp_offensive",
			Description: "Offensive Content",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
