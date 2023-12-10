package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

func DivideFood(t FodderCalculator, nCows int) (float64, error) {

	fodderAmount, err := t.FodderAmount(nCows)

	if err != nil {
		return fodderAmount, err
	}

	fatteningFactor, err := t.FatteningFactor()

	if err != nil {
		return fatteningFactor, err
	}

	var val = float64(fodderAmount * fatteningFactor / float64(nCows))
	return val, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(t FodderCalculator, nCows int) (float64, error) {

	if nCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(t, nCows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	message string
	nCows   int64
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", e.nCows, e.message)
}

func ValidateNumberOfCows(nCows int) error {
	if nCows == 0 {
		return &InvalidCowsError{nCows: int64(nCows), message: "no cows don't need food"}
	} else if nCows < 0 {
		return &InvalidCowsError{nCows: int64(nCows), message: "there are no negative cows"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
