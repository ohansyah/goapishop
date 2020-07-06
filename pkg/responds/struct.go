package responds

// Response as stuct
type Response struct {
	APIVersion string      `json:"api_version"`
	Code       int         `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// NewResponse create new instance of Response
func NewResponse() Response {
	response := Response{}
	response.APIVersion = "0.0.0"
	response.Code = 400
	response.Success = false
	response.Message = ""
	response.Data = nil
	return response
}
