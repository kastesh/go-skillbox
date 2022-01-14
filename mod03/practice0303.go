package main

import "fmt"

/*
Допишите программу, которая меняет местами значения переменных.
Добейтесь того, чтобы в переменной "a" лежало значение "b", в "b" лежало значение "a".
*/

func main() {
	fmt.Println("Обмен местами значений переменных")
	a := 42
	b := 153
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	holder := a
	a = b
	b = holder

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
