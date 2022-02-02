package main

import (
	"fmt"
	"math"
)

/*
Что нужно сделать
Достаточно часто при передаче по Сети или сохранении больших объёмов данных приходится выбирать
тип с минимальным размером памяти, чтобы экономить трафик или место на диске.
Напишите программу, в которую пользователь вводит два числа (int16), а программа выводит,
в какой минимальный тип данных можно сохранить результат умножения этих чисел.

Советы и рекомендации
Обратите внимание, что положительный результат можно сохранить в меньшем типе
за счёт использования uint8, uint16.
Чтобы не возникло проблем с переполнением в процессе умножения,
числа считывайте в int16, а перед умножением переведите их в int32.

*/

func main() {
	fmt.Print("Введите первое число: ")
	var a int16
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	var b int16
	fmt.Scan(&b)

	c := int32(a) * int32(b)
	result := ""

	if c >= 0 {
		result = testPostiveInt(c)
	} else {
		result = testNegativInt(c)
	}

	fmt.Printf("Результат умножений %d*%d можно сохраниить в тип %s.\n",
		a,
		b,
		result)
}

func testNegativInt(c int32) string {
	switch {
	case c >= int32(math.MinInt8):
		return "int8"
	case c >= int32(math.MinInt16):
		return "int16"
	default:
		return "int32"
	}
}

func testPostiveInt(c int32) string {
	switch {
	case c <= int32(math.MaxUint8):
		return "uint8"
	case c <= int32(math.MaxUint16):
		return "uint16"
	default:
		return "int32"

	}
}
