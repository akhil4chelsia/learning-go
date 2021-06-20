package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "the quick brown FOX, jumps over the lazy DOG."

	puncCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	for _, ch := range s {
		if unicode.IsPunct(ch) {
			puncCount++
		}
		if unicode.IsLower(ch) {
			lowerCount++
		}
		if unicode.IsUpper(ch) {
			upperCount++
		}
		if unicode.IsSpace(ch) {
			spaceCount++
		}
	}
	fmt.Printf("puncCount: %d, lowerCount: %d, upperCount:%d, spaceCount:%d\n",
		puncCount,
		lowerCount,
		upperCount,
		spaceCount)
}
