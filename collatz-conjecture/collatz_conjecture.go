package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	// panic("Please implement the CollatzConjecture function")

	var steps = 0

	if n <= 0 {
		return steps, errors.New("Incorrect input, should be more than zero")
	}

	for n > 1 {

		if n%2 == 0 {

			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
