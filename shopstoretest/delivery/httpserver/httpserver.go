package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shopstoretest/service/authorizationservice"
	"shopstoretest/service/authservice"
	"shopstoretest/service/categoryservice"
	"shopstoretest/service/productservice"
	"shopstoretest/service/userservice"
)

type Server struct {
	userService          userservice.Service
	authorizationService authorizationservice.Service
	categoryService      categoryservice.Service
	authService          authservice.Service
	productService       productservice.Service
}

func (s Server) Serve() {
	e := echo.New()

	e.GET("/test", s.testHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/users")
	categoryGroup := e.Group("/category")
	basketGroup := e.Group("/basket")
	productGroup := e.Group("/product")

	e.GET("/health-check", s.heathCheck)

	userGroup.POST("/register", s.userRegister)
	userGroup.POST("/login", s.userLogin)
	userGroup.GET("/profile", s.userProfile)

	categoryGroup.POST("/add", s.addCategory)
	basketGroup.POST("/add", s.addBasket)
	basketGroup.POST("/remove", s.removeBasket)

	basketGroup.GET("/show", s.showBaskets)
	productGroup.POST("/order", s.order)
	productGroup.POST("/add", s.addProduct)
	productGroup.GET("/show", s.showProductsByCategory)
	e.Logger.Fatal(e.Start(":5555"))
}

func New(userService userservice.Service, authService authservice.Service,
	authorizationService authorizationservice.Service, categoryService categoryservice.Service,
	productService productservice.Service) Server {
	srv := Server{
		userService:          userService,
		authorizationService: authorizationService,
		categoryService:      categoryService,
		authService:          authService,
		productService:       productService,
	}

	return srv
}
