package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListPlatformsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "platforms",
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
        "platforms": [
            {
                "id": "8",
                "name": "Desktop Browsers"
            },
            {
                "id": "3",
                "name": "Android"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListPlatformsRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []platrofm{
		{
			Id:   "8",
			Name: "Desktop Browsers",
		},
		{
			Id:   "3",
			Name: "Android",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
