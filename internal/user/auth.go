package user

import (
	res "api_olshop/pkg/responds"
	validator "api_olshop/pkg/validator"
	"api_olshop/queries"
	"net/http"
)

// Login user
func Login(w http.ResponseWriter, r *http.Request) {
	var response res.Response
	register := FormLogin{
		Phone:    r.FormValue("username"),
		Email:    r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// validate input data
	validate := validator.Validate(register)
	if validate != "" {
		response.Message = validate
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// validate email and phone
	user := queries.GetUserByEmailPhone(register.Email, register.Phone)
	if user.ID == 0 {
		response.Message = "Email or Phone Number not found"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// ComparePasswords password
	plain := []byte(r.FormValue("password"))
	if ComparePasswords(user.Password, plain) == false {
		response.Message = "Your password is incorrect"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// insert token profile

	response.Success = true
	response.Data = user
	res.ResSuccess(w, r, response)
}
