package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestScanStatusRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "projects/3460244ac44e78c09ece2d24519cf71e/scan-status",
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
        "status": "done",
        "num_of_ads": 2
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ScanStatusRequest("3460244ac44e78c09ece2d24519cf71e").Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := scanStatusResponce{
		Status:   "done",
		NumOfAds: 2,
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}

func TestScanStatusRequestWrongNumOfAds(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "projects/3460244ac44e78c09ece2d24519cf71e/scan-status",
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
        "status": "pending",
        "num_of_ads": "-1"
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ScanStatusRequest("3460244ac44e78c09ece2d24519cf71e").Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := scanStatusResponce{
		Status:   "pending",
		NumOfAds: -1,
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
