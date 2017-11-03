package scrabble

import "strings"

func scrabble(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}

func Score(input string) (score int) {
	for _, letter := range strings.ToUpper(input) {
		score += scrabble(letter)
	}
	return
}
