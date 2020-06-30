package responds

import (
	"api_olshop/pkg/dtos"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

// ResSuccess return data
func ResSuccess(w http.ResponseWriter, data dtos.Response) {
	data.APIVersion = viper.Get("api_version").(string)
	data.Code = 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// ResErr return data
func ResErr(w http.ResponseWriter, data dtos.Response, code int) {
	data.APIVersion = viper.Get("api_version").(string)
	data.Code = code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
