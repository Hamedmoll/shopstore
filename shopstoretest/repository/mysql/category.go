package mysql

import (
	"database/sql"
	"fmt"
	"shopstoretest/entity"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (mysql MySQLDB) CheckExistCategory(name string) (bool, error) {
	const op = "mysql.check exist category"

	row := mysql.DB.QueryRow("select * from categories where name = ?", name)
	category := entity.Category{}
	var createdAt []uint8

	err := row.Scan(&category.ID, &category.Name, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, nil
		}

		return false, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
	}

	return true, nil
}

//i assume before call this function check the existence by CheckExist method

func (mysql MySQLDB) GetCategoryByName(name string) (entity.Category, error) {
	row := mysql.DB.QueryRow("select * from categories where name = ?", name)
	category := entity.Category{}
	var createdAt []uint8

	err := row.Scan(&category.ID, &category.Name, &createdAt)
	if err != nil {

		return entity.Category{}, fmt.Errorf("cant scan %w", err)
	}

	return category, nil
}

func (mysql MySQLDB) AddCategory(category entity.Category) (entity.Category, error) {
	res, exErr := mysql.DB.Exec("insert into categories(name) values(?)", category.Name)
	if exErr != nil {

		return entity.Category{}, exErr
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.Category{}, fmt.Errorf("cant get last id after save new category in table %w", iErr)
	}

	category.ID = uint(id)

	fmt.Println("\n its here \n")
	return category, nil
}
