package iban

import (
	"errors"
)

var (
	ErrInvalidPattern = errors.New("Does not match pattern")
)

type IBANRegistry interface {
	Validator(cc CountryCode) (Validator, error)
	Generator(cc CountryCode) (Generator, error)
}

type Validator interface {
	Validate(iban string) bool
}

type Generator interface {
	Generate(bban string) *IBAN
}

type registry struct {
	countries map[CountryCode]Country
}

func (r *registry) Validator(cc CountryCode) (Validator, error) {
	return nil, errors.New("Not implemented")
}
func (r *registry) Generator(cc CountryCode) (Generator, error) {
	return nil, errors.New("Not implemented")
}
