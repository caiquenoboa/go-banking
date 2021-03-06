package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/caiquenoboa/go-banking/domain"
	"github.com/caiquenoboa/go-banking/service"
	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
)

func sanityChecker() {
	variables := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASSWORD", "DB_ADDRESS", "DB_PORT", "DB_NAME"}

	for _, variable := range variables {
		if os.Getenv(variable) == "" {
			log.Fatal("Variable Enviroment " + variable + " is not defined")
		}
	}
}

func Start() {

	sanityChecker()

	router := mux.NewRouter()

	customerHandler, accountHandler := initHandlers()

	//define routes
	router.
		HandleFunc("/customers", customerHandler.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.newAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	authMiddleware := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(authMiddleware.authorizationHandler())

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	//starting server
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}

func initHandlers() (customerHandler CustomerHandlers, accountHandler AccountHandler) {
	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()

	customerRepository := domain.NewCustomerRepositoryDb(dbClient)
	accountRepository := domain.NewAccountRepositoryDb(dbClient)

	custumerService := service.NewCustomerService(customerRepository)
	accountService := service.NewAccountService(accountRepository)

	customerHandler = CustomerHandlers{custumerService}
	accountHandler = AccountHandler{accountService}

	return
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
