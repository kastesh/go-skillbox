package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("Введите х: ")
	var x float64
	fmt.Scan(&x)

	fmt.Print("С точностью до какого знака нужно ввычислить экспоненту: ")
	var n int
	fmt.Scan(&n)

	epsilon := 1 / math.Pow(10.0, float64(n))

	fact := 1
	k := 0
	result := 0.0
	resultPrev := 0.0
	for {
		if k > 0 {
			fact *= k
		}
		result += math.Pow(x, float64(k)) / float64(fact)

		if math.Abs(result-resultPrev) < epsilon {
			fmt.Println("Экспонента равна: ", result)
			break
		}

		resultPrev = result
		k++

	}
}
