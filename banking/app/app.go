package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func Start() {

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/costomers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
