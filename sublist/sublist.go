package sublist

type Relation string

func Sublist(l1, l2 []int) Relation {
	switch {
	case len(l1) == len(l2) && isSub(l1, l2):
		return "equal"
	case len(l1) > len(l2) && isSub(l2, l1):
		return "superlist"
	case isSub(l1, l2):
		return "sublist"
	}
	return "unequal"
}

func isSub(l1, l2 []int) bool {
	for j := 0; j < len(l2)-len(l1)+1; j++ {
		match := true
		for i, n := range l1 {
			if n != l2[i+j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
