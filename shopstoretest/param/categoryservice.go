package param

type AddCategoryRequest struct {
	Name string `json:"name"`
}

type AddCategoryResponse struct {
	CategoryInfo CategoryInfo `json:"category_info"`
}

type CategoryInfo struct {
	Name string `json:"name"`
}
