package geoedge

const (
	TestedApiVersion     = "3.4.25"
	apiUrl               = "https://api.geoedge.com/rest/analytics/v3/"
	requestStatusSuccess = "Success"
)

type apiMethod struct {
	httpMethod string
	path       string
}

var (
	apiMethods = map[string]apiMethod{
		// Projects
		"List Projects":         {httpMethod: "GET", path: "projects?%v"},
		"Get Project":           {httpMethod: "GET", path: "projects/%v"},
		"Add Project":           {httpMethod: "POST", path: "projects"},
		"Add Multiple Projects": {httpMethod: "POST", path: "projects/bulk"},
		"Edit Project":          {httpMethod: "PUT", path: "projects/%v"},
		"Delete Projects":       {httpMethod: "DELETE", path: "projects/%v"},
		"Search Projects":       {httpMethod: "GET", path: "projects?%v"},
		"Scan Projects":         {httpMethod: "POST", path: "projects/%v/scan"},
		"Scan Status":           {httpMethod: "GET", path: "projects/%v/scan-status"},

		// Alerts
		"List Alerts":           {httpMethod: "GET", path: "alerts"},
		"List Alert Categories": {httpMethod: "GET", path: "alerts/categories"},
		"Add Alert":             {httpMethod: "POST", path: "alerts"},
		"Edit Alert":            {httpMethod: "PUT", path: "alerts/%v"},
		"Delete Alerts":         {httpMethod: "DELETE", path: "alerts/%v"},
		"Alert Projects Bind":   {httpMethod: "POST", path: "alerts/%v/projects-bind"},
		"Trigger Types":         {httpMethod: "GET", path: "alerts/trigger-types"},
		"Get Alert":             {httpMethod: "GET", path: "alerts/%v"},

		// Alerts History
		"Get Alert History":       {httpMethod: "GET", path: "alertshistory/%v"},
		"Search Alerts History":   {httpMethod: "GET", path: "alerts/history?%v"},
		"Whitelist Alert History": {httpMethod: "POST", path: "alertshistory/%v/whitelist"},

		// Ads
		"Get Ad":     {httpMethod: "GET", path: "ads/%v"},
		"Search Ads": {httpMethod: "GET", path: "ads?%v"},

		// Landing Pages
		"Search LPs": {httpMethod: "GET", path: "lps?%v"},

		// Locations
		"List Locations": {httpMethod: "GET", path: "locations"},

		// Platforms
		"List Platforms": {httpMethod: "GET", path: "platforms"},

		// Emulation Sets
		"List Emulation Sets": {httpMethod: "GET", path: "emulation-sets"},

		// Emulators
		"List Emulators": {httpMethod: "GET", path: "emulators"},

		// Networks
		"List Networks": {httpMethod: "GET", path: "networks"},

		// Usage
		"List Usage": {httpMethod: "GET", path: "usage"},
	}
)
