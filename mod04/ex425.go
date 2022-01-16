package main

import "fmt"

/*
Каждый из 9000 мужчин посещает барбершоп раз в месяц (30 дней)
Один барбер способен обслужить одного клиента за 1 час
Смена барбера — 8 часов
*/

func main() {
	fmt.Print("Сколько мужчин проживает в городе: ")
	var man int
	fmt.Scan(&man)

	fmt.Print("Сколько барберов уже работает в городе: ")
	var barber int
	fmt.Scan(&barber)

	// все барберы мугут обслужит за месяц такое количество людей
	oneBarberPerf := 8 * 30
	barberPerformancePerMonth := barber * oneBarberPerf

	// сколько слотов нужно под текущее количество мужчин
	manNeedsPerMonth := man

	fmt.Println("Такое количество барберов сможет обслужить", barberPerformancePerMonth, "за месяц.")
	if manNeedsPerMonth <= barberPerformancePerMonth {
		fmt.Println("Количество барберов достаточно для стрижки всех бород в городе")
	} else {
		fmt.Println("Количество барберов недостаточно")

		barberMaxInt := manNeedsPerMonth / oneBarberPerf
		barberMaxDevide := manNeedsPerMonth / oneBarberPerf
		if barberMaxDevide > 0 {
			barberMaxInt += 1
		}
		fmt.Println("Мы должны увеличить количество барберов до:", barberMaxInt)
		fmt.Println("Такое количество барберов сможет обслужить", barberMaxInt*oneBarberPerf, "человек.")
	}

}
