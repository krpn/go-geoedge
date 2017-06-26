# go-geoedge - A Go library for GeoEdge Ad Verification API

[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/krpn/go-geoedge/blob/master/LICENSE)

This library provides simple interface for all [GeoEdge Ad Verification API methods](https://www.geoedge.com/geoedge_api_examples). 

Tested on API version: **3.4.25**.


## Installation

```
go get github.com/krpn/go-geoedge
```

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/krpn/go-geoedge"
)

const myAuthorizationToken = "myAwesomeToken"

func main() {
	api := geoedge.NewApi(myAuthorizationToken)

	projects, err := api.ListProjectsRequest(geoedge.ListProjectsRequestLimit(3)).Do()
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		projectDetails, err := api.GetProjectRequest(project.Id).Do()
		if err != nil {
			panic(err)
		}

		fmt.Println(projectDetails)
	}
}

```