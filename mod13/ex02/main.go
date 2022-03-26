package main

import (
	"fmt"
	"os"
)

/*
Напишите функцию, которая принимает в качестве аргументов
указатели на два типа int и меняет их значения местами.
*/

func swapPlace(a, b *int) {
	if a == nil || b == nil {
		fmt.Println("There are empty pointers. Exit!")
		os.Exit(255)
	}

	c := *a
	*a = *b
	*b = c
}
func main() {

	var a1, a2 int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a1)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&a2)

	fmt.Printf("a1=%d a2=%d\n", a1, a2)
	swapPlace(&a1, &a2)
	fmt.Printf("a1=%d a2=%d\n", a1, a2)

}
