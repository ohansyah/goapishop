package routes

import (
	"api_olshop/internal/contact"
	"api_olshop/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequest handling every request using mux
func HandleRequest() {
	r := mux.NewRouter()

	// token
	// r.HandleFunc("/token/get", homePage).Methods("GET")

	// product
	// r.HandleFunc("/api/product", product.List).Methods("GET")
	// r.HandleFunc("/api/product/get", product.Get).Methods("GET")
	// r.HandleFunc("/api/product/create", product.Create).Methods("POST")
	// r.HandleFunc("/api/product/update", product.Update).Methods("POST")
	// r.HandleFunc("/api/product/delete", product.Delete).Methods("POST")

	// contact
	r.HandleFunc("/api/contact/create", contact.Create).Methods("POST")

	http.Handle("/", r)
	r.Use(middleware.LoggingMiddleware)
	fmt.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
