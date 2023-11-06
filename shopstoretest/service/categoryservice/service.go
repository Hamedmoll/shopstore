package categoryservice

import (
	"shopstoretest/cfg"
	"shopstoretest/entity"
	"shopstoretest/repository/mysql"
	"shopstoretest/service/authorizationservice"
)

type AuthService interface {
	CheckAccess(userID uint, permission entity.PermissionTitle) (bool, error)
}
type Service struct {
	AuthService AuthService
	Repository  Repository
}

type Repository interface {
	CheckExistCategory(name string) (bool, error)
	AddCategory(category entity.Category) (entity.Category, error)
	GetCategoryByName(name string) (entity.Category, error)
}

func New(cfg cfg.Cfg) Service {
	myRepo := mysql.New(cfg.DataBaseCfg)
	authService := authorizationservice.New(cfg)

	return Service{
		AuthService: authService,
		Repository:  myRepo,
	}
}
