package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListEmulationSetsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "emulation-sets",
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
        "emulation-sets": [
            {
                "id": 1,
                "description": "Desktop"
            },
            {
                "id": 2,
                "description": "Mobile iOS"
            },
            {
                "id": 3,
                "description": "Mobile Android"
            },
            {
                "id": 4,
                "description": "Mobile Windows"
            },
            {
                "id": 5,
                "description": "Mobile BlackBerry"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListEmulationSetsRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []emulationSet{
		{
			Id:          1,
			Description: "Desktop",
		},
		{
			Id:          2,
			Description: "Mobile iOS",
		},
		{
			Id:          3,
			Description: "Mobile Android",
		},
		{
			Id:          4,
			Description: "Mobile Windows",
		},
		{
			Id:          5,
			Description: "Mobile BlackBerry",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
