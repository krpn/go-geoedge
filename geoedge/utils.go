package geoedge

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func structToUrlencodedFormat(structObj interface{}) ([]byte, error) {
	structBytes, err := json.Marshal(structObj)
	if err != nil {
		return nil, err
	}

	tempMap := make(map[string]interface{})
	err = json.Unmarshal(structBytes, &tempMap)
	if err != nil {
		return nil, err
	}

	if len(tempMap) == 0 {
		return nil, nil
	}

	// Sorting map (need for tests)
	var keys []string
	for key := range tempMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if len(keys) == 0 {
		return nil, nil
	}

	var result string
	for _, key := range keys {
		val := url.QueryEscape(fmt.Sprint(tempMap[key]))
		val = strings.Replace(val, "%40", "@", -1) // fix emails

		result += fmt.Sprintf("%v=%v&", key, val)
	}
	result = result[:len(result)-1]

	return []byte(result), nil
}
