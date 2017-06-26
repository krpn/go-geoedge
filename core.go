package geoedge

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type geoedgeApiCoreInterface interface {
	makeRequest(method, url string, body []byte, headers map[string]string) ([]byte, error)
}

type geoedgeApiCore struct{}

func (api *geoedgeApiCore) makeRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
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
