package products

import (
	"apigo/entity"
	"apigo/repository/products"
	"fmt"
)

type ProductByIdUsecase struct {
	repository products.ByIdProductRepository
}

func NewProductByIdUsecase(repo products.ProductRepository) ProductByIdUsecase {
	return ProductByIdUsecase{
		repository: products.ByIdProductRepository(repo),
	}
}

func (pr *ProductByIdUsecase) GetByIdProduct(product_id int) (*entity.Products, error) {
	product, err := pr.repository.ByIdProduct(product_id)

	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar o produto com Id %d: %w", product_id, err)
	}

	return product, nil
}
