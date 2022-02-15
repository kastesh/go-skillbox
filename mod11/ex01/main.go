package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software.
*/
func main() {
	s := "Go  is an Open source programming Language aбракАдабра" +
		"that makes it Easy to build simple, reliable, and efficient Software."

	isFirstLetter := 1
	countUpperCases := 0
	for _, rune := range s {
		alfa := strconv.QuoteRune(rune)
		if isFirstLetter == 1 {
			isFirstLetter = 0
			if unicode.IsUpper(rune) {
				fmt.Println(alfa)
				countUpperCases++
			}
		}

		if unicode.IsSpace(rune) || unicode.IsPunct(rune) {
			isFirstLetter = 1
		}
	}
	fmt.Printf("Строка содержит %d слов с большой буквы.\n", countUpperCases)
}
