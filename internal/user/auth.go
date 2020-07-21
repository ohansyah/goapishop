package user

import (
	"api_olshop/models"
	res "api_olshop/pkg/responds"
	validator "api_olshop/pkg/validator"
	"api_olshop/queries"
	"net/http"
	"strings"
	"time"
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

	// validate token
	var tokenString = strings.Join(r.Header["Token"], ", ")
	var tokenData = queries.GetTokenData(tokenString)
	if tokenData.ID == 0 {
		// failed login
		response.Message = "Login failed, pleace contact customer services"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// validate token profiles
	var tokenProfiles = queries.GetTokenProfile(tokenData.ID)
	if tokenProfiles.ID == 0 {
		// create token profile
		tokenProfileData := models.TokenProfile{
			TokenID:      tokenData.ID,
			UserID:       user.ID,
			LastActivity: time.Now(),
		}
		queries.CreateTokenProfile(tokenProfileData)
	} else {
		// update token profile
		queries.UpdateTokenProfile(tokenProfiles.ID, tokenData.ID, user.ID)
	}

	response.Success = true
	response.Data = user
	res.ResSuccess(w, r, response)
}
