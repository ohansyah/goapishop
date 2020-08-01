package user

import (
	res "api_olshop/pkg/responds"
	validator "api_olshop/pkg/validator"
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
	register := FormRegis{
		Name:    r.FormValue("name"),
		Address: r.FormValue("address"),
		RoleID:  r.FormValue("role_id"),
		Phone:   r.FormValue("phone"),
		Email:   r.FormValue("email"),
		Status:  "1",
	}

	// validate input data
	validate := validator.Validate(register)
	if validate != "" {
		response.Message = validate
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// validate password
	validatePassword := validator.ValidatePassword(r.FormValue("password"), r.FormValue("retype_password"))
	if validatePassword != "" {
		response.Message = validatePassword
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// validate duplicate email and phone
	checkRegister := queries.GetUserByEmailPhone(register.Email, register.Phone)
	if checkRegister.ID > 0 {
		response.Message = "Email or Phone Number exist"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// bycript password
	pwd := []byte(r.FormValue("password"))
	register.Password = HashAndSalt(pwd)

	// insert register
	user := queries.Register(register.Name, register.Address, register.RoleID, register.Phone, register.Email, register.Password)
	if user.ID == 0 {
		response.Message = "Registration Failed!"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	response.Success = true
	response.Data = user
	res.ResSuccess(w, r, response)
}
