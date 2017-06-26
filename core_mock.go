package geoedge

import (
	"bytes"
	"fmt"
	"reflect"
)

type MockData struct {
	Method     string
	Url        string
	Body       []byte
	Headers    map[string]string
	ResultBody []byte
	ResultErr  error
}

type geoedgeApiCoreMock struct {
	data []MockData
}

func (api *geoedgeApiCoreMock) makeRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	for _, val := range api.data {
		if method == val.Method &&
			url == val.Url &&
			bytes.Compare(body, val.Body) == 0 &&
			reflect.DeepEqual(headers, val.Headers) {

			return val.ResultBody, val.ResultErr
		}
	}

	notFoundBody := []byte(`{
    "status": {
        "code": "UnknownAPIMethod",
        "message": "You need to call the API with a valid method."
    }
}`)

	return notFoundBody, fmt.Errorf("Status code: 404 at %v: %v", url, string(notFoundBody))
}
