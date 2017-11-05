package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)

	if len(input) < 2 {
		return false
	}

	if len(input)%2 == 1 {
		input = "0" + input
	}

	sum := 0
	for i, c := range input {
		if !unicode.IsDigit(c) {
			return false
		}
		d := int(c - '0')
		if i%2 == 0 && d != 9 {
			d = 2 * d % 9
		}
		sum += d
	}
	return sum%10 == 0
}
