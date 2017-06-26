package geoedge

// Request
type listAlertCategoriesRequest struct {
	api *geoedgeApi
}

func (api *geoedgeApi) ListAlertCategoriesRequest() *listAlertCategoriesRequest {
	return &listAlertCategoriesRequest{api: api}
}

// Responce
type listAlertCategoriesResponce struct {
	AlertCategories []alertCategory `json:"categories"`
}

// Do
func (request *listAlertCategoriesRequest) Do() (alertCategories []alertCategory, err error) {
	responce := listAlertCategoriesResponce{}
	err = request.api.makeRequest("List Alert Categories", "", nil, nil, &responce)
	if err != nil {
		return
	}

	alertCategories = responce.AlertCategories
	return
}
