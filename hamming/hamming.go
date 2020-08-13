// Package hamming provides the Hamming Distance function for DNA calculations.
package hamming

import "errors"

// Distance calculates the number of differences between two strands of DNA.
func Distance(a, b string) (int, error) {
	rs1, rs2 := []rune(a), []rune(b)

	if len(rs1) != len(rs2) {
		return 0, errors.New("dna strands are not the same length")
	}

	var distance int
	for i := 0; i < len(rs1); i++ {
		if rs1[i] != rs2[i] {
			distance++
		}
	}

	return distance, nil
}
