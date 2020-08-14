// Package scrabble manages functions for the game of Scrabble
package scrabble

import "strings"

var scrableScores = map[string]int{
	"a": 1,
	"e": 1,
	"i": 1,
	"o": 1,
	"u": 1,
	"l": 1,
	"n": 1,
	"r": 1,
	"s": 1,
	"t": 1,
	"d": 2,
	"g": 2,
	"b": 3,
	"c": 3,
	"m": 3,
	"p": 3,
	"f": 4,
	"h": 4,
	"v": 4,
	"w": 4,
	"y": 4,
	"k": 5,
	"j": 8,
	"x": 8,
	"q": 10,
	"z": 10,
}

// Score returns an int representing the scrabble word score of a given string.
func Score(word string) int {
	sanitizedWord := strings.ToLower(word)

	var score int
	for i := 0; i < len(word); i++ {
		score += scrableScores[string(sanitizedWord[i])]
	}

	return score
}
