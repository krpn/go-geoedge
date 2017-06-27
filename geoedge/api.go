package geoedge

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type api struct {
	url            string
	regularHeaders map[string]string
	core           apiCoreInterface
}

func getApi(authorizationToken string, core apiCoreInterface) *api {
	parsedUrl, _ := url.Parse(apiUrl)
	return &api{
		url: apiUrl,
		regularHeaders: map[string]string{
			"Host":          parsedUrl.Host,
			"Authorization": authorizationToken,
		},
		core: core,
	}
}

func NewApi(authorizationToken string) *api {
	return getApi(authorizationToken, &apiCore{})
}

func NewApiMocked(authorizationToken string, mockData []MockData) *api {
	return getApi(authorizationToken, &apiCoreMock{data: mockData})
}

func (api *api) makeRequest(
	method string, urlParams string,
	body []byte,
	extraHeaders map[string]string,
	responceStruct interface{}) (err error) {

	methodConfig, ok := apiMethods[method]
	if !ok {
		return fmt.Errorf("Unknown GeoEdge Ad Verification API Method: %v.", method)
	}

	fullUrl := api.url + methodConfig.path

	if urlParams != "" || strings.Contains(fullUrl, "%v") {
		fullUrl = fmt.Sprintf(fullUrl, urlParams)
	}

	if fullUrl[len(fullUrl)-1:] == "?" {
		fullUrl = fullUrl[:len(fullUrl)-1]
	}

	var headers map[string]string
	if extraHeaders != nil {
		headers = make(map[string]string)
		for key, val := range api.regularHeaders {
			headers[key] = val
		}
		for key, val := range extraHeaders {
			headers[key] = val
		}
	} else {
		headers = api.regularHeaders
	}

	respBody, err := api.core.makeRequest(methodConfig.httpMethod, fullUrl, body, headers)
	if err != nil {
		return
	}

	respRaw := new(rawResponse)
	err = json.Unmarshal(respBody, respRaw)
	if err != nil {
		return
	}

	if respRaw.Status.Code != requestStatusSuccess {
		return fmt.Errorf("GeoEdge Ad Verification API unsuccessfull result: %v, %v",
			respRaw.Status.Code,
			respRaw.Status.Message)
	}

	if responceStruct != nil {
		err = json.Unmarshal(fixKeys(respRaw.Response), responceStruct)
		if err != nil {
			return
		}
	}

	return
}

func fixKeys(source []byte) []byte {
	result := string(source)

	replacements := map[string]string{
		`"sequence": []`:       `"sequence": {}`,
		`"locations": []`:      `"locations": {}`,
		`"emulation_sets": []`: `"emulation_sets": {}`,
		`"emulators": []`:      `"emulators": {}`,
		`"tree": []`:           `"tree": {}`,
		`"network": ""`:        `"network": {}`,
	}

	for oldStr, newStr := range replacements {
		result = strings.Replace(result, oldStr, newStr, -1)
	}

	return []byte(result)
}
