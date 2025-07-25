package router

import (
	"apigo/handlers"
	"apigo/infra"
	"apigo/repository"
	"apigo/usecase"
	"log"
	"net/http"

	handlerAddress "apigo/handlers/address"
	repostiroyAddress "apigo/repository/address"
	usecaseAddress "apigo/usecase/address"

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

	updateProductRepository := repositoryProduct.NewByIdProductRepository(dbConnection)
	updateProductUsecase := usecaseProduct.NewProductByIdUsecase(repositoryProduct.ProductRepository(updateProductRepository))
	updateProductHalnder := handlerProduct.NewProductByIdHandler(updateProductUsecase)

	r.HandleFunc("/produtos", productHandler.GetProduct)
	r.HandleFunc("/produtos/cadastrar", createProductHandler.Create)
	r.HandleFunc("/produtos/editar", putProductHandler.PutProduct)
	r.HandleFunc("/produto/delete", deleteHandlerProduct.Delete)
	r.HandleFunc("/produto", updateProductHalnder.GetByIdProduct)

	addressRepository := repostiroyAddress.NewAddressCreate(dbConnection)
	addressUasecas := usecaseAddress.NewAddress(repostiroyAddress.AddressPostReropsitory(addressRepository))
	addressHandler := handlerAddress.NewAddressHandler(addressUasecas)

	getAddressRespository := repostiroyAddress.NewGetAddress(dbConnection)
	getAddressUsecase := usecaseAddress.NewAddressUsecase(getAddressRespository)
	getAddressHandler := handlerAddress.NewGetAddresshandler(getAddressUsecase)

	r.HandleFunc("/endereco/cadastrar", addressHandler.Create)
	r.HandleFunc("/endereco", getAddressHandler.GetAddress)

	return r
}
