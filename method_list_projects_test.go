package geoedge

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestListProjectsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "projects?limit=50&offset=100",
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
        "prev_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `result?offset=50&limit=50",
        "next_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `result?offset=150&limit=50",
        "projects_total": 20,
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

	result, prevPage, nextPage, projectsTotal, err := api.ListProjectsRequest(
		ListProjectsRequestOffset(100),
		ListProjectsRequestLimit(50),
	).DoExtra()
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
	expectedPrevPage := apiUrl + "result?offset=50&limit=50"
	expectedNextPage := apiUrl + "result?offset=150&limit=50"
	expectedProjectsTotal := 20

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
	if prevPage != expectedPrevPage {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedPrevPage, prevPage)
	}
	if nextPage != expectedNextPage {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedNextPage, nextPage)
	}
	if projectsTotal != expectedProjectsTotal {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedProjectsTotal, projectsTotal)
	}
}
