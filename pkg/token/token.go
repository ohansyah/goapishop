package token

import (
	res "api_olshop/pkg/responds"
	"api_olshop/queries"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// GenerateJWT token
func GenerateJWT(exptime time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorizes"] = true
	claims["exp"] = exptime.Unix()

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
	var response res.Response
	var data ResToken
	exptime := time.Now().Add(time.Hour * 24)
	data.DeviceID = r.FormValue("device_id")
	data.DeviceType = r.FormValue("device_type")

	// validate apps name and secret key
	if r.FormValue("secret_key") != viper.Get("secret_key") {
		response.Message = "secret_key invalid"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	tokenString, err := GenerateJWT(exptime)
	tokenRefresh, err := GenerateJWT(time.Now().Add(time.Hour * 24 * 30))
	if err != nil {
		response.Message = "Error generating token string"
		res.ResErr(w, r, response, http.StatusBadRequest)
		return
	}

	// check if device has token
	deviceToken := queries.GetTokenByDevID(data.DeviceID)
	if deviceToken.ID > 0 {
		// update if has
		queries.UpdateToken(deviceToken.ID, tokenString, tokenRefresh, exptime)

	} else {
		// create if not
		queries.CreateToken(data.DeviceID, data.DeviceType, tokenString, tokenRefresh, exptime)
	}

	data.TokenCode = tokenString
	data.RefreshToken = tokenRefresh
	response.Success = true
	response.Data = data

	// set header
	r.Header.Set("Token", tokenString)
	r.Header.Set("RefreshToken", tokenRefresh)

	res.ResSuccess(w, r, response)
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

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response.Success = true
		response.Message = "true"
		res.ResSuccess(w, r, response)

	} else {
		response.Message = err.Error()
		res.ResErr(w, r, response, http.StatusForbidden)
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

		res.ResErr(w, r, response, http.StatusForbidden)
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
		res.ResErr(w, r, response, http.StatusForbidden)
		return
	}
}
