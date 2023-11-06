package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
)

func (s Server) userRegister(c echo.Context) error {
	req := param.UserRegisterRequest{}

	bErr := c.Bind(&req)
	if bErr != nil {
		//fmt.Println("\n\n\n here \n\n\n")
		return echo.NewHTTPError(http.StatusBadRequest, errormsg.BadRequest)
	}

	res, err := s.userService.Register(req)
	fmt.Printf("\n\n\n %w \n\n\n", err)
	if err != nil {

		return echo.NewHTTPError(errorCodeAndMessage(err))
	}

	return c.JSON(http.StatusCreated, res)
}

func (s Server) userLogin(c echo.Context) error {
	req := param.UserLoginRequest{}

	bErr := c.Bind(&req)
	if bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, errormsg.BadRequest)
	}

	res, lErr := s.userService.Login(req)
	if lErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(lErr))
	}

	return c.JSON(http.StatusOK, res)
}

func (s Server) userProfile(c echo.Context) error {
	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	req := param.UserProfileRequest{ID: claim.ID}

	res, pErr := s.userService.Profile(req)

	if pErr != nil {

		return echo.NewHTTPError(errorCodeAndMessage(pErr))
	}

	return c.JSON(http.StatusOK, res)
}

func (s Server) testHandler(c echo.Context) error {
	category := c.QueryParam("category")

	fmt.Println("\n\n\n", category, "\n\n\n")

	return nil
}
