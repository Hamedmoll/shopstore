package productservice

import (
	"shopstoretest/cfg"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (s Service) Order(req param.AddOrderRequest) (param.AddOrderResponse, error) {
	const op = "product service.order"

	showReq := param.ShowBasketRequest{ID: req.UserID}
	res, sErr := s.ShowBasket(showReq)
	if sErr != nil {

		return param.AddOrderResponse{}, richerror.New(op).WithError(sErr)
	}

	items := res.Baskets
	cost := int(res.TotalPrice)

	user, gErr := s.Repository.GetUserByID(req.UserID)
	if gErr != nil {

		return param.AddOrderResponse{}, richerror.New(op).WithError(gErr)
	}

	if user.Credit-cost >= cfg.MinimumCredit {
		user.Credit -= cost
		_, uErr := s.Repository.UpdateUserWithCredit(user.Credit-cost, req.UserID)
		if uErr != nil {

			return param.AddOrderResponse{}, richerror.New(op).WithError(uErr)
		}

		orderedItems, oErr := s.Repository.AddOrder(items)
		if oErr != nil {

			return param.AddOrderResponse{}, richerror.New(op).WithError(oErr)
		}

		response := param.AddOrderResponse{}
		response.Items = orderedItems

		remReq := param.RemoveBasketRequest{UserID: req.UserID}
		_, rErr := s.RemoveBasketItems(remReq)
		if rErr != nil {

			return param.AddOrderResponse{}, richerror.New(op).WithError(rErr)
		}

		return response, nil
	} else {

		return param.AddOrderResponse{}, richerror.New(op).WithMessage(errormsg.DontHaveCredit).WithKind(richerror.KindDontHaveCredit)
	}
}
