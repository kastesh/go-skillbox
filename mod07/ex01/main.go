package main

import (
	"fmt"
	"math"
)

/*
Выведите, сколько зеркальных билетов находится среди всех шестизначных номеров от 100000 до 999999.
*/
func main() {

	/*  Размышления у камина :)
	Нам нужно от 100000 до 999999
	соответственно нам нужны пары вида 123 321 внутри
	Соответственно, цикл делаем по половине, формируем пары и считаем.
	Учитываем, что 230 032 в случае int - второе дает нам 32 и не проходит соответствие
	И по факту получается что надо делать обход просто от 100 до 999 и посчитать сколько это
	Итого, найти "середину" числа минимального и максимального.
	От него посчитать сколько чисел между,
	так как эти серединные числа + обратное число даст нам нужное в заданном диапазоне
	*/

	minNum := 100000
	maxNum := 999999
	// нужно найти середину
	digitsNumberInMin := 0
	digitsNumberInMax := 0

	// ищем разрядность для min числа
	saveNum := minNum
	for {
		saveNum = saveNum / 10
		digitsNumberInMin++
		if saveNum == 0 {
			break
		}
	}

	// ищем разрядность максимального числа
	saveNum = maxNum
	for {
		saveNum = saveNum / 10
		digitsNumberInMax++
		if saveNum == 0 {
			break
		}
	}
	fmt.Println("Количество цифр", minNum, "равно", digitsNumberInMin)
	fmt.Println("Количество цифр", maxNum, "равно", digitsNumberInMax)

	digitsSearchNumberMin := digitsNumberInMin / 2
	digitsSearchNumberMax := digitsNumberInMax / 2

	// тут в случае явно заданных чисел можно оставить как есть
	// если цифры произвольные, то надо еще правильно выбрать диапазон для min/max
	min := minNum / int(math.Pow(float64(10), float64(digitsSearchNumberMin)))
	max := maxNum / int(math.Pow(float64(10), float64(digitsSearchNumberMax)))
	fmt.Println("Составлять половинки мы будем из чисел из диапазона от",
		min, "до", max)

	mirrorCnt := max - min + 1

	fmt.Println("В диапазоне от",
		minNum, "до", maxNum, "найдено", mirrorCnt,
		"зарекальных билетов.")

	/*
		это нам нужно если хотим напечатать все зеркальные цифры
		//fmt.Println(min, max)
		mirrorCnt = 0
		for i := min; i <= max; i++ {

			revertI := ""
			processI := i
			base := 10
			baseRevert := min

			for {
				mod10 := processI % base
				whole10 := processI / base
				processI = whole10
				revertI += strconv.Itoa(mod10)

				baseRevert = baseRevert / base

				if whole10 == 0 {
					break
				}
			}

			mirror := strconv.Itoa(i) + revertI
			fmt.Println(mirror)
			mirrorCnt++

		}
		fmt.Println("В диапазоне от",
			min, "до", max, "найдено", mirrorCnt,
			"зарекальных билетов.")

	*/
}
