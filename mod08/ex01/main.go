package main

import "fmt"

func main() {
	fmt.Print("Введите название месяца: ")
	var month string
	fmt.Scan(&month)

	var season string

	switch month {
	case "декабрь", "январь", "февраль":
		season = "зима"
	case "март", "апрель", "май":
		season = "весна"
	case "июнь", "июль", "август":
		season = "лето"
	case "сентябрь", "октябрь", "ноябрь":
		season = "осень"
	default:
		fmt.Println("Не удалось определить время года для месяца", month)
		return
	}
	fmt.Println("Время года", season)

}
