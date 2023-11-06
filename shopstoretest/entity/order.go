package entity

type OrderItem struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Price     uint `json:"price"`
	Count     uint `json:"count"`
	Total     uint `json:"total_price"`
}
