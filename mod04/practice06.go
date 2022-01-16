package main

import "fmt"

/*
Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов,
на K групп. Напишите программу, которая поможет старосте сделать это:
он вводит N, K и порядковый номер студента, а программа определяет, в какую группу он попадёт.
*/
func main() {

	fmt.Print("Введите количество людей на курсе: ")
	var people int
	fmt.Scan(&people)

	fmt.Print("Введите количество групп: ")
	var groups int
	fmt.Scan(&groups)

	fmt.Print("Введите номер студента: ")
	var student int
	fmt.Scan(&student)

	groupSize := people / groups
	leftPeople := people % groups
	if leftPeople > 0 {
		groupSize += 1
	}
	fmt.Println("Количество людей в одной группе:", groupSize)

	studentGroup := (student % groupSize) + 1
	fmt.Println("Номер группы для студента:", studentGroup)

}
