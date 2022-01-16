package main

import "fmt"

func main() {

	fmt.Print("Введите стоимость первого товара: ")
	var productCost1 int
	fmt.Scan(&productCost1)

	fmt.Print("Введите стоимость второго товара: ")
	var productCost2 int
	fmt.Scan(&productCost2)

	fmt.Print("Введите стоимость третьего товара: ")
	var productCost3 int
	fmt.Scan(&productCost3)

	totalCost := productCost1 + productCost2 + productCost3
	if totalCost >= 10000 {
		fmt.Println("Сумма товаров больше 10т. Вам положена скида 10%")
		totalCost -= totalCost * 10 / 100
	}
	fmt.Println("Необходимо заплатить:", totalCost)
}
