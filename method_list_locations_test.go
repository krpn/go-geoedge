package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListLocationsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "locations",
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
        "locations": [
            {
                "id": "Y0",
                "description": "Atlanta, GA",
                "region": "United States - Cities"
            },
            {
                "id": "E7",
                "description": "Austria, A1",
                "region": "Mobile Carriers"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListLocationsRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []location{
		{
			Id:          "Y0",
			Description: "Atlanta, GA",
			Region:      "United States - Cities",
		},
		{
			Id:          "E7",
			Description: "Austria, A1",
			Region:      "Mobile Carriers",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
