package main

import "fmt"

func main() {
	fmt.Print("Введите стоимость товара: ")
	var productCost int
	fmt.Scan(&productCost)

	fmt.Print("Введите стоимость доставки: ")
	var deliveryCost int
	fmt.Scan(&deliveryCost)

	fmt.Print("Введите размер скидки: ")
	var discount int
	fmt.Scan(&discount)

	price := productCost + deliveryCost - discount
	fmt.Println("====================")
	fmt.Println("Итого:", price)
}
