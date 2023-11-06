package mysql

import (
	"database/sql"
	"shopstoretest/entity"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (mysql MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.is phone number unique"

	row := mysql.DB.QueryRow("select * from users where phone_number = ?", phoneNumber)

	user := entity.User{}
	var createdAt []uint8
	var role string

	uErr := row.Scan(&user.ID, &user.Name, &user.Password, &user.Credit, &user.PhoneNumber, &createdAt, &role)
	user.Role = entity.MapToRoleEntity(role)

	if uErr != nil {
		if uErr == sql.ErrNoRows {

			return true, nil
		}

		return false, richerror.New(op).WithError(uErr).WithMessage(errormsg.CantScan).WithKind(richerror.KindUnexpected)
	}

	return false, nil
}

func (mysql MySQLDB) Register(user entity.User) (entity.User, error) {
	const op = "mysql.register"

	res, exErr := mysql.DB.Exec("insert into users(name, hashed_password, phone_number) values(?, ?, ?)",
		user.Name, user.Password, user.PhoneNumber)
	if exErr != nil {

		return entity.User{}, richerror.New(op).WithError(exErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantExecute)
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.User{}, richerror.New(op).WithError(iErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantCallLastInsertMethod)
	}

	user.ID = uint(id)

	return user, nil
}

func (mysql MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.get user by phone number"
	row := mysql.DB.QueryRow("select * from users where phone_number = ?", phoneNumber)
	user := entity.User{}
	var createdAt []uint8
	var roleString string

	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Credit, &user.PhoneNumber, &createdAt, &roleString)
	if err != nil {

		return entity.User{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
	}

	return entity.User{
		ID:          user.ID,
		Role:        entity.MapToRoleEntity(roleString),
		Name:        user.Name,
		Password:    user.Password,
		Credit:      user.Credit,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (mysql MySQLDB) GetUserByID(ID uint) (entity.User, error) {
	const op = "mysql.get user by id"

	row := mysql.DB.QueryRow("select * from users where `id` = ?", ID)
	user := entity.User{}
	var createdAt []uint8
	var roleString string

	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Credit, &user.PhoneNumber, &createdAt, &roleString)
	if err != nil {

		return entity.User{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
	}

	return entity.User{
		ID:          user.ID,
		Role:        entity.MapToRoleEntity(roleString),
		Name:        user.Name,
		Password:    user.Password,
		Credit:      user.Credit,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (mysql MySQLDB) UpdateUserWithCredit(credit int, id uint) (entity.User, error) {
	const op = "mysql.update user with credit"

	_ = mysql.DB.QueryRow("update users set credit = ? where id = ?", credit, id)

	modifyUser, gErr := mysql.GetUserByID(id)
	if gErr != nil {

		return entity.User{}, richerror.New(op).WithError(gErr)
	}

	return modifyUser, nil
}
