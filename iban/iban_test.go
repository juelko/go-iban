package iban

import (
	"errors"
	"testing"
)

func TestValidateChecksum(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		want error
	}{
		{
			desc: "Happy path",
			in:   "FI2112345600000785",
			want: nil,
		},
		{
			desc: "Err",
			in:   "FI2212345600000785",
			want: ErrInvalidChecksum,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			got := validateChecksum(tC.in)

			if !errors.Is(got, tC.want) {
				t.Errorf("want: %s , got: %s", tC.want, got)
			}

		})
	}
}

func TestGenerate(t *testing.T) {
	testCases := []struct {
		desc string
		iban *IBAN
		in   string
		want string
	}{
		{
			desc: "Happy path",
			iban: &IBAN{cc: &CountryCode{code: "FI"}},
			in:   "12345600000785",
			want: "FI2112345600000785",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.iban.generate(tC.in)

			if got.String() != tC.want {
				t.Errorf("want: %s , got: %s", tC.want, got)
			}
		})
	}
}
