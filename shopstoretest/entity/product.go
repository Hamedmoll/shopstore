package entity

type Product struct {
	ID          uint   `json:"id"`
	Price       uint   `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	Count       uint   `json:"count"`
}
