package contact

import (
	"api_olshop/queries"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Create contact mux
func Create(w http.ResponseWriter, r *http.Request) {

	// get data
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")

	// save operation
	response := queries.CreateContact(name, email, phone, address)

	// res
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}
