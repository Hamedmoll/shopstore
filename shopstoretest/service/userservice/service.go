package userservice

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"shopstoretest/cfg"
	"shopstoretest/entity"
	"shopstoretest/repository/mysql"
	"shopstoretest/service/authservice"
)

type Service struct {
	Repository  Repository
	authService AuthService
}

type AuthService interface {
	CreateAccessToken(id uint, role entity.Role) (string, error)
	CreateRefreshToken(id uint, role entity.Role) (string, error)
	ParseToken(barerToken string) (authservice.Claim, error)
}

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(id uint) (entity.User, error)
}

func New(cfg cfg.Cfg) Service {
	myRepo := mysql.New(cfg.DataBaseCfg)
	authSrv := authservice.New(cfg)

	return Service{
		Repository:  myRepo,
		authService: authSrv,
	}
}

func sha3Hash(input string) string {

	// Create a new hash & write input string
	hash := sha3.New256()
	_, _ = hash.Write([]byte(input))

	// Get the resulting encoded byte slice
	sha3 := hash.Sum(nil)

	// Convert the encoded byte slice to a string
	return fmt.Sprintf("%x", sha3)
}
