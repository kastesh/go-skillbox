package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
a10 10 20b 20 30c30 30 dd,
*/

func main() {
	s := "a10 10 20b 20 30c30 30 dd"

	res := ""
	/*for index, rune := range s {
		fmt.Printf("%#U starts at byte position %d\n",
			rune,
			index)
	}*/

	for len(s) > 0 {
		s = strings.TrimSpace(s)
		index := strings.Index(s, " ")

		word := s
		if index >= 0 {
			word = s[:index]
			//fmt.Printf("string:[%s]\nlen:[%d]\nindex:[%d]\n", s, len(s), index)
		}

		_, err := strconv.Atoi(word)
		if err == nil {
			res += word + " "
		}

		if index >= 0 {
			s = s[index:]
		} else {
			s = ""
		}
		//fmt.Printf("end:[%s]\n", s)
	}
	fmt.Println("В строке содержатся числа в десятичном формате: ",
		res)
}
