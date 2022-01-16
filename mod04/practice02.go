package main

import "fmt"

func main() {
	limit := 5

	fmt.Print("Введите первое число: ")
	var a1 int
	fmt.Scan(&a1)

	fmt.Print("Введите второе число: ")
	var a2 int
	fmt.Scan(&a2)

	fmt.Print("Введите третье число: ")
	var a3 int
	fmt.Scan(&a3)

	if a1 > limit {
		fmt.Print("Первое число больше ", limit, ".\n")
	} else {
		if a2 > limit {
			fmt.Print("Второе число больше ", limit, ".\n")
		} else {
			if a3 > limit {
				fmt.Print("Третье число больше ", limit, ".\n")
			} else {
				fmt.Print("В введенных числах нет числа больше ", limit, "\n")
			}
		}
	}

}
