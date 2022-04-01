package main

import (
	"fmt"
	"math/rand"
)

/*
Напишите программу, которая с помощью функции генерирует
три случайные точки в двумерном пространстве (две координаты),
а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.
*/

func generatePoints(base int) (x, y int) {
	x = rand.Intn(base)
	y = rand.Intn(base)
	return
}

func transformPoints(x *int, y *int) {
	*x = *x*2 + 10
	*y = *y*-3 - 5
}

func strange(base int) {
	x1, y1 := generatePoints(base)
	fmt.Printf("Initial state:\nx1=%d y1=%d\n", x1, y1)
	transformPoints(&x1, &y1)
	fmt.Printf("Transforms to:\nx1=%d y1=%d\n", x1, y1)
}
func printHeader(id int) {
	fmt.Printf("======= %02d generation ======\n", id)
}

func printFooter() {
	fmt.Printf("==============================\n")
}

func main() {

	for id := 1; id <= 3; id++ {
		printHeader(id)
		fmt.Print("Input base int: ")
		var answer int
		fmt.Scan(&answer)

		strange(answer)
		printFooter()
	}
}
