package userservice

import (
	"shopstoretest/entity"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (s Service) Register(req param.UserRegisterRequest) (param.UserRegisterResponse, error) {
	const op = "user service.register"
	unique, uErr := s.Repository.IsPhoneNumberUnique(req.PhoneNumber)

	if uErr != nil {

		return param.UserRegisterResponse{}, richerror.New(op).WithError(uErr)
	}

	if !unique {

		return param.UserRegisterResponse{}, richerror.New(op).WithKind(richerror.KindNotUnique).WithMessage(errormsg.NotUniquePhoneNumber)
	}

	newUser := entity.User{
		ID:          0,
		Role:        entity.UserRole,
		Name:        req.Name,
		Password:    sha3Hash(req.Password),
		Credit:      0,
		PhoneNumber: req.PhoneNumber,
	}

	createdUser, cErr := s.Repository.Register(newUser)

	if cErr != nil {

		return param.UserRegisterResponse{}, richerror.New(op).WithError(cErr)
	}

	userInfo := param.UserInfo{
		ID:          createdUser.ID,
		Name:        createdUser.Name,
		PhoneNumber: createdUser.PhoneNumber,
		Credit:      0,
	}

	res := param.UserRegisterResponse{
		UserInfo: userInfo,
	}

	return res, nil
}
