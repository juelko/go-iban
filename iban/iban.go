package iban

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Errors exported from this file
var ErrInvalidChecksum = errors.New("Invalid checksum")

type IBAN struct {
	cc  *CountryCode
	str string
}

func (iban *IBAN) String() string {
	return iban.str
}

func (iban *IBAN) CountryCode() *CountryCode {
	return iban.cc
}

func (iban *IBAN) Print() string {

	var b strings.Builder

	l := len(iban.str)

	for i := 0; i < l; i = i + 4 {

		if i+4 > l {

			b.WriteString(iban.str[i:])

		} else {

			b.WriteString(" ")
			b.WriteString(iban.str[i : i+4])

		}
	}

	return b.String()
}

func (iban *IBAN) generate(bban string) *IBAN {

	var b strings.Builder

	b.WriteString(iban.cc.String())

	b.WriteString(addChecksum(replaceChars(bban + iban.cc.String())))

	iban.str = b.String()
	return iban
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
	i.SetString(s+"00", 10)
	cs := int(98 - i.Mod(i, big.NewInt(97)).Int64())

	var b strings.Builder
	if cs < 10 {
		b.WriteRune('0')
	}
	b.WriteString(strconv.Itoa(cs))
	b.WriteString(s[:len(s)-4])
	return b.String()
}

func validateChecksum(iban string) error {

	var b strings.Builder
	b.WriteString(iban[4:])
	b.WriteString(iban[:4])

	i := new(big.Int)

	i.SetString(replaceChars(b.String()), 10)

	cs := i.Mod(i, big.NewInt(97)).Int64()
	if cs != 1 {
		return ErrInvalidChecksum
	}

	return nil
}
