package products

import (
	"apigo/entity"
	"apigo/repository/products"
)

type ProductUsecase struct {
	repository products.ProductRepository
}

func NewProductUsecase(repo products.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pr *ProductUsecase) GetProduct() ([]entity.Products, error) {
	return pr.repository.GetProduct()
}
