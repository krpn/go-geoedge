package geoedge

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestSearchAdsRequest(t *testing.T) {
	parsedUrl, _ := url.Parse(apiUrl)
	api := NewApiMocked(
		"testToken",
		[]MockData{
			{
				Method: "GET",
				Url:    apiUrl + "ads?limit=50&max_datetime=2015-07-10+00%3A00%3A00&min_datetime=2015-07-09+00%3A00%3A00&offset=100",
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
        "prev_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `ads?min_datetime=2015-07-09+00%3A00%3A00&max_datetime=2015-07-10+00%3A00%3A00&offset=50&limit=50",
        "next_page": "` + strings.Replace(apiUrl, `/`, `\/`, -1) + `ads?min_datetime=2015-07-09+00%3A00%3A00&max_datetime=2015-07-10+00%3A00%3A00&offset=150&limit=50",
        "ads": [
            {
                "id": "ybLbcC_wK3o8BEjHIkOcBw"
            },
            {
                "id": "de3hGjzHFa64_LIxI3AV8A"
            },
            {
                "id": "7iPbOgeHNBkeHSHaUy7RfA"
            },
            {
                "id": "EwvVWRzpuIS3fa6IWZam9A"
            },
            {
                "id": "_sFbmNDG1_yrAz_3Yrgkxw"
            },
            {
                "id": "ueKT--lfBtFfT8Ls6qA7Xw"
            }
        ]
    }
}`),
				ResultErr: nil,
			},
		},
	)

	minDatetime, _ := time.Parse("2006-01-02 15:04:05", "2015-07-09 00:00:00")
	maxDatetime, _ := time.Parse("2006-01-02 15:04:05", "2015-07-10 00:00:00")

	result, prevPage, nextPage, err := api.SearchAdsRequest(
		minDatetime,
		SearchAdsRequestMaxDatetime(maxDatetime),
		SearchAdsRequestOffset(100),
		SearchAdsRequestLimit(50),
	).DoExtra()
	if err != nil {
		t.Error(err)
	}

	expectedResult := []string{
		"ybLbcC_wK3o8BEjHIkOcBw",
		"de3hGjzHFa64_LIxI3AV8A",
		"7iPbOgeHNBkeHSHaUy7RfA",
		"EwvVWRzpuIS3fa6IWZam9A",
		"_sFbmNDG1_yrAz_3Yrgkxw",
		"ueKT--lfBtFfT8Ls6qA7Xw",
	}
	expectedPrevPage := apiUrl + "ads?min_datetime=2015-07-09+00%3A00%3A00&max_datetime=2015-07-10+00%3A00%3A00&offset=50&limit=50"
	expectedNextPage := apiUrl + "ads?min_datetime=2015-07-09+00%3A00%3A00&max_datetime=2015-07-10+00%3A00%3A00&offset=150&limit=50"

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
