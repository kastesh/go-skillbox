package main

import "fmt"

func main() {
	/*
		Поздравляю, username, теперь вы зарегистрированы!
		Ваш пароль: password
		Ваш возраст: 32
	*/
	var name string
	var password string
	var age int

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)

	fmt.Print("Введите Ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("Поздравляю, ", name, ", теперь вы зарегистрированы!\n")
	fmt.Println("Ваш пароль:", password)
	fmt.Println("Ваш возраст:", age)

}
