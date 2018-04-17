package allergies

var substances = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

// Allergies returns list of allergies based on score
func Allergies(score uint) (res []string) {
	for i, a := range substances {
		if score&(1<<uint(i)) > 0 {
			res = append(res, a)
		}
	}
	return
}

// AllergicTo tests if allergic score includes substance
func AllergicTo(score uint, substance string) bool {
	for _, a := range Allergies(score) {
		if substance == a {
			return true
		}
	}
	return false
}
