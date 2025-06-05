package products

import (
	"apigo/repository/products"
)

type DeleteProductUsecase struct {
	repository products.DeleteProdcutRepository
}

func NewDeleteProductUsecase(repo products.DeleteProdcutRepository) DeleteProductUsecase {
	return DeleteProductUsecase{
		repository: repo,
	}
}

func (pr *DeleteProductUsecase) Delete(product_id int) error {
	return pr.repository.Delete(product_id)
}
