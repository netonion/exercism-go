package grains

import "errors"

const numSquares = 64

func Square(n int) (uint64, error) {
	if n < 1 || n > numSquares {
		return 0, errors.New("Input must be within 1 to 64.")
	}
	return 1 << uint(n-1), nil
}

func Total() uint64 {
	return 1<<numSquares - 1
}
