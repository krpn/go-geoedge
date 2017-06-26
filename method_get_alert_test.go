package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetAlertRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "alerts/7d8a14b95a8418f3054492cfc5e2bfad",
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
        "alerts": [
            {
                "id": "7d8a14b95a8418f3054492cfc5e2bfad",
                "name": "Alert 1",
                "enabled": "1",
                "global_alert": "1",
                "notification_time": "2",
                "emails": [
                    "test@example.com",
                    "test2@example.com"
                ],
                "locations": {
                    "AU": "Australia",
                    "NZ": "New Zealand",
                    "US": "United States"
                },
                "triggers": [
                    [
                        {
                            "trigger_type_id": "15",
                            "params": [
                                ""
                            ]
                        },
                        {
                            "trigger_type_id": "12",
                            "params": [
                                "1B"
                            ]
                        },
                        {
                            "trigger_type_id": "4",
                            "params": [
                                "Computers & Electronics",
                                "Finance & Insurance",
                                "Games",
                                "Health",
                                "Shopping",
                                "Travel"
                            ]
                        }
                    ]
                ],
                "projects": [
                    {
                        "60336dd34e59a0568cea9b8e0847f7cd": "Project Name 1"
                    }
                ]
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.GetAlertRequest("7d8a14b95a8418f3054492cfc5e2bfad").Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []alert{
		{
			Id:               "7d8a14b95a8418f3054492cfc5e2bfad",
			Name:             "Alert 1",
			Enabled:          "1",
			GlobalAlert:      "1",
			NotificationTime: "2",
			Emails:           []string{"test@example.com", "test2@example.com"},
			Locations: map[string]string{
				"AU": "Australia",
				"NZ": "New Zealand",
				"US": "United States",
			},
			Triggers: [][]struct {
				TriggerTypeId string   `json:"trigger_type_id"`
				Params        []string `json:"params"`
			}{
				{
					{
						TriggerTypeId: "15",
						Params:        []string{""},
					},
					{
						TriggerTypeId: "12",
						Params:        []string{"1B"},
					},
					{
						TriggerTypeId: "4",
						Params: []string{
							"Computers & Electronics",
							"Finance & Insurance",
							"Games",
							"Health",
							"Shopping",
							"Travel",
						},
					},
				},
			},
			Projects: []map[string]string{
				{
					"60336dd34e59a0568cea9b8e0847f7cd": "Project Name 1",
				},
			},
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
