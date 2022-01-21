package main

import (
	"fmt"
	"os"
)

/*
Программное обеспечение банкоматов постоянно решает задачу,
как имеющимися купюрами сформировать сумму, введённую пользователем.
Попробуйте решить похожую задачу и определить, сможет ли пользователь заплатить за товар без сдачи или нет.
Для этого он будет вводить стоимость товара и номиналы трёх монет.
*/

func main() {

	fmt.Print("Стоимость товара: ")
	var price int
	fmt.Scan(&price)

	fmt.Println("Какие моенты есть у Вас в кошельке!")
	fmt.Print(" Первая: ")
	var coin1 int
	fmt.Scan(&coin1)

	fmt.Print(" Вторая: ")
	var coin2 int
	fmt.Scan(&coin2)

	fmt.Print(" Третья: ")
	var coin3 int
	fmt.Scan(&coin3)

	noChange := false

	if price > coin1+coin2+coin3 {
		fmt.Println("В вашем кармане недостаточно момент, чтобы купить это!")
		os.Exit(1)
	}

	if price == coin1+coin2+coin3 {
		noChange = true
	} else if price == coin1+coin2 {
		noChange = true
	} else if price == coin1+coin3 {
		noChange = true
	} else if price == coin3+coin2 {
		noChange = true
	} else if price == coin1 || price == coin2 || price == coin3 {
		noChange = true
	}

	if noChange {
		fmt.Println("Вы можете купить товар без сдачи.")
	} else {
		fmt.Println("Вам дадут сдачу.")
	}

}
