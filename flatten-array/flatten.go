package flatten

import (
	"reflect"
)

// Flatten nested slices recursively
func Flatten(i interface{}) []interface{} {
	res := []interface{}{}
	s := i.([]interface{})
	for _, item := range s {
		if reflect.TypeOf(item) != nil {
			if reflect.TypeOf(item).Kind() == reflect.Slice {
				res = append(res, Flatten(item)...)
			} else {
				res = append(res, item)
			}
		}
	}
	return res
}
