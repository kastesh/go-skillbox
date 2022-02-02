package main

import (
	"fmt"
	"math"
)

/*
В данном модуле мы рассмотрели примеры по целочисленным типам, их размерам в памяти и то,
что происходит при её переполнении.
Напишите программу, которая в цикле с использованием встроенных констант
(на предельные значения целых чисел, в пакете math) будет подсчитывать,
сколько приходится переполнений чисел типа uint8, uint16 в диапазоне от 0 до uint32.
*/

func main() {

	testInt8 := int8(0)
	testInt16 := int16(0)

	//fmt.Println(math.MaxUint32)

	max := uint32(0)

	overflowInt8 := 0
	for {
		if max == math.MaxUint32 {
			break
		}
		if testInt8 == math.MaxInt8 {
			overflowInt8++
		}
		//fmt.Println(overInt8)
		max++
		testInt8++
	}

	max = 0
	overflowInt16 := 0
	for {
		if max == math.MaxUint32 {
			break
		}
		if testInt16 == math.MaxInt16 {
			overflowInt16++
		}
		//fmt.Println(overInt8)
		max++
		testInt16++
	}

	fmt.Println("В диапазоне от 0 до uint32 тип uint8 переполняется",
		overflowInt8, "раз.")
	fmt.Println("В диапазоне от 0 до uint32 тип uint16 переполняется",
		overflowInt16, "раз.")
}
