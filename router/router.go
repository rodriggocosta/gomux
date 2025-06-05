package router

import (
	"apigo/handlers"
	"apigo/infra"
	"apigo/repository"
	"apigo/usecase"
	"log"
	"net/http"

	handlerProduct "apigo/handlers/products"
	repositoryProduct "apigo/repository/products"
	usecaseProduct "apigo/usecase/products"
)

func NewRouter() *http.ServeMux {

	r := http.NewServeMux()

	dbConnection, err := infra.Connect()
	if err != nil {
		log.Fatalf("Error ao conectar ao database: %v", err)
	}

	//	defer dbConnection.Close()

	customerRepo := repository.NewCustomerRepository(dbConnection)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	customerHandlers := handlers.NewCustomerHandler(customerUsecase)

	r.HandleFunc("/clientes", customerHandlers.GetCustomer)
	r.HandleFunc("/clientes/cadastrar", customerHandlers.PostCustomer)
	r.HandleFunc("/cliente", customerHandlers.GetCustomerById)
	r.HandleFunc("/cliente/delete", customerHandlers.DeleteById)
	r.HandleFunc("/cliente/editar", customerHandlers.Update)

	productRepository := repositoryProduct.NewProductRepository(dbConnection)
	productUsecase := usecaseProduct.NewProductUsecase(productRepository)
	productHandler := handlerProduct.NewProductHandler(productUsecase)

	createProductsRepository := repositoryProduct.NewProductsCreateRepository(dbConnection)
	createProductUsercase := usecaseProduct.NewProductsCreateUsecase(repositoryProduct.ProductCreateRepository(createProductsRepository))
	createProductHandler := handlerProduct.NewProductsCreateHandlers(createProductUsercase)

	putProductRepository := repositoryProduct.NewPutProductRepository(dbConnection)
	putProductUsecase := usecaseProduct.NewPutProduct(repositoryProduct.PutProductRepository(putProductRepository))
	putProductHandler := handlerProduct.NewPutProudcHandler(putProductUsecase)

	deleteProductRepository := repositoryProduct.NewDeleteProductRepository(dbConnection)
	deleteProductUsecase := usecaseProduct.NewDeleteProductUsecase(repositoryProduct.DeleteProdcutRepository(deleteProductRepository))
	deleteHandlerProduct := handlerProduct.NewDeleteProductHandler(deleteProductUsecase)

	r.HandleFunc("/produtos", productHandler.GetProduct)
	r.HandleFunc("/produtos/cadastrar", createProductHandler.Create)
	r.HandleFunc("/produtos/editar", putProductHandler.PutProduct)
	r.HandleFunc("/produtos/delete", deleteHandlerProduct.Delete)

	return r
}
