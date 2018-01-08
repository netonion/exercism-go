package queenattack

import (
	"errors"
)

var errInput = errors.New("incorrect input")

// CanQueenAttack returns if two queens can attack each other
func CanQueenAttack(w, b string) (bool, error) {
	if !validQueen(w) || !validQueen(b) || w == b {
		return false, errInput
	}
	return w[0] == b[0] || w[1] == b[1] || w[0]-b[0] == w[1]-b[1] || w[0]-b[0] == b[1]-w[1], nil
}

func validQueen(q string) bool {
	return len(q) == 2 &&
		q[0] >= byte('a') &&
		q[0] <= byte('h') &&
		q[1] >= byte('1') &&
		q[1] <= byte('8')
}
