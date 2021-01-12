package iban

import (
	"errors"
	"regexp"
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

type Country struct {
	name    string
	code    CountryCode
	sepa    bool
	lenght  int
	pattern regexp.Regexp
}

func (c Country) Validate(iban string) error {
	if !c.pattern.MatchString(iban) {
		return ErrInvalidPattern
	}

	return validateChecksum(iban)
}

func (c Country) Generate(bban string) (*IBAN, error) {
	if !c.pattern.MatchString(c.code.String() + "00" + bban) {
		return nil, ErrInvalidPattern
	}

	i := IBAN{cc: c.code}

	return i.generate(bban), nil
}

type CountryCode struct {
	code string
}

func (cc *CountryCode) String() string {
	return cc.code
}
