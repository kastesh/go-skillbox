package main

import (
	"fmt"
	"regexp"
)

/*
Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software.
*/
func main() {
	s := "Go  is an Open source programming Language " +
		"that makes it Easy to build simple, reliable, and efficient Software."

	r, _ := regexp.Compile("[ ,.]+")
	u, _ := regexp.Compile("^[A-Z]")

	count := 0
	for len(s) > 0 {
		postition := r.FindStringIndex(s)
		start := postition[0]
		end := postition[1]

		word := s[:start]
		//fmt.Printf("[%s]\n", word)

		upperPosition := u.FindStringIndex(word)
		if len(upperPosition) > 0 {
			count++
		}
		//fmt.Println(upperPosition)
		s = s[end:]
	}
	fmt.Printf("Строка содержит %d слов с большой буквы.\n", count)
}
