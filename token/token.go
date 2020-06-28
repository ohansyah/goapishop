package token

import (
	"api_olshop/dtos"
	"encoding/json"
	"fmt"
	"net/http"
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
	claims["exp"] = time.Now().Add(time.Second * 3).Unix()

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
	tokenString, err := GenerateJWT()
	if err != nil {
		response.Success = false
		response.Message = "Error generating token string"
	} else {
		response.Success = true
		response.Data = tokenString
	}

	// res
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}
