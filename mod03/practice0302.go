package main

import "fmt"

/*
Напишите программу, которая симулирует работу маршрутки.
Программа умеет объявлять остановки и узнавать у пользователя,
сколько человек вышло на этой остановке и сколько зашло. Реализуйте маршрут из четырёх остановок,
а в конце рассчитайте, сколько денег заработано, при условии, что билет стоит 20 рублей.

Расходы следующие:
четверть — на зарплату водителя;
пятая часть — на топливо;
пятая часть — на налоги;
пятая часть — на ремонт машины.
*/
func main() {
	ticketPrice := 20
	earnedAmount := 0

	pasEntered := 0
	pasLeft := 0
	pasCount := 0
	pasTotal := 0

	fmt.Println("Программа симуляции работы маршрутного такси.")
	fmt.Println("============================")
	// start: обработка остановки 1
	street := "Улица Программистов"
	fmt.Print("Первая остановка «", street, "».\n")
	fmt.Print("В салоне пассажиров: ", pasCount, "\n")
	fmt.Print("Пассажиров вышло на остановке: ", pasLeft, "\n")
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&pasEntered)
	fmt.Print("Отправляемся с остановки «", street, "».\n")
	pasCount += (pasEntered - pasLeft)
	pasTotal += pasEntered
	fmt.Print("В салоне пассажиров: ", pasCount, "\n\n")

	fmt.Println("-----------Едем---------")
	// start: обработка остановки 2
	street = "Пойди туда не знаю куда"
	fmt.Print("Прибываем на остановку «", street, "».\n")
	fmt.Print("В салоне пассажиров: ", pasCount, "\n")
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&pasLeft)
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&pasEntered)
	fmt.Print("Отправляемся с остановки «", street, "».\n")
	pasCount += (pasEntered - pasLeft)
	pasTotal += pasEntered
	fmt.Print("В салоне пассажиров: ", pasCount, "\n\n")

	// start: обработка остановки 3
	street = "На деревню к дедушке"
	fmt.Print("Прибываем на остановку «", street, "».\n")
	fmt.Print("В салоне пассажиров: ", pasCount, "\n")
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&pasLeft)
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&pasEntered)
	fmt.Print("Отправляемся с остановки «", street, "».\n")
	pasCount += (pasEntered - pasLeft)
	pasTotal += pasEntered
	fmt.Print("В салоне пассажиров: ", pasCount, "\n\n")

	// start: обработка остановки 4
	street = "Гадюкино"
	fmt.Print("Прибываем на остановку «", street, "».\n")
	fmt.Print("В салоне пассажиров: ", pasCount, "\n")
	// финальная выходят все и никто не входит
	fmt.Print("Пассажиров вышло на остановке: ", pasCount, "\n")
	fmt.Print("Пассажиров вошло на остановке: ", 0, "\n")
	fmt.Print("Конечная остановка «", street, "».\n\n")

	fmt.Println("-----------Считаем---------")
	earnedAmount = pasTotal * ticketPrice
	salary := earnedAmount / 4
	fuel := earnedAmount / 5
	taxes := earnedAmount / 5
	repair := earnedAmount / 5
	revenue := earnedAmount - (salary + fuel + taxes + repair)

	fmt.Print("Количество человек на маршруте: ", pasTotal, "\n")
	fmt.Print("Всего заработали: ", earnedAmount, " рублей.\n")
	fmt.Print("Зарплата водителя: ", salary, " рублей.\n")
	fmt.Print("Расходы на топливо: ", fuel, " рублей.\n")
	fmt.Print("Налоги: ", taxes, " рублей.\n")
	fmt.Print("Расходы на ремонт машины: ", repair, " рублей.\n")
	fmt.Print("Итого доход: ", revenue, " рублей.\n")
}
