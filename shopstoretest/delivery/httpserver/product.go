package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/entity"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
)

func (s Server) addProduct(c echo.Context) error {
	req := param.AddProductRequest{}
	bErr := c.Bind(&req)
	if bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, errormsg.BadRequest)
	}

	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	haveAccess, hErr := s.authorizationService.CheckAccess(claim.ID, entity.AddProductPermission)
	if hErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(hErr))
	}

	if !haveAccess {

		return echo.NewHTTPError(http.StatusUnauthorized, errormsg.DontHaveAccess)
	}

	res, err := s.productService.AddProduct(req)
	if err != nil {

		return echo.NewHTTPError(errorCodeAndMessage(err))
	}

	return c.JSON(http.StatusCreated, res)
}

func (s Server) showProductsByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	req := param.ShowByCategoryRequest{CategoryStr: category}

	res, sErr := s.productService.ShowByCategory(req)
	if sErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(sErr))
	}

	return c.JSON(http.StatusOK, res)
}
