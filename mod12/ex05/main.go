package main

import "fmt"

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func generate(n int, opened int, closed int, answer string, result *[]string) {
	//fmt.Println("-->", answer, opened, closed, n)
	if opened+closed == n*2 {
		fmt.Println(answer)
		*result = append(*result, answer)
		return
	}
	if opened < n {
		generate(n, opened+1, closed, answer+"(", result)
	}
	if opened > closed {
		generate(n, opened, closed+1, answer+")", result)
	}

}
func main() {
	fmt.Print("Введите количество пар: ")
	var pairs int
	fmt.Scan(&pairs)

	result := make([]string, 0)

	generate(pairs, 0, 0, "", &result)

	printSlice(result)
}
