package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListEmulatorsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "emulators",
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
        "emulators": [
            {
                "id": "90",
                "description": "iPad iOS 8",
                "category": "Apple iOS"
            },
            {
                "id": "67",
                "description": "iPad iOS 7.1",
                "category": "Apple iOS"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListEmulatorsRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []emulator{
		{
			Id:          "90",
			Description: "iPad iOS 8",
			Category:    "Apple iOS",
		},
		{
			Id:          "67",
			Description: "iPad iOS 7.1",
			Category:    "Apple iOS",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
