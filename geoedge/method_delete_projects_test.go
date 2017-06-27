package geoedge

import (
	"net/url"
	"testing"
)

func TestDeleteProjectsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "DELETE",
				Url:    apiUrl + "projects/2460244ac44e78c09ece2d24519cf71e,3460244ac44e78c09ece2d24519cf71e",
				Body:   nil,
				Headers: map[string]string{
					"Host":          parsedUrl.Host,
					"Authorization": "testToken",
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

	err := api.DeleteProjectRequest([]string{"2460244ac44e78c09ece2d24519cf71e", "3460244ac44e78c09ece2d24519cf71e"}).Do()
	if err != nil {
		t.Error(err)
	}
}
