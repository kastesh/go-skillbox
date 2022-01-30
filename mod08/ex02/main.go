package main

import "fmt"

/*
Пользователь вводит будний день недели в сокращённой форме (пн, вт, ср, чт, пт)
и получает развёрнутый список всех последующих рабочих дней, включая пятницу.
*/
func main() {
	fmt.Print("Дни недели.")
	fmt.Print("Введите будний день недели: пн, вт, ср, чт, пт:")
	var day string
	fmt.Scan(&day)

	workDays := ""

	switch day {
	case "пн":
		//workDays = "понедельник\nвторник\nсреда\nчетверг\nпятница\n"
		workDays += "понедельник\n"
		fallthrough
	case "вт":
		workDays += "вторник\n"
		fallthrough
	case "ср":
		workDays += "среда\n"
		fallthrough
	case "чт":
		workDays += "четверг\n"
		fallthrough
	case "пт":
		workDays += "пятница\n"
	default:
		fmt.Println("Не удалось определить рабочие дни для", workDays)
		return
	}

	fmt.Print("Рабочие дни:\n", workDays)
}
