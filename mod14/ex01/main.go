package main

import (
	"fmt"
)

/*
Напишите функцию, которая на вход получает число и возвращает true,
если число четное, и false, если нечётное.


Рекомендация
Программа запрашивает у пользователя или генерирует случайное число,
передает в функцию в качестве аргумента и выводит в консоль результат её работы
*/

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Print("Please input base int: ")
	var answer int
	fmt.Scan(&answer)

	/*
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(answer) // generate randdom string
	*/

	fmt.Println("Input:", answer)

	if isEven(answer) {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
