package main

import "fmt"

func main() {
	fmt.Print("Введите целое число: ")
	var a int
	fmt.Scan(&a)

	var modA int
	if a <= 0 {
		modA = -a
	} else {
		modA = a
	}
	fmt.Println("Модуль числа равен: ", modA)

}
