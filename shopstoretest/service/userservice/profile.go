package userservice

import (
	"shopstoretest/param"
	"shopstoretest/pkg/richerror"
)

func (s Service) Profile(req param.UserProfileRequest) (param.UserProfileResponse, error) {
	const op = "user service.profile"

	user, gErr := s.Repository.GetUserByID(req.ID)
	if gErr != nil {

		return param.UserProfileResponse{}, richerror.New(op).WithError(gErr)
	}

	userInfo := param.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Credit:      user.Credit,
	}

	return param.UserProfileResponse{UserInfo: userInfo}, nil
}
