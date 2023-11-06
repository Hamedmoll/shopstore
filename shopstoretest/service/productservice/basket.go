package productservice

import (
	"shopstoretest/entity"
	"shopstoretest/param"
	"shopstoretest/pkg/organize"
	"shopstoretest/pkg/richerror"
)

func (s Service) AddBasketItem(req param.AddToBasketRequest) (param.AddToBasketResponse, error) {
	const op = "product service.add basket item"

	user, guErr := s.Repository.GetUserByID(req.UserID)
	if guErr != nil {

		return param.AddToBasketResponse{}, richerror.New(op).WithError(guErr)
	}

	product, gpErr := s.Repository.GetProductByID(req.ProductID)
	if gpErr != nil {

		return param.AddToBasketResponse{}, richerror.New(op).WithError(gpErr)
	}

	//fmt.Printf("\n\n %v \n\n", req.Count)

	newBasket := entity.BasketItem{
		ID:        0,
		UserID:    user.ID,
		ProductID: product.ID,
		Price:     product.Price,
		Count:     req.Count,
	}

	createdBasket, aErr := s.Repository.AddBasket(newBasket)

	if aErr != nil {

		return param.AddToBasketResponse{}, richerror.New(op).WithError(aErr)
	}

	res := param.AddToBasketResponse{Basket: createdBasket}

	return res, nil
}

func (s Service) ShowBasket(req param.ShowBasketRequest) (param.ShowBasketResponse, error) {
	const op = "product service.show basket"

	basket, gErr := s.Repository.GetBasketsByID(req.ID)
	if gErr != nil {

		return param.ShowBasketResponse{}, richerror.New(op).WithError(gErr)
	}

	//fmt.Println("hi there")
	//fmt.Printf("\n\n\n %+v \n\n\n", basket)

	organizeItems, _ := organize.OrganizeBasketItem(basket)

	organizeItemList := make([]entity.BasketItem, 0)

	var total uint

	for _, item := range organizeItems {
		organizeItemList = append(organizeItemList, item)
		total += item.Price * item.Count
	}

	res := param.ShowBasketResponse{
		Baskets:    organizeItemList,
		TotalPrice: total,
	}

	return res, nil
}

func (s Service) RemoveBasketItems(req param.RemoveBasketRequest) (param.RemoveBasketResponse, error) {
	const op = "product service.remove basket items"

	rErr := s.Repository.RemoveBasketItems(req.UserID)
	if rErr != nil {

		return param.RemoveBasketResponse{}, richerror.New(op).WithError(rErr)
	}

	res := param.RemoveBasketResponse{}

	return res, nil
}
