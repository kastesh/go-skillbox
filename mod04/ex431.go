package main

import "fmt"

func main() {
	fmt.Print("Введите первое число: ")
	var t1 int
	fmt.Scan(&t1)

	fmt.Print("Введите второе число: ")
	var t2 int
	fmt.Scan(&t2)

	if t1 > t2 {
		fmt.Print(t1, ">", t2, "\n")
	} else if t1 == t2 {
		fmt.Print(t1, "=", t2, "\n")
	} else {
		fmt.Print(t1, "<", t2, "\n")
	}
}
