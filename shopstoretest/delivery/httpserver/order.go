package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
)

func (s Server) order(c echo.Context) error {
	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)

	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	req := param.AddOrderRequest{}
	cErr := c.Bind(&req)

	if cErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, errormsg.BadRequest)
	}

	if claim.ID != req.UserID {

		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	res, oErr := s.productService.Order(req)
	//fmt.Printf("\n\n %v here \n\n", oErr)

	if oErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(oErr))
	}

	return c.JSON(http.StatusOK, res)
}
