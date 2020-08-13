// Package raindrops provides functions for the sound of raindrops.
package raindrops

import "strconv"

// Convert returns a raindrop string if the input is a factorial of 3, 5 or 7.
// Or returns the input as a string if not a factorial.
func Convert(input int) string {
	var rain string

	if input%3 == 0 {
		rain += "Pling"
	}
	if input%5 == 0 {
		rain += "Plang"
	}
	if input%7 == 0 {
		rain += "Plong"
	}
	if rain == "" {
		rain = strconv.FormatInt(int64(input), 10)
	}

	return rain
}
