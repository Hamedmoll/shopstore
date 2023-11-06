package productservice

import (
	"shopstoretest/param"
	"shopstoretest/pkg/richerror"
)

func (s Service) AddProduct(req param.AddProductRequest) (param.AddProductResponse, error) {
	const op = "product service.add product"
	createdProduct, aErr := s.Repository.AddProduct(req)
	if aErr != nil {

		return param.AddProductResponse{}, richerror.New(op).WithError(aErr)
	}

	productInfo := param.ProductInfo{
		Name:        createdProduct.Name,
		Count:       createdProduct.Count,
		Description: createdProduct.Description,
		Category:    req.Category,
		Price:       createdProduct.Price,
	}

	res := param.AddProductResponse{ProductInfo: productInfo}

	return res, nil
}
