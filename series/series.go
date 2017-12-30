package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (res []string) {
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n if n is valid.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}
	return UnsafeFirst(n, s), true
}
