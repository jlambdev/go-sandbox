package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amt, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (amt * ff) / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, cows)
}

type InvalidCowsError struct {
	numCows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			numCows: cows,
			message: "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			numCows: cows,
			message: "no cows don't need food",
		}
	}

	return nil
}
