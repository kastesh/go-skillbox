package main

import (
	"fmt"
)

/*
Напишите функцию, которая принимает в качестве аргументов два числа типа int,
вычисляет сумму чётных чисел заданного диапазона и выводит результат в консоль.

Если первое число, переданное в качестве аргумента,
будет больше второго, просто поменяйте их местами.
*/

func sumEven(a, b uint64) (sum uint64) {
	sum = 0
	min := a
	max := b
	if a > b {
		min = b
		max = a
	}
	fmt.Println("Summa of even numbers in the range from",
		min, "to", max)

	for i := min; i <= max; i++ {
		if i%2 == 0 {
			//fmt.Println(i)
			sum += i
		}
	}

	return sum
}

func main() {
	var a1, a2 uint64
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a1)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&a2)

	sum := sumEven(a1, a2)
	fmt.Println("Summa: ", sum)

}
