package geoedge

import (
	"net/url"
	"testing"
)

func TestEditAlertRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "PUT",
				Url:    apiUrl + "alerts/1297cef3cdf39c5525b457d34c6b05ae",
				Body:   []byte(`locations=CA%2CY0&name=Alert+1+-+altered&notification_emails=test3%40example.com&notification_time=0&trigger_ad_auto_play_sound=1&trigger_ad_video_max_length=30&trigger_scan_malware=0`),
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

	err := api.EditAlertRequestRequest(
		"1297cef3cdf39c5525b457d34c6b05ae",
		EditAlertRequestName("Alert 1 - altered"),
		EditAlertRequestLocations([]string{"CA", "Y0"}),
		EditAlertRequestNotificationEmails([]string{"test3@example.com"}),
		EditAlertRequestNotificationTime(0),
		EditAlertRequestTriggerScanMalware(0),
		EditAlertRequestTriggerAdAutoPlaySound(1),
		EditAlertRequestTriggerAdVideoMaxLength(30),
	).Do()
	if err != nil {
		t.Error(err)
	}
}
