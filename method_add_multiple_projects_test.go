package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestAddMultipleProjectsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "POST",
				Url:    apiUrl + "projects/bulk",
				Body:   []byte(`[{"name":"Test Project 1","tag":"http:\/\/www.bbc.com\/a","auto_scan":0,"scan_type":1,"locations":"GB,AU","emulators":"65"},{"name":"Test Project 2","tag":"http:\/\/www.bbc.com\/b","auto_scan":0,"scan_type":1,"locations":"GB,AU","emulators":"68"}]`),
				Headers: map[string]string{
					"Host":          parsedUrl.Host,
					"Authorization": "testToken",
					"Content-Type":  "application/json",
				},
				ResultBody: []byte(`{
    "status": {
        "code": "Success",
        "message": "Success"
    },
    "response": {
        "projects": [
            {
                "status": {
                    "code": "Success",
                    "message": "Success"
                },
                "project_id": "591b644cee4d41ff0ffaab920463d536"
            },
            {
                "status": {
                    "code": "Success",
                    "message": "Success"
                },
                "project_id": "e83785b63a858fcc7d18231964088502"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.AddMultipleProjectsRequest(
		api.AddProjectRequestRequest(
			"Test Project 1",
			"http://www.bbc.com/a",
			0,
			1,
			[]string{"GB", "AU"},
			AddProjectRequestEmulators([]string{"65"}),
		),
		api.AddProjectRequestRequest(
			"Test Project 2",
			"http://www.bbc.com/b",
			0,
			1,
			[]string{"GB", "AU"},
			AddProjectRequestEmulators([]string{"68"}),
		),
	).Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []addMultipleProjectsResult{
		{
			Status: responceStatus{
				Code:    "Success",
				Message: "Success",
			},
			ProjectId: "591b644cee4d41ff0ffaab920463d536",
		},
		{
			Status: responceStatus{
				Code:    "Success",
				Message: "Success",
			},
			ProjectId: "e83785b63a858fcc7d18231964088502",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
