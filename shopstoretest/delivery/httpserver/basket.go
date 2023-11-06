package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
)

func (s Server) addBasket(c echo.Context) error {
	req := param.AddToBasketRequest{}
	bErr := c.Bind(&req)
	if bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, errormsg.BadRequest)
	}

	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	if claim.ID != req.UserID {

		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	res, aErr := s.productService.AddBasketItem(req)
	if aErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(aErr))
	}

	return c.JSON(http.StatusCreated, res)
}

func (s Server) removeBasket(c echo.Context) error {
	req := param.RemoveBasketRequest{}
	bErr := c.Bind(&req)
	//fmt.Println(req.ProductID, req.UserID, "\n\n")
	//fmt.Printf("\n", "here", req.UserID, " ", req.ProductID, "\n\n\n", bErr, "\n\n")
	if bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, errormsg.BadRequest)
	}

	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	if claim.ID != req.UserID {

		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	res, aErr := s.productService.RemoveBasketItems(req)
	if aErr != nil {

		return aErr
	}

	return c.JSON(http.StatusOK, res)
}

func (s Server) showBaskets(c echo.Context) error {
	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	req := param.ShowBasketRequest{ID: claim.ID}

	res, sErr := s.productService.ShowBasket(req)
	if sErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(sErr))
	}

	return c.JSON(http.StatusOK, res)
}
