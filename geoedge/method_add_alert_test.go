package geoedge

import (
	"net/url"
	"testing"
)

func TestAddAlertRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "POST",
				Url:    apiUrl + "alerts",
				Body:   []byte(`locations=NZ%2CAU&name=Alert+1&notification_emails=test@example.com%2Ctest2@example.com&notification_time=2&trigger_ad_auto_play_sound=1&trigger_ad_video_max_length=15&trigger_scan_malware=1`),
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
        "alert_id": "3d47d72feb37425f989d7f9d4276facc"
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.AddAlertRequestRequest(
		"Alert 1",
		AddAlertRequestLocations([]string{"NZ", "AU"}),
		AddAlertRequestNotificationEmails([]string{"test@example.com", "test2@example.com"}),
		AddAlertRequestNotificationTime(2),
		AddAlertRequestTriggerScanMalware(1),
		AddAlertRequestTriggerAdAutoPlaySound(1),
		AddAlertRequestTriggerAdVideoMaxLength(15),
	).Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := "3d47d72feb37425f989d7f9d4276facc"

	if result != expectedResult {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
