package roman

import (
	"errors"
	"strings"
)

var numerals = []struct {
	d int
	r string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// FromInt takes an integer and returns a string representation
// of that integer in Roman numerals.
func FromInt(i int) (s string, e error) {

	if i == 0 {
		e = errors.New("unable to convert zero to Roman numerals")
		return
	}

	if i < 0 {
		e = errors.New("unable to convert negative number to Roman numerals")
		return
	}

	for _, n := range numerals {
		for i >= n.d {
			s += n.r
			i -= n.d
			if i == 0 {
				return
			}
		}
	}

	return
}

// ToInt takes a string of Roman numerals and returns the integer
// representation of the number
//
// Note: ToInt assumes that the input is a valid Roman numberal
func ToInt(s string) (i int, e error) {

	if IsValid(s, false) != true {
		e = errors.New("unable to parse Roman numeral")
		return
	}

	for len(s) > 0 {
		for _, n := range numerals {
			if strings.HasPrefix(s, n.r) {
				i += n.d
				s = strings.Replace(s, n.r, "", 1)
			}
		}
	}

	return
}

// IsValid takes a string of Roman numerals and returns the a
// boolean to indicate whether the string is a valid representation
// of a number
//
// Note: setting strict = true will validate that the efficient
//       form of a number of used. This requires decoding and
//       re-encoding the number to determine it's efficient form
//       and this is significantly slower.
//
// Rules:
// - only the characters I, V, X, L, C, D, and M are valid
// - powers of 10 (I, X, C, and M) can only repeat three times
// - powers of 5 cannot repeat
// - numerals (or valid pairs) should appear in decending order
func IsValid(s string, strict bool) (v bool) {

	// empty strings aren't considered valid
	if s == "" {
		return false
	}

	// valid Roman numerals (and their decimal representation)
	validNums := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	// maximum number of times a numeral can appear consecutively
	countNums := map[string]uint{
		"I": 3, "V": 1, "X": 3, "L": 1, "C": 3, "D": 1, "M": 3,
	}

	var lastChar string
	var consecutiveCount uint

	// loop over the string representation of the Roman numerals and
	// check the following
	//  - each character is in hte valid set
	//  - it doesn't violet the limit on consecutive occurances
	for _, c := range s {

		if _, ok := validNums[string(c)]; ok == false {
			return false
		}

		if string(c) == lastChar {
			consecutiveCount++

			if consecutiveCount >= countNums[string(c)] {
				return false
			}
		} else {
			consecutiveCount = 0
		}
		lastChar = string(c)
	}

	// set beyond the maximum value of any given numeral
	lastNum := 20000
	roman := s

	// validate that numerals appear in descending order
	for len(roman) > 0 {
		for _, n := range numerals {
			if strings.HasPrefix(roman, n.r) {

				if n.d > lastNum {
					return false
				}

				roman = strings.Replace(roman, n.r, "", 1)
				lastNum = n.d
			}
		}
	}

	if strict {
		i, _ := ToInt(s)
		r, _ := FromInt(i)
		if r != s {
			return false
		}

	}

	return true

}
