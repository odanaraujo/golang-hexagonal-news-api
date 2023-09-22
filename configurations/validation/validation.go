package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/exception"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		_ = en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *exception.Exception {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return exception.NewBadRequestError("Invalid field type")
	}
	if errors.As(validation_err, &jsonValidationError) {
		var errorsCauses []exception.Causes
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := exception.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return exception.NewBadRequestValidationError("some fields are invalid", errorsCauses)
	}
	return exception.NewBadRequestError("error trying to convert fields")
}
