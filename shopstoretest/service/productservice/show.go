package productservice

import (
	"shopstoretest/param"
	"shopstoretest/pkg/errormsg"
	"shopstoretest/pkg/richerror"
)

func (s Service) ShowByCategory(req param.ShowByCategoryRequest) (param.ShowByCategoryResponse, error) {
	const op = "product service.show by category"
	exist, eErr := s.Repository.CheckExistCategory(req.CategoryStr)
	if eErr != nil {

		return param.ShowByCategoryResponse{}, richerror.New(op).WithError(eErr)
	}

	if !exist {

		return param.ShowByCategoryResponse{}, richerror.New(op).WithKind(richerror.KindNotFound).WithMessage(errormsg.NotFound)
	}

	products, gErr := s.Repository.GetProductByCategory(req.CategoryStr)
	if gErr != nil {

		return param.ShowByCategoryResponse{}, richerror.New(op).WithError(gErr)
	}

	res := param.ShowByCategoryResponse{Products: products}

	return res, nil
}

/*func (s Service) ShowByCategory(req param.ShowByCategoryRequest) (param.ShowByCategoryResponse, error) {
	exist, eErr := s.Repository.CheckExistCategory(req.CategoryStr)
	if eErr != nil {

		return param.ShowByCategoryResponse{}, fmt.Errorf("cant check existence %w", eErr)
	}

	if !exist {

		return param.ShowByCategoryResponse{}, fmt.Errorf("category !")
	}
}*/
