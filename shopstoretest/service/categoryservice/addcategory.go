package categoryservice

import (
	"shopstoretest/entity"
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (s Service) AddCategory(req param.AddCategoryRequest) (param.AddCategoryResponse, error) {
	const op = "category service.add category"
	exist, eErr := s.Repository.CheckExistCategory(req.Name)
	if eErr != nil {

		return param.AddCategoryResponse{}, richerror.New(op).WithError(eErr)
	}

	if exist {

		return param.AddCategoryResponse{}, richerror.New(op).WithKind(richerror.KindNotFound).WithMessage(errormsg.NotFound)
	}

	newCategory := entity.Category{
		ID:   0,
		Name: req.Name,
	}

	createdCategory, aErr := s.Repository.AddCategory(newCategory)
	if aErr != nil {

		return param.AddCategoryResponse{}, richerror.New(op).WithError(aErr)
	}

	categoryInfo := param.CategoryInfo{Name: createdCategory.Name}

	res := param.AddCategoryResponse{
		CategoryInfo: categoryInfo,
	}

	return res, nil
}
