package main

import "fmt"

func main() {
	fmt.Print("Как Вас зовут? ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Вы ввели имя: ", name)
}
