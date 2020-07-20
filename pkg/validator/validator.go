package validator

import (
	"log"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// Validate data
func Validate(register input) string {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	validate := validator.New()
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}

	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	err := validate.Struct(register)
	var messages = ""
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			messages += e.Translate(trans) + ". "
		}
	}
	return messages
}

// ValidatePassword and retype password
func ValidatePassword(pass string, retype string) string {
	var messages = ""
	if len(pass) < 6 {
		messages += "Password minimum 6 character."
	}

	if pass != retype {
		messages += "Password does not match."
	}

	if IsAlphNum(pass) == false {
		messages += "Password must contains Alphabet and Numeric."
	}

	return messages
}

// IsAlphNum regex a-z A-Z 0-1
var IsAlphNum = regexp.MustCompile(`[a-zA-Z][0-9]+|[0-9][a-zA-Z]+`).MatchString

type input interface{}
