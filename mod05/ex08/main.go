package main

import (
	"fmt"
	"os"
)

/*
Ну и какой же компьютер без игр? Давайте научим его играть в «Угадай число».
Пользователь загадывает число от 1 до 10 (включительно).
Программа пытается это число угадать, для этого она выводит число,
а пользователь должен ответить: угадала программа, больше загаданное число или меньше.
*/
func main() {

	fmt.Print(" Прошу загадайте число от 1 до 10 (включительно).\n",
		"Я буду предлагать варианты ответа.\n",
		"Если загаданное тобой число, больше - скажи \"1\"\n",
		"Если загаданное тобой число, меньше - скажи \"-1\"\n",
		"Если я угадала - скажи \"0\"\n")

	// варианты раундов
	minRound := 1
	maxRound := 5

	//
	// варианты цифр
	min := 1
	max := 10
	center := 0
	answer := -2
	for round := minRound; round < maxRound; round++ {
		fmt.Printf("------- Раунд %d ---------\n", round)
		total := 0
		cnt := 0
		for i := min; i <= max; i++ {
			//fmt.Print(i, " ")
			cnt++
			total += i
		}
		//fmt.Print("\n")
		center = total / cnt

		fmt.Printf(" Вы загадали число %d? ", center)
		fmt.Scan(&answer)

		if answer == 0 {
			round = maxRound
		} else if answer == 1 {
			min = center + 1 // 1 => 6
		} else if answer == -1 {
			max = center - 1 // 10 => 4
		} else {
			fmt.Println("Я так не играю!")
			os.Exit(1)
		}
	}

	if answer == 0 {
		fmt.Println("Ура! Мы справились")
	} else {
		fmt.Println("Что-то пошло не так!")
	}
}
