package main

import "fmt"

func main() {
	startHight := 100 // начальный размер
	growth := 50      // прирост за день
	losses := 20      // гусеницы сьели за ночь
	growthForSale := 300
	fmt.Println("-------------Начальные условия-------------")
	fmt.Print("Начальный размер бамбука (см): ", startHight, "\n")
	fmt.Print("Прирост за день (см): ", growth, "\n")
	fmt.Print("Потери за ночь (см): ", losses, "\n")
	fmt.Print("Рост для продажи(см): ", growthForSale, "\n")

	higthThirdDay := startHight + (growth-losses)*2 + growth
	fmt.Print("Высота на третий день: ", higthThirdDay, "\n")

	numberDays := (growthForSale - startHight) / (growth - losses)
	hightNumberDays := startHight + (growth-losses)*numberDays + growth
	fmt.Println("Чтобы продать бамбук нужно подождать", numberDays, "дней.")
	fmt.Println("В конце", numberDays, "размер бабмука", hightNumberDays, "см.")

	fmt.Println("-------------Поехали-------------")
	fmt.Println("Сейчас Вы сможете выбрать условия для бабмбука!")
	fmt.Print("Введите начальный размер бамбука(см): ")
	fmt.Scan(&startHight)
	fmt.Print("Прирост за день(см): ")
	fmt.Scan(&growth)
	fmt.Print("Потери за ночь(см): ")
	fmt.Scan(&losses)
	fmt.Print("Высота бамбука, достаточная для продажи(см): ")
	fmt.Scan(&growthForSale)

	dayNumber := 3
	fmt.Print("На какой день будем смотреть его рост (см): ")
	fmt.Scan(&dayNumber)

	higthThirdDay = startHight + (growth-losses)*(dayNumber-1) + growth
	fmt.Print("Высота на ", dayNumber, " день: ", higthThirdDay, "\n")

	numberDays = (growthForSale - startHight) / (growth - losses)
	hightNumberDays = startHight + (growth-losses)*numberDays + growth
	fmt.Println("Чтобы продать бамбук нужно подождать", numberDays, "дней.")
	fmt.Println("В конце", numberDays, "размер бабмука", hightNumberDays, "см.")

}
