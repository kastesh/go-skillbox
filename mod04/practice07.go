package main

import "fmt"

/*
Примеры расчёта:

с дохода 100000 рублей придётся заплатить:
30% × (100000 − 50000) + 20% × (50000 − 10000) + 13% × 10000 =

15 000 + 8000 + 1300 = 24300

с дохода 30000 рублей придётся заплатить:
20% × (30000 − 10000) + 13% × 10000 = 4000 + 1300 = 5300

*/
func main() {
	fmt.Print("Введите Ваш доход: ")
	var income int
	fmt.Scan(&income)

	highTax := 0   // высокий налог
	mediumTax := 0 // средний налог
	minTax := 0    // минимальный налог

	highIncome := income - 50000

	if highIncome > 0 {
		fmt.Println("Повышенный налог с дохода в", highIncome)
		highTax = highIncome * 30 / 100
		fmt.Println("Размер повышенного налога составил", highTax)
		// осталось заплатить
		income = 50000
	}

	mediumIncome := income - 10000
	if mediumIncome > 0 {
		fmt.Println("Средний налог с дохода в", mediumIncome)
		mediumTax = mediumIncome * 20 / 100
		fmt.Println("Размер среднего налога составил", mediumTax)
		// осталось заплатить
		income = 10000
	}

	minTax = income * 13 / 100
	fmt.Println("Обычный налог с дохода в", income)
	fmt.Println("Размер обычного налога составил", minTax)

	totalTax := highTax + mediumTax + minTax
	fmt.Println("Общая сумма налога составила", totalTax)
}
