package helpers

import (
	"api_olshop/dtos"
	"api_olshop/langs"

	"gopkg.in/go-playground/validator.v8"
)

// GenerateValidateResponse function
func GenerateValidateResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.Validation

	// get validation error
	// var validationErrors := err.(validator.ValidationErrors)
	var validationErrors = err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		// get field and rule (tag)
		field, rule := value.Field, value.Tag

		// create validation object
		validation := dtos.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}

		// add validation object to validations
		validations = append(validations, validation)
	}

	// set Validation response
	response.Validations = validations

	return response
}
