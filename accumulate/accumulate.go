package accumulate

// Accumulate applies conv to each words in given
func Accumulate(given []string, conv func(string) string) (res []string) {
	for _, s := range given {
		res = append(res, conv(s))
	}
	return
}
