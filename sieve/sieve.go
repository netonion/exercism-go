package sieve

// Sieve generates all the prime numbers under limit
func Sieve(limit int) (res []int) {
	multiples := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if !multiples[i] {
			res = append(res, i)
			for j := i * i; j <= limit; j += i {
				multiples[j] = true
			}
		}
	}
	return
}
