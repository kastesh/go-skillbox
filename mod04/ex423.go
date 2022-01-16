package main

import "fmt"

func main() {
	fmt.Println("Расстояние от Москвы до Рязани 200км")
	fmt.Print("Введите среднюю скорость (км/ч), с которой двидется автомобиль: ")
	var speed int
	fmt.Scan(&speed)
	time := 200 / speed
	if time <= 2 {
		fmt.Println("Вы успели!")
	}
	fmt.Println("Автомобиль проедет данное расстрояние за:", 200/speed, "ч.")

}
