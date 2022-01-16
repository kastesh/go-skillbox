package main

import "fmt"

/*
Банкомат.
Введите сумму снятия со счёта:
500
Операция успешно выполнена.
Вы сняли 500 рублей.
*/
func main() {
	fmt.Println("Банкомат.")
	fmt.Print("Введите сумму снятия со счета: ")
	var sum int
	fmt.Scan(&sum)

	if sum%100 == 0 {
		fmt.Print("Операция успешно выполнена.\n",
			"Вы сняли ", sum, " рублей.\n")
	} else {
		fmt.Print("Невозможно выполнить операцию.\n")
	}

}