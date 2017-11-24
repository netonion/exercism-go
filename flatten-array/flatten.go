package flatten

// Flatten nested slices recursively
func Flatten(i interface{}) []interface{} {
	res := []interface{}{}
	switch t := i.(type) {
	case []interface{}:
		for _, item := range t {
			res = append(res, Flatten(item)...)
		}
	case interface{}:
		res = append(res, t)
	}
	return res
}
