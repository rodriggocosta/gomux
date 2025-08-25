package router

import (
	"apigo/handlers"
	"apigo/infra"
	"apigo/repository"
	"apigo/usecase"
	"log"
	"net/http"

	handlerAddress "apigo/handlers/address"
	repositoryAddress "apigo/repository/address"
	usecaseAddress "apigo/usecase/address"

	handlerProduct "apigo/handlers/products"
	repositoryProduct "apigo/repository/products"
	usecaseProduct "apigo/usecase/products"

	handlerUser "apigo/handlers/users"
	repositoryUser "apigo/repository/users"
	usecaseUser "apigo/usecase/users"
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

	addressRepository := repositoryAddress.NewAddressCreate(dbConnection)
	addressUasecas := usecaseAddress.NewAddress(repositoryAddress.AddressPostReropsitory(addressRepository))
	addressHandler := handlerAddress.NewAddressHandler(addressUasecas)

	getAddressRespository := repositoryAddress.NewGetAddress(dbConnection)
	getAddressUsecase := usecaseAddress.NewAddressUsecase(getAddressRespository)
	getAddressHandler := handlerAddress.NewGetAddresshandler(getAddressUsecase)

	updateAddressRepository := repositoryAddress.NewUpdateAddressRepository(dbConnection)
	upadateAddressUsecase := usecaseAddress.NewPutAddress(updateAddressRepository)
	updateAddressHandler := handlerAddress.NewPutAddressHandler(upadateAddressUsecase)

	deleteAddressRepository := repositoryAddress.NewDeleteAddressRespository(dbConnection)
	deleteAddressUsecase := usecaseAddress.NewDeleteAddress(repositoryAddress.DeleteAddressRepository(deleteAddressRepository))
	deleAddressHandler := handlerAddress.NewDeleteAddressHandler(deleteAddressUsecase)

	userGetRepository := repositoryUser.NewUserGet(dbConnection)
	userGetUsecase := usecaseUser.NewUserGet(userGetRepository)
	userGetHandler := handlerUser.NewUserHandler(userGetUsecase)

	r.HandleFunc("/endereco/cadastrar", addressHandler.Create)
	r.HandleFunc("/endereco", getAddressHandler.GetAddress)
	r.HandleFunc("/endereco/editar", updateAddressHandler.PutAddress)
	r.HandleFunc("/endereco/excluir", deleAddressHandler.Delete)

	r.HandleFunc("/user", userGetHandler.GetUser)

	return r
}
