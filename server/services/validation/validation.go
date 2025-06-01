package validation

import "github.com/go-playground/validator/v10"

type ValidationService interface {
	Validate(sub interface{}) error
}

type goPlayGroundValidator struct {
	v *validator.Validate
}

func New() *goPlayGroundValidator {
	return &goPlayGroundValidator{v: validator.New()}
}

func (g *goPlayGroundValidator) Validate(sub interface{}) error {
	return g.v.Struct(sub)
}
