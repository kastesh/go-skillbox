package main

import "fmt"

/*
Напишите программу, которая принимает на вход цену товара и скидку.
Посчитайте и верните на экран сумму скидки.
Скидка должна быть не больше 30% от цены товара и не больше 2000 рублей.
*/
func main() {
	limitDiscountPerc := 30.0
	limitDiscount := 2000.0
	fmt.Print("Введите цену товара в рублях: ")
	var cost float64
	fmt.Scan(&cost)

	fmt.Print("Введите скидку на товар в %: ")
	var discountPerc float64
	fmt.Scan(&discountPerc)

	fmt.Println("---------- Расчет скидки ----------")
	if discountPerc > limitDiscountPerc {
		fmt.Print("Скидка на товар не может превышать ", limitDiscountPerc, "%!\n")
		fmt.Print("Скидка уменьшается до ", limitDiscountPerc, "%!\n")
		discountPerc = limitDiscountPerc
	}

	discount := cost * discountPerc / 100
	if discount > limitDiscount {
		fmt.Printf("Преварительное значение скидки равно %.2f\n", discount)
		fmt.Print("Скидка на товар не может превышать ", limitDiscount, " рублей!\n")
		fmt.Print("Скидка уменьшается до ", limitDiscount, " рублей!\n")
		discount = limitDiscount
	}
	fmt.Println("---------- Итого ----------")
	fmt.Println("Скидка составила:", discount, "рублей.")
}
