package geoedge

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestSearchLpsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "lps?limit=50&max_datetime=2015-07-12+00%3A00%3A00&min_datetime=2015-07-08+00%3A00%3A00&offset=100",
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
        "prev_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `lps?min_datetime=2015-07-08+00%3A00%3A00&max_datetime=2015-07-12+00%3A00%3A00&offset=50&limit=50",
        "next_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `lps?min_datetime=2015-07-08+00%3A00%3A00&max_datetime=2015-07-12+00%3A00%3A00&offset=150&limit=50",
        "ads": [
            {
                "id": "7iPbOgeHNBkeHSHaUy7RfA"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	minDatetime, _ := time.Parse("2006-01-02 15:04:05", "2015-07-08 00:00:00")
	maxDatetime, _ := time.Parse("2006-01-02 15:04:05", "2015-07-12 00:00:00")

	result, prevPage, nextPage, err := api.SearchLpsRequest(
		minDatetime,
		SearchLpsRequestMaxDatetime(maxDatetime),
		SearchLpsRequestOffset(100),
		SearchLpsRequestLimit(50),
	).DoExtra()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []string{
		"7iPbOgeHNBkeHSHaUy7RfA",
	}
	expectedPrevPage := apiUrl + "lps?min_datetime=2015-07-08+00%3A00%3A00&max_datetime=2015-07-12+00%3A00%3A00&offset=50&limit=50"
	expectedNextPage := apiUrl + "lps?min_datetime=2015-07-08+00%3A00%3A00&max_datetime=2015-07-12+00%3A00%3A00&offset=150&limit=50"

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedResult, result)
	}
	if prevPage != expectedPrevPage {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedPrevPage, prevPage)
	}
	if nextPage != expectedNextPage {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedNextPage, nextPage)
	}
}
