package raindrops

import "strconv"

func Convert(num int) string {
	res := ""
	if num%3 == 0 {
		res += "Pling"
	}
	if num%5 == 0 {
		res += "Plang"
	}
	if num%7 == 0 {
		res += "Plong"
	}
	if len(res) == 0 {
		return strconv.Itoa(num)
	}
	return res
}
