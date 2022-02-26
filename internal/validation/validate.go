package validation

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func NewValidator(translator ut.Translator) (*validator.Validate, error) {
	validate := validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		return nil, err
	}

	return validate, nil
}

func NewTranslatorEN() (ut.Translator, error) {
	en := en.New()
	uni := ut.New(en, en)
	trans, ok := uni.GetTranslator("en")
	if !ok {
		return nil, errors.New("failed to get default english translation for validator")
	}

	return trans, nil
}
