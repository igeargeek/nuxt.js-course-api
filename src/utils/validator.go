package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var trans ut.Translator

func NewValidateTranslation() ut.Translator {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)
	}

	return trans
}

func SetInvalidFields(errs validator.ValidationErrors) []map[string]string {
	invalidFields := make([]map[string]string, 0)
	for _, e := range errs {
		errors := map[string]string{}
		errors[ToSnakeCase(e.Field())] = e.Translate(trans)
		invalidFields = append(invalidFields, errors)
	}
	return invalidFields
}
