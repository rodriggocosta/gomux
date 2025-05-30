package products

import (
	"apigo/entity"
	"apigo/repository/products"
)

type ProductCreateUsecase struct {
	repository products.ProductCreateRepository
}

func NewProductsCreateUsecase(repo products.ProductCreateRepository) ProductCreateUsecase {
	return ProductCreateUsecase{
		repository: repo,
	}
}

func (pr *ProductCreateUsecase) Create(product entity.Products) (entity.Products, error) {
	productId, err := pr.repository.Create(product)
	if err != nil {
		return entity.Products{}, err
	}
	product.Product_id = productId

	return product, nil
}
