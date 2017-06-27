package geoedge

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type apiCoreInterface interface {
	makeRequest(method, url string, body []byte, headers map[string]string) ([]byte, error)
}

type apiCore struct{}

func (api *apiCore) makeRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Close = true

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status code: %v at %v: %v", resp.StatusCode, url, string(respBodyBytes))
	}

	return respBodyBytes, nil
}

type MockData struct {
	Method     string
	Url        string
	Body       []byte
	Headers    map[string]string
	ResultBody []byte
	ResultErr  error
}

type apiCoreMock struct {
	data []MockData
}

func (api *apiCoreMock) makeRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
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
