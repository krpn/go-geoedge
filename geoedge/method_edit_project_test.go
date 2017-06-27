package geoedge

import (
	"net/url"
	"testing"
)

func TestEditProjectRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "PUT",
				Url:    apiUrl + "projects/1a755a2adf6301380b5ed35fb303767c",
				Body:   []byte(`emulators=67&locations=AU%2CNZ%2CUS&name=Test+Project+2&tag=http%3A%2F%2Fwww.yahoo.com%2F`),
				Headers: map[string]string{
					"Host":          parsedUrl.Host,
					"Authorization": "testToken",
					"Content-Type":  "application/x-www-form-urlencoded",
				},
				ResultBody: []byte(`{
    "status": {
        "code": "Success",
        "message": "Success"
    }
}`),
				ResultErr: nil,
			},
		},
	)

	err := api.EditProjectRequest(
		"1a755a2adf6301380b5ed35fb303767c",
		EditProjectRequestName("Test Project 2"),
		EditProjectRequestTag("http://www.yahoo.com/"),
		EditProjectRequestLocations([]string{"AU", "NZ", "US"}),
		EditProjectRequestEmulators([]string{"67"}),
	).Do()
	if err != nil {
		t.Error(err)
	}
}
