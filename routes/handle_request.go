package routes

import (
	"api_olshop/internal/contact"
	"api_olshop/middleware"
	"api_olshop/token"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequest handling every request using mux
func HandleRequest() {
	r := mux.NewRouter()

	// token
	r.HandleFunc("/api/token/gnerate", token.Generate).Methods("POST")

	// contact
	r.HandleFunc("/api/contact/create", contact.Create).Methods("POST")

	http.Handle("/", r)
	r.Use(middleware.LoggingMiddleware)
	fmt.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
