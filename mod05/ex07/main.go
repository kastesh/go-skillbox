package main

import (
	"fmt"
	"os"
)

/*
Напишите программу, в которую пользователь будет вводить четырёхзначный номер билета,
а программа будет выводить, является ли он зеркальным, счастливым или обычным билетом.
*/
func main() {
	fmt.Print("Введите номер билета: ")
	var ticket int
	fmt.Scan(&ticket)

	if ticket > 9999 || ticket < 1000 {
		fmt.Println("Номер билета должен состоять из 4х цифр.")
		os.Exit(1)
	}

	p01 := ticket / 1000
	p02 := (ticket - p01*1000) / 100
	p03 := (ticket - p01*1000 - p02*100) / 10
	p04 := ticket - p01*1000 - p02*100 - p03*10

	//fmt.Println(p01, p02, p03, p04)
	if p01+p02 == p03+p04 {
		fmt.Printf("Билет %d - счастливый.\n", ticket)
	} else {
		fmt.Printf("Билет %d - обычный.\n", ticket)
	}
}
