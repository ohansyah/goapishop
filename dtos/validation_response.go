package dtos

// ValidationResponse as struct
type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validations`
}
