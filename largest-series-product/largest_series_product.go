package lsproduct

import (
	"errors"
)

var err = errors.New("input incorrect")

// LargestSeriesProduct comoputs the largest product of a contiguous substring of digits of length n
func LargestSeriesProduct(digits string, n int) (int, error) {
	if n < 0 || len(digits) < n {
		return -1, err
	}
	max := 0
	for i := 0; i <= len(digits)-n; i++ {
		prod := 1
		for j := 0; j < n; j++ {
			d := int(digits[i+j] - '0')
			if d > 9 { // d is never less than zero, because byte is uint8
				return -1, err
			}
			prod *= d
		}
		if prod > max {
			max = prod
		}
	}
	return max, nil
}
