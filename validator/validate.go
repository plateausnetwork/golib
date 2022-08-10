//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validate interface {
	Struct(s interface{}) error
	RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
}

type validateImpl struct {
	playground *validator.Validate
}

func NewValidate() Validate {
	return &validateImpl{
		playground: validator.New(),
	}
}

func (v *validateImpl) Struct(s interface{}) error {
	return v.playground.Struct(s)
}

func (v *validateImpl) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.playground.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}
