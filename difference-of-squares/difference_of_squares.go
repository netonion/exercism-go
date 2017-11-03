package diffsquares

func SquareOfSums(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
