package main

import (
	"fmt"
)

func main() {
	fmt.Print("Введите первое число: ")
	var t1 int
	fmt.Scan(&t1)

	fmt.Print("Введите второе число: ")
	var t2 int
	fmt.Scan(&t2)

	fmt.Print("Какая сумма двух введенных чисел? ")
	var userSum int
	fmt.Scan(&userSum)
	sum := t1 + t2
	if sum == userSum {
		fmt.Println("Браво! Вы дали верный ответ!")
	} else {
		fmt.Println("Как жаль! Ответ - неверен!")
		fmt.Println("Правильный ответ:", sum)
	}
}
