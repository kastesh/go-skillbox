package main

import (
	"fmt"
	"math"
)

/*
Что компьютеры делают быстрее людей? Выполняют вычисления!

Напишите программу, решающую квадратные уравнения.
Пользователь вводит коэффициенты a, b и c квадратного уравнения. Программа должна вывести корни уравнения.
Для решения уравнения необходимо:

Вычислить дискриминант по формуле: D = (b × b − 4 × a × c)

Если D > 0, будет два различных корня, которые находятся по формуле.
Если D = 0, будет один корень, который находится по формуле.
Если D < 0, то вывести, что корней нет.
*/
func main() {

	fmt.Println("Найдем корни квадратного уравнения: ax2 + bx + c = 0")
	fmt.Print(" Введите коэффициент a: ")
	var a float64
	fmt.Scan(&a)

	fmt.Print(" Введите коэффициент b: ")
	var b float64
	fmt.Scan(&b)

	fmt.Print(" Введите коэффициент c: ")
	var c float64
	fmt.Scan(&c)

	// D = b*b − 4ac
	D := math.Pow(b, 2) - 4*a*c

	fmt.Println("--------------------")
	if D < 0 {
		fmt.Println("Уравнение не имеет решения.")
	} else if D == 0 {
		fmt.Println("Уравнение имеет одно решение.")
		x := -b / (2 * a)
		fmt.Printf("Значение x равно %.2f.\n", x)
	} else {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Уравнение имеет два решения.")
		fmt.Printf("Значение x1 равно %.2f.\n", x1)
		fmt.Printf("Значение x2 равно %.2f.\n", x2)
	}
}
