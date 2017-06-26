package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListNetworksRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "networks",
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
        "networks": [
            {
                "id": "913",
                "name": "152 media"
            },
            {
                "id": "87",
                "name": "24 Interactive"
            },
            {
                "id": "59",
                "name": "24\/7 Media"
            },
            {
                "id": "452",
                "name": "33Across"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListNetworksRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []network{
		{
			Id:   "913",
			Name: "152 media",
		},
		{
			Id:   "87",
			Name: "24 Interactive",
		},
		{
			Id:   "59",
			Name: "24/7 Media",
		},
		{
			Id:   "452",
			Name: "33Across",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
