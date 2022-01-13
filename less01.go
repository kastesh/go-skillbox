package main

import "fmt"

func main() {
	fmt.Println("Введите целое число")
	var i int
	fmt.Scan(&i)
	fmt.Println("Квадрат числа", i, "равен", i*i)
}
