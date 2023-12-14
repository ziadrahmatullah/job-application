package appvalidator

import "github.com/go-playground/validator/v10"

var Validator AppValidator

type AppValidator interface {
	Validate(objWithStructTag any) error
}

func SetValidator(validator AppValidator) {
	Validator = validator
}

type Impl struct {
	validator *validator.Validate
}

func NewAppValidatorImpl() *Impl {
	return &Impl{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (i *Impl) Validate(objWithStructTag any) error {
	return i.validator.Struct(objWithStructTag)
}
