package iban

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var ErrInvalidChecksum = errors.New("Invalid checksum")

type CountryCode struct {
	code string
}

func (cc *CountryCode) String() string {
	return cc.code
}

type IBAN struct {
	cc  CountryCode
	str string
}

func (i *IBAN) String() string {
	return i.str
}

func (i *IBAN) CountryCode() CountryCode {
	return i.cc
}

func (i *IBAN) generate(bban string) *IBAN {

	var b strings.Builder

	b.WriteString(replaceChars(bban))

	b.WriteString("00")

	i.str = i.cc.String() + addChecksum(b.String())
	return i
}

func replaceChars(s string) string {

	var b strings.Builder

	for _, v := range s {
		if v >= 65 && v <= 90 {
			fmt.Fprintf(&b, "%v", v-55)
		}
		if v >= 48 && v <= 57 {
			fmt.Fprintf(&b, "%v", v-48)
		}
	}

	return b.String()
}

func addChecksum(s string) string {
	i := new(big.Int)
	i.SetString(s, 10)
	cs := int(98 - i.Mod(i, big.NewInt(97)).Int64())

	var b strings.Builder
	if cs < 10 {
		b.WriteRune('0')
	}
	b.WriteString(strconv.Itoa(cs))
	b.WriteString(s)
	return b.String()
}

func validateChecksum(iban string) error {

	i := new(big.Int)

	i.SetString(replaceChars(iban), 10)

	cs := i.Mod(i, big.NewInt(97)).Int64()
	if cs != 1 {
		return ErrInvalidChecksum
	}

	return nil
}
