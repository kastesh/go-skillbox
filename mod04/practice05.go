package main

import "fmt"

/*
по понедельникам должна
применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый;
по пятницам, если сумма чека превышает 10 000 рублей,
включается дополнительная скидка в размере 5%;
если число гостей в одной компании превышает пять человек, а
втоматически включается надбавка на обслуживание 10%.
*/
func main() {
	fmt.Print("Введите день недели (пн - 1, вт - 2 и т.д.): ")
	var day int
	fmt.Scan(&day)

	fmt.Print("Введите сумму чека: ")
	var check int
	fmt.Scan(&check)

	fmt.Print("Ввведите количество гостей: ")
	var guests int
	fmt.Scan(&guests)

	guestsOverehead := 5
	plusGuests := 0
	// в любой день увеличивется стоимость, если гостей больше 5
	if guests >= guestsOverehead {
		fmt.Println("Сумма чека увеличена на 10% из-за количество гостей.")
		plusGuests = check * 10 / 100
	}
	// по понедельникам
	byMondey := 0
	if day == 1 {
		fmt.Println("Сумма чека уменьшена на 10%. Хвала понедельнику!")
		byMondey = check * 10 / 100
	}

	byFridey := 0
	//  пятница и сумма
	if day == 5 && check >= 10000 {
		fmt.Println("Сумма чека менььшена на 5%. Пятница и большой чек!")
		byFridey = check * 5 / 100
	}

	check = check + plusGuests - byMondey - byFridey
	fmt.Print("Итоговая сумма: ", check, " рублей.\n")
}
