package main

import (
	"fmt"
)

func main() {
	min := 100000
	max := 999999

	distance := 0
	prevLacky := 0
	for i := min; i <= max; i++ {
		left := i / 1000
		right := i % 1000
		var testDistance int

		leftSum := left%10 + (left/10)%10 + left/100
		rightSum := right%10 + (right/10)%10 + right/100
		// пропускаем, если не счастливый
		if leftSum != rightSum {
			continue
		}
		// если есть запомненный предыдущий, считаем
		if prevLacky > 0 {
			testDistance = i - prevLacky
		}
		// если дистанция больше сохраненной, сохраняем
		if testDistance > distance {
			distance = testDistance
		}
		prevLacky = i
		//fmt.Println(i)
	}

	fmt.Println("Наибольшая дистанция", distance)
}
