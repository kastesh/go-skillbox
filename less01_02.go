package main

import "fmt"

func main() {
	/*
		Калькулятор стоимости товара со скидкой.
		Стоимость товара: 6400
		Стоимость доставки: 350
		Размер скидки: 700
		---------
		Итого: 6050
	*/
	var productCost int  // стоимость продукта
	var deliveryCost int // стомость доставки
	var discount int     // скидка

	fmt.Println("Калькулятор стоимости товара со скидкой.")
	fmt.Println("Введите значения!")
	fmt.Print("Стоимость товара: ")
	fmt.Scan(&productCost)

	fmt.Print("Стоимость доставки: ")
	fmt.Scan(&deliveryCost)

	fmt.Print("Размер скидки: ")
	fmt.Scan(&discount)

	fmt.Println("---------")
	fmt.Println("Итого:",
		productCost+deliveryCost-discount)
}
