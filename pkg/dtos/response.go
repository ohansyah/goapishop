package dtos

// Response as stuct
type Response struct {
	APIVersion string      `json:"api_version"`
	Code       int         `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
