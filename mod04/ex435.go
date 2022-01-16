package main

import "fmt"

func main() {
	fmt.Println("Пожалуйста, загадайте число то 1 до 5!")
	fmt.Println("Я буду спрашивать Вас, если угадала нажмите 1, если нет - 0")
	fmt.Print("Число больше трех? ")
	var a int
	fmt.Scan(&a)

	if a == 0 {
		// x <= 3; 1 2 3
		fmt.Print("Число больше двух? ")
		fmt.Scan(&a)
		if a == 1 {
			fmt.Println("Вы загадали 3!")
		} else if a == 0 {
			fmt.Println("Число меньше 2?")
			fmt.Scan(&a)
			if a == 0 {
				fmt.Println("Вы загадали 2!")
			} else if a == 1 {
				fmt.Println("Вы загадал 1!")
			}
		}
	}

}
