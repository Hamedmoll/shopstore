package param

import "shopstoretest/entity"

type AddToBasketRequest struct {
	ProductID uint `json:"product_id"`
	UserID    uint `json:"user_id"`
	Count     uint `json:"count"`
}

type AddToBasketResponse struct {
	Basket entity.BasketItem `json:"basket"`
}

type ShowBasketRequest struct {
	ID uint `json:"id"`
}

type ShowBasketResponse struct {
	Baskets    []entity.BasketItem `json:"basket"`
	TotalPrice uint                `json:"total_price"`
}

type RemoveBasketRequest struct {
	UserID uint `json:"user_id"`
}

type RemoveBasketResponse struct {
}
