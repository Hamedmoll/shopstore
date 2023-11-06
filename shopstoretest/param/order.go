package param

import "shopstoretest/entity"

type AddOrderRequest struct {
	UserID uint `json:"id"`
}

type AddOrderResponse struct {
	Items []entity.OrderItem `json:"items"`
}
