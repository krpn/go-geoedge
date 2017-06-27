package geoedge

import (
	"net/url"
	"testing"
)

func TestWhitelistAlertHistoryRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "POST",
				Url:    apiUrl + "alertshistory/z2COSod4AKP80fmr_zm04w/whitelist",
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

	err := api.WhitelistAlertHistoryRequest("z2COSod4AKP80fmr_zm04w").Do()
	if err != nil {
		t.Error(err)
	}
}
