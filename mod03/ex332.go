package main

import "fmt"

func main() {
	var name string
	var password string

	fmt.Print("Username: ")
	fmt.Scan(&name)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	fmt.Print("Поздравляю, ", name, ", вы успешно зашли!\n")

}
