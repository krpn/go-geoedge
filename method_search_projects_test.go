package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestSearchProjectsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "projects?tag=http%3A%2F%2Fwww.yahoo.com%2F",
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
        "projects": [
            {
                "id": "1a755a2adf6301380b5ed35fb303767c",
                "name": "Test Project 2",
                "auto_scan": 1,
                "creation_time": 1431984877
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.SearchProjectsRequest(
		SearchProjectsRequestTag("http://www.yahoo.com/"),
	).Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []projectBase{
		{
			Id:           "1a755a2adf6301380b5ed35fb303767c",
			Name:         "Test Project 2",
			AutoScan:     1,
			CreationTime: 1431984877,
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
