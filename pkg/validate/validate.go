package validate

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	// Trans translate validate msg
	Trans ut.Translator
	// Validation validator.Validate
	Validation *validator.Validate
)

func init() {
	e := en.New()
	uni := ut.New(e, e)
	Trans, _ = uni.GetTranslator("zh")
	Validation = validator.New()
	en_translations.RegisterDefaultTranslations(Validation, Trans)
}

// Validate the field
func Validate(s interface{}) error {
	err := Validation.Struct(s)
	if err == nil {
		return nil
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors.New("interface conversion failed")
	}
	var msg string
	rType := reflect.TypeOf(s)
	for _, e := range errs {
		field := e.Field()
		fieldZhName, _ := rType.FieldByName(field)
		msg += strings.Replace(e.Translate(Trans), field, fieldZhName.Tag.Get("zh"), -1)
	}
	return errors.New(msg)
}
