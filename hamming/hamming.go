// Package hamming provides the Hamming Distabce function for DNA calculations.
package hamming

import (
	"errors"
)

// Distance calculates the number of differences between two strands of DNA.
func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return -1, errors.New("DNA strands are not the same length.")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
