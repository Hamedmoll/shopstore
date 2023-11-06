package param

type AddProductRequest struct {
	Name        string `json:"name"`
	Count       uint   `json:"count"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       uint   `json:"price"`
}

type AddProductResponse struct {
	ProductInfo ProductInfo `json:"product_info"`
}

type ProductInfo struct {
	Name        string `json:"name"`
	Count       uint   `json:"count"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       uint   `json:"price"`
}

type ShowByCategoryRequest struct {
	CategoryStr string `json:"category"`
}

type ShowByCategoryResponse struct {
	Products []ProductInfo `json:"products"`
}
