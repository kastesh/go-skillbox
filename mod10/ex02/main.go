package main

import (
	"fmt"
	"math"
)

/*
Пусть пользователь вводит сумму, которую он кладёт в банк,
ежемесячный процент капитализации и количество лет, в течение которых будет открыт вклад.
Необходимо ежемесячно пересчитывать сумму,
округляя её до целого количества копеек в меньшую сторону.
Например, если после начисления процентов остаток включает 35,78 копейки,
то оставляем только 35 копеек, а дробную часть отбрасываем.
По окончании работы программы необходимо вывести,
какую сумму получит вкладчик на руки и какая сумма будет зачислена в пользу банка за счёт округления копеек.



Для 1000 рублей,
1% и одного года программа должна вывести 1126,78 и 0,04350000000022192 (возможно меньше знаков).

Для 1000 рублей, 1% и
десяти лет программа должна вывести 3299,41 и 0,5216000000013992 (возможно меньше знаков).
*/
func main() {

	fmt.Print("Введите сумму вклада: ")
	var savings float64
	fmt.Scan(&savings)

	fmt.Print("Введите количество лет: ")
	var years int
	fmt.Scan(&years)

	fmt.Print("Введите ежемесячный процент: ")
	var percent float64
	fmt.Scan(&percent)

	var months int
	months = years * 12

	totalSavings := savings
	bankPennies := 0.0

	/*
		Для округления копеек можно умножить получившееся после капитализации процентов
		число на 100, округлить в меньшую сторону, затем поделить опять на 100.
		Отбрасываемую часть можно получить вычитанием текущего значения остатка на счёте и того,
		который был до округления.
	*/
	for i := 0; i < months; i++ {

		// количество заработанного
		monthSavings := totalSavings * percent / 100
		monthPennies := math.Trunc(monthSavings*100) / 100
		monthDiff := monthSavings - monthPennies
		totalSavings += monthPennies
		bankPennies += monthDiff
	}
	fmt.Printf("По завершению срока хранения Выы получите: %.2f рублей.\n", totalSavings)
	fmt.Printf("Удержано банком: %.4f\n", bankPennies)

}
