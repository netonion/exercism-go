package summultiples

// SumMultiples returns the sum of multiples of divisros up to limit
func SumMultiples(limit int, divisors ...int) (res int) {
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				res += i
				break
			}
		}
	}
	return
}
