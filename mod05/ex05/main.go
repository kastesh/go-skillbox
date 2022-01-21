package main

import "fmt"

/*
Напишите программу, которая позволит пользователю ввести три процентные ставки
и выведет в консоль две наиболее выгодные,
а если все процентные ставки равны — сообщит, что ставки равнозначны.

Задачу можно решить несколькими способами — например, от противного.
*/
func main() {

	fmt.Print(" Введите ставку первого банка: ")
	var a1 float64
	fmt.Scan(&a1)

	fmt.Print(" Введите ставку второго банка: ")
	var a2 float64
	fmt.Scan(&a2)

	fmt.Print(" Введите ставку третьего банка:: ")
	var a3 float64
	fmt.Scan(&a3)

	position1 := a1
	position2 := a2
	position3 := a3

	if a3 > a1 && a3 > a2 {
		position1 = a3
		position3 = a1
	} else if a2 > a1 {
		position1 = a2
		position2 = a1
	}
	if position3 > position2 {
		position2, position3 = position3, position2
	}

	//fmt.Println(position1, position2, position3)

	if position1 == position2 && position2 == position3 {
		fmt.Println("Ставки во всех банках равнозначны.")
	} else {
		fmt.Print("Наиболее выгодные ставки - ", position1, "% и ", position2, "%.\n")
	}
}
