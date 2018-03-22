package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode encrypts input string using crypto square method
func Encode(pt string) string {
	var rs []rune
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			rs = append(rs, unicode.ToLower(r))
		}
	}
	r := int(math.Sqrt(float64(len(rs))))
	if r*(r+1) < len(rs) {
		r++
	}
	c := r
	if r*c < len(rs) {
		c++
	}

	var sb strings.Builder
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if k := j*c + i; k < len(rs) {
				sb.WriteRune(rs[k])
			} else {
				sb.WriteRune(' ')
			}
		}
		if i < c-1 {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}
