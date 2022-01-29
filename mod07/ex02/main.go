package main

import "fmt"

/*
Одним из видов компьютерного искусства является псевдографика,
когда из символов создаются картины.
Попробуйте вывести два изображения. Запросите у пользователя размер шахматной доски в клетках и
выведите шахматную доску на экран.
Белые поля выведите звёздочкой, а чёрные — пробелом.
*/

func main() {
	fmt.Print("Ширина: ")
	var width int
	fmt.Scan(&width)

	fmt.Print("Высота: ")
	var lenght int
	fmt.Scan(&lenght)

	black := 0
	for i := 0; i < lenght; i++ {
		for j := 0; j < width; j++ {
			char := "*"
			if j%2 == black {
				char = " "
			}
			fmt.Print(char)
		}
		black = (black + 1) % 2
		fmt.Printf("\n")
	}
}
