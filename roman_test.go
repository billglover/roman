package roman_test

import (
	"errors"
	"testing"

	"github.com/billglover/roman"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var toRomanTests = []struct {
	from int
	to   string
	err  error
}{
	{1, "I", nil},
	{2, "II", nil},
	{3, "III", nil},
	{4, "IV", nil},
	{5, "V", nil},
	{6, "VI", nil},
	{7, "VII", nil},
	{8, "VIII", nil},
	{9, "IX", nil},
	{1954, "MCMLIV", nil},
	{4999, "MMMMCMXCIX", nil},
	{0, "", errors.New("unable to convert zero to Roman numerals")},
	{-1, "", errors.New("unable to convert negative number to Roman numerals")},
}

func TestToRoman(t *testing.T) {
	for _, tt := range toRomanTests {
		out, err := roman.FromInt(tt.from)
		if out != tt.to {
			t.Fatalf("testing %d, expected '%s' got '%s' %v", tt.from, tt.to, out, ballotX)
		}
		t.Logf("testing %d, expected '%s' got '%s' %v", tt.from, tt.to, out, checkMark)

		if tt.err != nil {
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Fatalf("expected error: '%s' got '%s' %v", tt.err.Error(), err.Error(), ballotX)
				}
			}

			if err == nil {
				t.Fatalf("expected error: '%s' got '%s' %v", tt.err.Error(), err.Error(), ballotX)
			}

			t.Logf("got error: '%s' %v", err.Error(), checkMark)
		}
	}
}

var validTests = []struct {
	number string
	valid  bool
}{
	{"I", true},
	{"V", true},
	{"X", true},
	{"L", true},
	{"C", true},
	{"D", true},
	{"M", true},
	{"A", false},
	{"", false},
	{"IA", false},
	{"AI", false},
	{"MMMCMXCIX", true},
	{"III", true},
	{"IIII", false},
	{"VV", false},
	{"XXX", true},
	{"XXXX", false},
	{"LL", false},
	{"CCC", true},
	{"CCCC", false},
	{"DD", false},
	{"MMM", true},
	{"MMMM", false},
	{"CXXX", true},
	{"XCXL", false},
	{"IXX", false},
	{"XIX", true},
}

func TestValid(t *testing.T) {
	for _, tt := range validTests {
		out := roman.IsValid(tt.number, true)
		if out != tt.valid {
			t.Errorf("testing '%s', expected '%v' got '%v' %v", tt.number, tt.valid, out, ballotX)
		} else {
			t.Logf("testing '%s', expected '%v' got '%v' %v", tt.number, tt.valid, out, checkMark)
		}
	}
}

var toIntTests = []struct {
	from int
	to   string
	err  error
}{
	{1, "I", nil},
	{2, "II", nil},
	{3, "III", nil},
	{4, "IV", nil},
	{5, "V", nil},
	{6, "VI", nil},
	{7, "VII", nil},
	{8, "VIII", nil},
	{9, "IX", nil},
	{1954, "MCMLIV", nil},
	{3999, "MMMCMXCIX", nil},
	{0, "", errors.New("unable to parse Roman numeral")},
	{0, "ABCD", errors.New("unable to parse Roman numeral")},
	{0, "IXX", errors.New("unable to parse Roman numeral")},
}

func TestToInt(t *testing.T) {
	for _, tt := range toIntTests {
		out, err := roman.ToInt(tt.to)
		if out != tt.from {
			t.Fatalf("testing '%s', expected %d got %d %v", tt.to, tt.from, out, ballotX)
		}
		t.Logf("testing '%s', expected %d got %d %v", tt.to, tt.from, out, checkMark)
		if tt.err != nil {
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Fatalf("expected error: '%s' got '%s' %v", tt.err.Error(), err.Error(), ballotX)
				}
			}

			if err == nil {
				t.Fatalf("expected error: '%s' got '%s' %v", tt.err.Error(), err.Error(), ballotX)
			}

			t.Logf("got error: '%s' %v", err.Error(), checkMark)
		}
	}
}
