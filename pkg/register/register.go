package register

import (
	res "api_olshop/pkg/responds"
	"api_olshop/queries"
	"net/http"
)

// GetRoles get data roles
func GetRoles(w http.ResponseWriter, r *http.Request) {
	var response res.Response
	rows := queries.GetRoles(1)
	response.Success = true
	response.Data = rows
	res.ResSuccess(w, r, response)
}

// Register new user
func Register(w http.ResponseWriter, r *http.Request) {
	var response res.Response
	var FormRegis FormRegis
	FormRegis.Name = r.FormValue("name")
	FormRegis.Address = r.FormValue("address")
	FormRegis.RoleID = r.FormValue("role_id")
	FormRegis.Phone = r.FormValue("phone")
	FormRegis.Email = r.FormValue("email")

	// insert register
	queries.Register(FormRegis.Name, FormRegis.Address, FormRegis.RoleID, FormRegis.Phone, FormRegis.Email)

	response.Success = true
	response.Data = "Success"
	res.ResSuccess(w, r, response)
}
