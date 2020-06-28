package token

import (
	"api_olshop/dtos"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	ResSuccess(w, response)
}

// ResSuccess return data
func ResSuccess(w http.ResponseWriter, data dtos.Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusOK)
}

// 	// /*
// 	// 	Parse takes the token string and a function for looking up the key.
// 	// 	The latter is especially useful if you use multiple keys for your application.
// 	// 	The standard is to use 'kid' in thebhead of the token to identify which key to use,
// 	// 	but the parsed token (head and claims) is provided to the callback, providing flexibility.
// 	// */
// 	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 	// 	// Don't forget to validate the alg is what you expect:
// 	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	// 	}

// 	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 	// 	return mySigningKey, nil
// 	// })

// 	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 	// 	fmt.Println(claims["foo"], claims["nbf"])
// 	// } else {
// 	// 	fmt.Println(err)
// 	// }
