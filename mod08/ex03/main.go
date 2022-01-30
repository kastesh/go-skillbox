package main

import (
	"fmt"
	"strings"
)

/*
Напишите функцию, которая посчитает,
сможет ли продавец в киоске обслужить всех покупателей.

В киоске каждый лимонад стоит пять долларов.
Клиенты стоят в очереди, чтобы купить у вас, и заказывают по одному лимонаду.
Каждый покупатель может купить только один лимонад и заплатить купюрами номиналом 5, 10 или 20 долларов.
Вы должны дать каждому покупателю сдачу с его купюры.

Обратите внимание, что сначала у вас нет сдачи.


*/

func main() {
	fiveDollarBill := 0 // сколько купюр в кассе 5$
	tenDollarBill := 0  // сколько купюр в кассе 10$
	twentyDollarBill := 0

	fmt.Println("Расчёт сдачи.")
	fmt.Print("Ввод: ")
	var queue string
	fmt.Scan(&queue)
	fmt.Println(queue)

	queue = strings.Trim(queue, "[]")
	//fmt.Println("|", queue, "|")

	customers := strings.Split(queue, ",")
	for _, customer := range customers {
		//fmt.Println(customer)
		isSell := false

		switch customer {
		case "5":
			isSell = true
			fiveDollarBill++
		case "10":
			if fiveDollarBill > 0 {
				fiveDollarBill--
				tenDollarBill++
				isSell = true
			}
		case "20":
			if fiveDollarBill > 0 && tenDollarBill > 0 {
				fiveDollarBill--
				tenDollarBill--
				twentyDollarBill++
				isSell = true
			} else if fiveDollarBill >= 3 {
				fiveDollarBill -= 3
				twentyDollarBill++
				isSell = true
			}
		default:
			fmt.Println("Можно использовать только купюры номиналом 5,10 или 20 долларов.")
			return
		}

		if isSell == false {
			fmt.Println("Вывод:", isSell)
			return
		}
	}

	fmt.Println("Вывод: true")
}
