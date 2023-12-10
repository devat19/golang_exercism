package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// panic("Implement the Distance function")
	if len(a) != len(b) {
		return 0, errors.New("both strings should be of same length to calculate hamming distance")
	}
	if a == b { // understand the ramifications of this - memory address and very large  strings
		return 0, nil
	}

	var distance = 0

	for i := 0; i < len(a); i++ {

		if a[i] == b[i] {
			continue
		}
		distance++
	}
	return distance, nil
}
