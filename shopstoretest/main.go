package main

import (
	_ "github.com/go-sql-driver/mysql"
	"shopstoretest/cfg"
	"shopstoretest/delivery/httpserver"
	"shopstoretest/service/authorizationservice"
	"shopstoretest/service/authservice"
	"shopstoretest/service/categoryservice"
	"shopstoretest/service/productservice"
	"shopstoretest/service/userservice"
	"time"
)

func main() {
	dataBaseCfg := cfg.DataBaseConfig{
		DataBaseHost:     cfg.DatabaseHost,
		DataBasePort:     cfg.DatabasePort,
		DataBaseName:     cfg.DatabaseName,
		DataBaseUser:     cfg.DatabaseUser,
		DataBasePassword: cfg.DatabasePassword,
		DataBaseProtocol: cfg.DatabaseProtocol,
	}

	serverCfg := cfg.ServerConfig{
		ServerHost: cfg.ServerHost,
		ServerPort: cfg.ServerPort,
	}

	authCfg := cfg.AuthConfig{
		SignKey:               cfg.JwtSignKey,
		AccessExpirationTime:  time.Hour * 6,
		RefreshExpirationTime: time.Hour * 24 * 7,
		AccessSubject:         "at",
		RefreshSubject:        "rt",
	}

	myCfg := cfg.New(dataBaseCfg, serverCfg, authCfg)

	uService, authService, authorizationService, categoryService, productService := setupServices(myCfg)

	server := httpserver.New(uService, authService, authorizationService, categoryService, productService)

	//req := param.LoginRequest{
	//	PhoneNumber: "091322234488",
	//	Password:    "123",
	//}

	//myRepo := mysql.New(dataBaseCfg)
	//Success, err := myRepo.Login(req.PhoneNumber, req.Password)

	//fmt.Println(Success, err, "\n\n\n")

	//uInfo, gErr := myRepo.GetProfileByPhoneNumber(req.PhoneNumber)

	//fmt.Println(uInfo, gErr)

	//res, err := myRepo.DB.Exec("insert into users(name, phone_number, hashed_password) values(?, ?, ?)", "hamed", "0912", "123")

	//fmt.Println(res, err)
	server.Serve()
}

func setupServices(cfg cfg.Cfg) (userservice.Service, authservice.Service,
	authorizationservice.Service, categoryservice.Service, productservice.Service) {
	userService := userservice.New(cfg)
	authService := authservice.New(cfg)
	authorizationService := authorizationservice.New(cfg)
	categoryService := categoryservice.New(cfg)
	productService := productservice.New(cfg)

	return userService, authService, authorizationService, categoryService, productService
}
