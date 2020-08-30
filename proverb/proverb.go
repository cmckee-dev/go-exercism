// Package proverb should have a package comment that summarizes what it's about.
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return []string{}
	}

	if len(rhyme) == 1 {
		return []string{fmt.Sprintf("And all for the want of a %s.", rhyme[0])}
	}

	var proverb []string

	for i := 0; i < len(rhyme)-1; i++ {
		str := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverb = append(proverb, str)
	}

	str := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverb = append(proverb, str)
	return proverb
}
