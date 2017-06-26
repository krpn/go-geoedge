package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListAlertCategoriesRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "alerts/categories",
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
        "categories": [
            {
                "id": "1",
                "name": "News"
            },
            {
                "id": "2",
                "name": "Offensive"
            },
            {
                "id": "3",
                "name": "Games"
            },
            {
                "id": "4",
                "name": "Religion"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ListAlertCategoriesRequest().Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []alertCategory{
		{
			Id:   "1",
			Name: "News",
		},
		{
			Id:   "2",
			Name: "Offensive",
		},
		{
			Id:   "3",
			Name: "Games",
		},
		{
			Id:   "4",
			Name: "Religion",
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
