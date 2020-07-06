package token

import (
	res "api_olshop/pkg/responds"
	"api_olshop/queries"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// GenerateJWT token
func GenerateJWT(appsSecretKey string, appsName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorizes"] = true
	claims["appsName"] = appsName
	claims["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()

	var mySigningKey = []byte(appsSecretKey)
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// Generate and return token
func Generate(w http.ResponseWriter, r *http.Request) {
	var response res.Response
	var data ResToken

	id64, _ := strconv.ParseUint(r.FormValue("device_id"), 10, 64)
	data.DeviceID = uint(id64)
	data.DeviceType = r.FormValue("device_type")

	// validate apps name and secret key
	var appsName = r.FormValue("name")
	var appsSecretKey = r.FormValue("secret_key")
	GetTokenApp := queries.GetTokenApp(appsName, appsSecretKey)
	if GetTokenApp.ID == 0 {
		response.Message = "Apps Name invalid"
		res.ResErr(w, response, http.StatusBadRequest)
		return
	}

	tokenString, err := GenerateJWT(appsSecretKey, appsName)
	if err != nil {
		response.Message = "Error generating token string"
		res.ResErr(w, response, http.StatusBadRequest)
		return
	}
	data.TokenCode = tokenString
	response.Success = true
	response.Data = data
	res.ResSuccess(w, response)
}

// Validate and return boolean
func Validate(w http.ResponseWriter, r *http.Request) {
	var response res.Response
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
		response.Message = err.Error()
		res.ResErr(w, response, http.StatusBadRequest)
		return
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

	var response res.Response
	tokenHeader := r.Header.Get("Token") //Grab the token from the header
	if tokenHeader == "" {               //Token is missing, returns with error code 403 Unauthorized
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
		response.Message = err.Error()
		res.ResErr(w, response, http.StatusForbidden)
		return
	}
}
