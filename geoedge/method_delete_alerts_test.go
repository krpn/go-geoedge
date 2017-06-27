package geoedge

import (
	"net/url"
	"testing"
)

func TestDeleteAlertsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "DELETE",
				Url:    apiUrl + "alerts/7d8a14b95a8418f3054492cfc5e2bfad,8e8a14b95a8418f3054492cfc5e2bfad",
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

	err := api.DeleteAlertRequest([]string{"7d8a14b95a8418f3054492cfc5e2bfad", "8e8a14b95a8418f3054492cfc5e2bfad"}).Do()
	if err != nil {
		t.Error(err)
	}
}
