package mysql

import (
	"shopstoretest/entity"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (mysql MySQLDB) GetUserPermissionTitles(userID uint) ([]entity.PermissionTitle, error) {
	const op = "mysql.get user permission title"
	rows, qErr := mysql.DB.Query("select permission_id from access_controls where actor_id = ?", userID)
	if qErr != nil {

		return nil, richerror.New(op).WithError(qErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantQuery)
	}

	permissionTitleIDs := make([]uint, 0)
	var tmpPermissionTitleID uint

	for rows.Next() {
		sErr := rows.Scan(&tmpPermissionTitleID)
		if sErr != nil {

			return nil, richerror.New(op).WithError(qErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
		}

		permissionTitleIDs = append(permissionTitleIDs, tmpPermissionTitleID)
	}

	permissionTitles := make([]entity.PermissionTitle, 0)
	var tmpPermissionTitle entity.PermissionTitle

	for _, id := range permissionTitleIDs {
		row := mysql.DB.QueryRow("select title from permissions where id = ?", id)
		sErr := row.Scan(&tmpPermissionTitle)
		if sErr != nil {

			return nil, richerror.New(op).WithError(qErr).WithKind(richerror.KindUnexpected).WithMessage(errormsg.CantScan)
		}

		permissionTitles = append(permissionTitles, tmpPermissionTitle)
	}

	return permissionTitles, nil
}
