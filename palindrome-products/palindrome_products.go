// Package palindrome calculates palindrome products
package palindrome

import (
	"errors"
	"strconv"
)

var errRange = errors.New("fmin > fmax")
var errEmpty = errors.New("no palindromes")

// Product reprents a palindrome product with its factorizations
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products accepts minimum and maximum factors and returns the smallest and bigest palindrome products
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errRange
		return
	}

	var found bool
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if product := i * j; isPalindrome(product) {
				if !found || (pmin.Product > product) {
					pmin = Product{product, [][2]int{{i, j}}}
				} else if pmin.Product == product {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}

				if !found || pmax.Product < product {
					pmax = Product{product, [][2]int{{i, j}}}
				} else if pmax.Product == product {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}

				found = true
			}
		}
	}

	if !found {
		err = errEmpty
		return
	}
	return pmin, pmax, nil
}

func isPalindrome(n int) bool {
	if n < 0 {
		n = -n
	}
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
