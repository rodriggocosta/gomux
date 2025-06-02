package products

import (
	"apigo/entity"
	"apigo/repository/products"
	"errors"

	"github.com/shopspring/decimal"
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
	priceDec, err := decimal.NewFromString(product.Price)
	if err != nil {
		return entity.Products{}, errors.New("Invalid price format")
	}
	if priceDec.Cmp(decimal.Zero) <= 0 {
		return entity.Products{}, errors.New("Price must be greater than zero")
	}

	productId, err := pr.repository.Create(product)
	if err != nil {
		return entity.Products{}, err
	}

	product.Product_id = productId

	return product, nil
}
