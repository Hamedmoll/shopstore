package mysql

import (
	"shopstoretest/entity"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (mysql MySQLDB) AddBasket(basket entity.BasketItem) (entity.BasketItem, error) {
	const op = "mysql.add basket"
	res, exErr := mysql.DB.Exec("insert into basket_items(user_id, product_id, price, count) values(?, ?, ?, ?)",
		basket.UserID, basket.ProductID, basket.Price, basket.Count)

	if exErr != nil {

		return entity.BasketItem{}, richerror.New(op).WithError(exErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantExecute)
	}

	id, lErr := res.LastInsertId()
	if lErr != nil {

		return entity.BasketItem{}, richerror.New(op).WithError(lErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantCallLastInsertMethod)
	}

	basket.ID = uint(id)

	return basket, nil
}

func (mysql MySQLDB) GetBasketsByID(id uint) ([]entity.BasketItem, error) {
	const op = "mysql.get basket by id"
	rows, qErr := mysql.DB.Query("select id, product_id, user_id, price, count from basket_items where user_id = ?", id)
	if qErr != nil {

		return nil, richerror.New(op).WithError(qErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantQuery)
	}

	var baskets = make([]entity.BasketItem, 0)
	var tmpBasket entity.BasketItem

	for rows.Next() {
		sErr := rows.Scan(&tmpBasket.ID, &tmpBasket.ProductID, &tmpBasket.UserID, &tmpBasket.Price, &tmpBasket.Count)
		if sErr != nil {

			return nil, richerror.New(op).WithError(sErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
		}

		baskets = append(baskets, tmpBasket)
	}

	return baskets, nil
}

func (mysql MySQLDB) RemoveBasketItems(userID uint) error {
	const op = "mysql.remove basket items"
	_, exErr := mysql.DB.Exec("delete from basket_items where user_id = ?", userID)
	if exErr != nil {

		return richerror.New(op).WithError(exErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantExecute)
	}

	return nil
}
