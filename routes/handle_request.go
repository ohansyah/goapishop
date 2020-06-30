package routes

import (
	"api_olshop/internal/contact"
	"api_olshop/middleware"
	"api_olshop/pkg/token"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// HandleRequest handling every request using mux
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	// token
	r.HandleFunc("/api/token/generate", token.Generate).Methods("POST")
	r.HandleFunc("/api/token/validate", token.Validate).Methods("POST")

	// contact
	r.HandleFunc("/api/contact/create", contact.Create).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Connected to port " + viper.Get("port").(string))
	log.Fatal(http.ListenAndServe(":"+viper.Get("port").(string), r))
}
