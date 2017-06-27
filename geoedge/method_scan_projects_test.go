package geoedge

import (
	"net/url"
	"reflect"
	"testing"
)

func TestScanProjectsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "POST",
				Url:    apiUrl + "projects/2460244ac44e78c09ece2d24519cf71e,3460244ac44e78c09ece2d24519cf71e/scan",
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
        "projects": {
            "2460244ac44e78c09ece2d24519cf71e": {
                "scans": [
                    "1a755a2adf6301380b5ed35fb303767c",
                    "301380b5ed35fb303767c5t55t5t5677"
                ]
            }
        }
    }
}`),
				ResultErr: nil,
			},
		},
	)

	result, err := api.ScanProjectRequest([]string{"2460244ac44e78c09ece2d24519cf71e", "3460244ac44e78c09ece2d24519cf71e"}).Do()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []scanProjectsResponceResult{
		{
			ProjectId: "2460244ac44e78c09ece2d24519cf71e",
			Scans:     []string{"1a755a2adf6301380b5ed35fb303767c", "301380b5ed35fb303767c5t55t5t5677"},
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
}
