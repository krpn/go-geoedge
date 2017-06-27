package geoedge

import (
	"net/url"
	"testing"
)

func TestAlertProjectsBindRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "POST",
				Url:    apiUrl + "alerts/3d47d72feb37425f989d7f9d4276facc/projects-bind",
				Body:   []byte(`projects=-1`),
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

	err := api.AlertProjectsBindRequest("3d47d72feb37425f989d7f9d4276facc", []string{"-1"}).Do()
	if err != nil {
		t.Error(err)
	}
}
