package geoedge

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type geoedgeApi struct {
	url            string
	regularHeaders map[string]string
	core           geoedgeApiCoreInterface
}

func getApi(authorizationToken string, core geoedgeApiCoreInterface) *geoedgeApi {
	parsedUrl, _ := url.Parse(apiUrl)
	return &geoedgeApi{
		url: apiUrl,
		regularHeaders: map[string]string{
			"Host":          parsedUrl.Host,
			"Authorization": authorizationToken,
		},
		core: core,
	}
}

func NewApi(authorizationToken string) *geoedgeApi {
	return getApi(authorizationToken, &geoedgeApiCore{})
}

func NewApiMocked(authorizationToken string, mockData []MockData) *geoedgeApi {
	return getApi(authorizationToken, &geoedgeApiCoreMock{data: mockData})
}

func (api *geoedgeApi) makeRequest(
	method string, urlParams string,
	body []byte,
	extraHeaders map[string]string,
	responceStruct interface{}) (err error) {

	methodConfig, ok := apiMethods[method]
	if !ok {
		return fmt.Errorf("Unknown Geoedge API Method: %v.", method)
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
		return fmt.Errorf("Geoedge unsuccessfull result: %v, %v",
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
