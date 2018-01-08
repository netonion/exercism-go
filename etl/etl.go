package etl

import (
	"strings"
)

// Transform performs extract, transform and load
func Transform(input map[int][]string) map[string]int {
	res := make(map[string]int)
	for i, slice := range input {
		for _, s := range slice {
			res[strings.ToLower(s)] = i
		}
	}
	return res
}
