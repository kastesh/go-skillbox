package main

import "fmt"

func main() {
	fmt.Println("Обмен местами значений переменных")
	a := 42
	b := 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// не работает со строками, болььше математических операций и по идее больше использованиия CPU
	a = a + b

	b = a - b
	a = a - b

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
