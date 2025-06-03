package products

import (
	"apigo/entity"
	"apigo/repository/products"
)

type PutProductUsecase struct {
	repository products.PutProductRepository
}

func NewPutProduct(repo products.PutProductRepository) PutProductUsecase {
	return PutProductUsecase{
		repository: repo,
	}
}

func (pr *PutProductUsecase) PutProduct(product_id int, product *entity.Products) error {
	return pr.repository.PutProduct(product_id, product)
}
