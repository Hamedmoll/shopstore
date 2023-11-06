package userservice

import (
	"fmt"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (s Service) Login(req param.UserLoginRequest) (param.UserLoginResponse, error) {
	const op = "user service.login"

	user, gErr := s.Repository.GetUserByPhoneNumber(req.PhoneNumber)

	if gErr != nil {

		return param.UserLoginResponse{}, fmt.Errorf("unexpected error %w", gErr)
	}

	if sha3Hash(req.Password) != user.Password {

		return param.UserLoginResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errormsg.InvalidUser)
	}

	userInfo := param.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Credit:      user.Credit,
	}

	accessToken, aErr := s.authService.CreateAccessToken(userInfo.ID, user.Role)
	if aErr != nil {

		return param.UserLoginResponse{}, richerror.New(op).WithError(aErr)
	}

	refreshToken, rErr := s.authService.CreateRefreshToken(userInfo.ID, user.Role)
	if rErr != nil {

		return param.UserLoginResponse{}, richerror.New(op).WithError(rErr)
	}

	tokens := param.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	res := param.UserLoginResponse{
		UserInfo: userInfo,
		Tokens:   tokens,
	}

	return res, nil
}
