package wordcount

import (
	"strings"
	"unicode"
)

// Frequency type
type Frequency map[string]int

// WordCount counts the frequency of words in the input
func WordCount(input string) Frequency {
	input = strings.Replace(input, ",", " ", -1)
	freq := make(Frequency)
	for _, word := range strings.Fields(strings.ToLower(input)) {
		trimmed := strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		freq[trimmed]++
	}
	return freq
}
