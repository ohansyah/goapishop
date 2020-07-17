package responds

import (
	"api_olshop/models"
	"api_olshop/queries"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// ResSuccess return data
func ResSuccess(w http.ResponseWriter, r *http.Request, data Response) {
	data.APIVersion = viper.Get("api_version").(string)
	data.Code = 200

	CreateTokenLog(r, data, "success")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// ResErr return data
func ResErr(w http.ResponseWriter, r *http.Request, data Response, code int) {
	data.APIVersion = viper.Get("api_version").(string)
	data.Code = code

	CreateTokenLog(r, data, "fail")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// CreateTokenLog create token log
func CreateTokenLog(r *http.Request, data Response, Status string) {
	// save response
	jsonData, _ := json.Marshal(data)
	var tokenString = strings.Join(r.Header["Token"], ", ")
	var tokendata = queries.GetTokenData(tokenString)

	tokenlog := models.TokenLog{
		TokenID:    tokendata.ID,
		UserAgent:  r.Header.Get("User-Agent"),
		Path:       r.RequestURI,
		Method:     r.Method,
		Request:    r.RequestURI,
		Response:   string(jsonData),
		Status:     Status,
		APIVersion: data.APIVersion,
	}
	queries.CreateTokenLog(tokenlog)
}
