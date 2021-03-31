package hamming

import (
	"errors"
)

// Distance function that calculates the strand differences in the DNA and computes a Hamming Distance between two strands
func Distance(a, b string) (int, error) {
	distance := 0
	rA := []rune(a)
	rB := []rune(b)

	if len(rA) != len(rB) {
		return distance, errors.New("length of strand a or b do not match")
	}

	for i := range rA {
		if rA[i] != rB[i] {
			distance++
		}

	}

	return distance, nil
}
