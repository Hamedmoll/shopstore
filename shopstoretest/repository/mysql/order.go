package mysql

import (
	"shopstoretest/entity"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (mysql MySQLDB) AddOrder(items []entity.BasketItem) ([]entity.OrderItem, error) {
	const op = "mysql.add order"

	orderedItems := make([]entity.OrderItem, 0)

	for _, item := range items {
		orderItem := entity.OrderItem{
			ID:        0,
			UserID:    item.UserID,
			ProductID: item.ProductID,
			Price:     item.Price,
			Count:     item.Count,
			Total:     item.Count * item.Price,
		}
		res, exErr := mysql.DB.Exec("insert into order_list(user_id, product_id, price, count, total_price) values(?, ?, ?, ?, ?)",
			orderItem.UserID, orderItem.ProductID, orderItem.Price, orderItem.Count, orderItem.Total)

		if exErr != nil {

			return nil, richerror.New(op).WithError(exErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantExecute)
		}

		id, lErr := res.LastInsertId()
		if lErr != nil {

			return nil, richerror.New(op).WithError(lErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantCallLastInsertMethod)
		}

		orderItem.ID = uint(id)

		orderedItems = append(orderedItems, orderItem)
	}

	return orderedItems, nil
}
