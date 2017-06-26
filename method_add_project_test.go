package geoedge

import (
	"net/url"
	"testing"
)

func TestAddProjectRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "POST",
				Url:    apiUrl + "projects",
				Body:   []byte(`auto_scan=0&emulators=65&locations=GB%2CAU&name=Test+Project+1&scan_type=1&tag=http%3A%2F%2Fwww.bbc.com%2F`),
				Headers: map[string]string{
					"Host":          parsedUrl.Host,
					"Authorization": "testToken",
					"Content-Type":  "application/x-www-form-urlencoded",
				},
				ResultBody: []byte(`{
    "status": {
        "code": "Success",
        "message": "Success"
    },
    "response": {
        "project_id": "4e4c504b6ce6fe44da8610c8364cb5b4"
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.AddProjectRequestRequest(
		"Test Project 1",
		"http://www.bbc.com/",
		0,
		1,
		[]string{"GB", "AU"},
		AddProjectRequestEmulators([]string{"65"}),
	).Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := "4e4c504b6ce6fe44da8610c8364cb5b4"

	if result != expectedResult {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
