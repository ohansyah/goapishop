package token

import (
	"api_olshop/pkg/dtos"
	res "api_olshop/pkg/responds"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// GenerateJWT token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorizes"] = true
	claims["user"] = "ohan"
	// claims["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()

	var mySigningKey = []byte(viper.Get("api_key").(string))
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// Generate and return token
func Generate(w http.ResponseWriter, r *http.Request) {
	var response dtos.Response
	var data ResToken

	id64, _ := strconv.ParseUint(r.FormValue("device_id"), 10, 64)
	data.DeviceID = uint(id64)
	data.DeviceType = r.FormValue("device_type")

	tokenString, err := GenerateJWT()
	if err != nil {
		response.Success = false
		response.Message = "Error generating token string"
	} else {
		data.TokenCode = tokenString
		response.Success = true
	}

	response.Data = data
	res.ResSuccess(w, response)
}

// Validate and return boolean
func Validate(w http.ResponseWriter, r *http.Request) {
	var response dtos.Response
	var tokenString = strings.Join(r.Header["Token"], ", ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.Get("api_key").(string)), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		response.Success = true
		response.Message = "true"
		res.ResSuccess(w, response)

	} else {
		response.Success = false
		response.Message = err.Error()
		res.ResErr(w, response, http.StatusBadRequest)
	}

}

// ValidateToken return data
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// token checks
	notAuth := []string{"/api/token/generate"} //List of endpoints that doesn't require auth
	requestPath := r.URL.Path                  //current request path
	//check if request does not need authentication, serve the request if it doesn't need it
	for _, value := range notAuth {

		if value == requestPath {
			next.ServeHTTP(w, r)
			return
		}
	}

	var response dtos.Response
	tokenHeader := r.Header.Get("Token") //Grab the token from the header
	if tokenHeader == "" {               //Token is missing, returns with error code 403 Unauthorized
		response.Success = false
		response.Message = "Missing auth token"

		res.ResErr(w, response, http.StatusForbidden)
		return
	}

	// start validate token
	var tokenString = strings.Join(r.Header["Token"], ", ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.Get("api_key").(string)), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		next.ServeHTTP(w, r)
	} else {
		response.Success = false
		response.Message = err.Error()
		res.ResErr(w, response, http.StatusForbidden)
		return
	}
}
