package perfect

import (
	"errors"
)

// ErrOnlyPositive represents input less than or equal to zero
var ErrOnlyPositive = errors.New("only positive")

// Classification type
type Classification int

const (
	ClassificationAbundant  Classification = iota // Abundant
	ClassificationDeficient                       // Deficient
	ClassificationPerfect                         // Perfect
)

// Classify determins if a number is perfect, abundant, or deficient
func Classify(n int64) (res Classification, err error) {
	if n <= 0 {
		err = ErrOnlyPositive
		return
	}

	sum := int64(1)
	for i := int64(2); i*i < n; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}

	if sum > n {
		res = ClassificationAbundant
	} else if n == 1 || sum < n {
		res = ClassificationDeficient
	} else {
		res = ClassificationPerfect
	}

	return
}
