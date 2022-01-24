package main

import "fmt"

/*
Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле следующим образом:
берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.
*/

func main() {
	fmt.Print("Введите первое число: ")
	var a int
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	var b int
	fmt.Scan(&b)

	sum := a + b

	cnt := a
	for cnt <= sum {
		fmt.Println(cnt)
		cnt++
	}
}
