package pangram

import "unicode"

// IsPangram returns weather the input string is a pangram
func IsPangram(input string) bool {
	counter := make(map[rune]bool)
	for _, r := range input {
		counter[unicode.ToLower(r)] = true
	}
	for r := 'a'; r <= 'z'; r++ {
		if !counter[r] {
			return false
		}
	}
	return true
}
