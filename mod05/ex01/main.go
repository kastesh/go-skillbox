package main

import "fmt"

/*
Напишите программу, которая поможет пользователю определить, к какой координатной четверти принадлежит точка.
Пользователь вводит числа X и Y, а программе необходимо вывести, какой координатной четверти принадлежит точка.
*/
func main() {
	fmt.Print("Введите координаты для x: ")
	var x int
	fmt.Scan(&x)

	fmt.Print("Введите координаты для y: ")
	var y int
	fmt.Scan(&y)

	quadrant := "unknown"
	if x > 0 && y > 0 {
		quadrant = "I"
	} else if x < 0 && y > 0 {
		quadrant = "II"
	} else if x < 0 && y < 0 {
		quadrant = "III"
	} else if x > 0 && y < 0 {
		quadrant = "IV"
	}

	if quadrant == "unknown" {
		fmt.Print("Точка (", x, ",", y, ") - это начало координат.\n")
	} else {
		fmt.Print("Точка (", x, ",", y, ") принадлежит ", quadrant, " четверти.\n")
	}
}
