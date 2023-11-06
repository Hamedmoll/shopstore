package mysql

import (
	"shopstoretest/entity"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (mysql MySQLDB) AddProduct(product param.AddProductRequest) (entity.Product, error) {
	const op = "mysql.add product"

	existCategory, eErr := mysql.CheckExistCategory(product.Category)
	if eErr != nil {

		return entity.Product{}, richerror.New(op).WithError(eErr)
	}

	if !existCategory {
		newCategory := entity.Category{Name: product.Category}
		_, aErr := mysql.AddCategory(newCategory)
		if aErr != nil {

			return entity.Product{}, richerror.New(op).WithError(aErr)
		}
	}

	category, gErr := mysql.GetCategoryByName(product.Category)
	if gErr != nil {

		return entity.Product{}, richerror.New(op).WithError(gErr)
	}

	res, exErr := mysql.DB.Exec("insert into products(name, price, description, category_id, count) values(?, ?, ?, ?, ?)",
		product.Name, product.Price, product.Description, category.ID, product.Count)
	if exErr != nil {

		return entity.Product{}, richerror.New(op).WithError(exErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantExecute)
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.Product{}, richerror.New(op).WithError(exErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantCallLastInsertMethod)
	}

	createdProduct := entity.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  category.ID,
		Count:       product.Count,
	}

	createdProduct.ID = uint(id)

	return createdProduct, nil
}

func (mysql MySQLDB) GetProductByCategory(name string) ([]param.ProductInfo, error) {
	const op = "mysql.get product by category"

	category, gErr := mysql.GetCategoryByName(name)
	if gErr != nil {

		return nil, richerror.New(op).WithError(gErr)

	}

	rows, qErr := mysql.DB.Query("select name, count, price, Description from products where category_id = ?", category.ID)
	if qErr != nil {

		return nil, richerror.New(op).WithError(qErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantQuery)
	}

	var tmpProduct param.ProductInfo
	var products []param.ProductInfo

	for rows.Next() {
		sErr := rows.Scan(&tmpProduct.Name, &tmpProduct.Count, &tmpProduct.Price, &tmpProduct.Description)
		if sErr != nil {

			return nil, richerror.New(op).WithError(sErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
		}

		tmpProduct.Category = name
		products = append(products, tmpProduct)
	}

	return products, nil
}

func (mysql MySQLDB) GetProductByID(id uint) (entity.Product, error) {
	const op = "mysql.get product by id"

	row := mysql.DB.QueryRow("select * from products where id = ?", id)
	product := entity.Product{}
	var createdAt []uint8

	err := row.Scan(&product.ID, &product.Name, &product.Count, &product.Price,
		&product.Description, &product.CategoryID, &createdAt)
	if err != nil {

		return entity.Product{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
	}

	return product, nil
}
