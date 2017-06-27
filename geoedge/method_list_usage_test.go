package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListUsageRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "usage",
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
        "usage": [
            {
                "daily_scans": "15",
                "daily_scans_automatic": "10",
                "daily_scans_manual": "5",
                "monthly_scans": "591",
                "monthly_scans_automatic": "500",
                "monthly_scans_manual": "91",
                "billing_monthly_scans": "167",
                "billing_monthly_scans_automatic": "135",
                "billing_monthly_scans_manual": "32"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListUsageRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []usage{
		{
			DailyScans:                   "15",
			DailyScansAutomatic:          "10",
			DailyScansManual:             "5",
			MonthlyScans:                 "591",
			MonthlyScansAutomatic:        "500",
			MonthlyScansManual:           "91",
			BillingMonthlyScans:          "167",
			BillingMonthlyScansAutomatic: "135",
			BillingMonthlyScansManual:    "32",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
