package main

import "fmt"

func main() {
	limit := 5

	fmt.Print("Введите первое число: ")
	var a1 int
	fmt.Scan(&a1)

	fmt.Print("Введите второе число: ")
	var a2 int
	fmt.Scan(&a2)

	fmt.Print("Введите третье число: ")
	var a3 int
	fmt.Scan(&a3)

	numberCount := 0
	if a1 >= limit {
		numberCount++
	}
	if a2 >= limit {
		numberCount++
	}
	if a3 >= limit {
		numberCount++
	}

	if numberCount > 0 {
		fmt.Println("Среди введенных чисел",
			numberCount, "больше или равны", limit)
	}
}
