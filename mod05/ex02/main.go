package main

import "fmt"

/*
Проверка пользовательского ввода на различные ограничения является частой задачей.
Попросите пользователя ввести три числа и проверьте, что хотя бы одно является положительным.
Результат проверки необходимо сообщить пользователю.
*/
func main() {
	fmt.Print(" Введите первое число: ")
	var a1 int
	fmt.Scan(&a1)

	fmt.Print(" Введите второе число: ")
	var a2 int
	fmt.Scan(&a2)

	fmt.Print(" Введите третье число: ")
	var a3 int
	fmt.Scan(&a3)

	if a1 > 0 || a2 > 0 || a3 > 0 {
		fmt.Println("Одно из трех числе положительное")
	} else {
		fmt.Println("Все числа отрицательные")
	}
}
