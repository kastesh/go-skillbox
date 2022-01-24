package main

import "fmt"

/*
Представьте, что у вас есть три корзины разной ёмкости.
Пользователю предлагается ввести, какое количество яблок помещается в каждую корзину.
После этого программа должна заполнить все корзины по очереди, учитывая, какие корзины уже заполнены,
строго соблюдая очерёдность заполнения и добавляя по одному яблоку в каждой итерации.

Используйте if и continue и break.

Пример: пользователь решил, что у корзин будет ёмкость на шесть, четыре и девять яблок.
Программа должна заполнить корзину 1 и в следующей итерации перейти к корзине 2,
далее в следующей итерации перейти к корзине 3. Если очередная корзина уже заполнена,
программа должна переходить к следующей по очереди, и так по кругу, пока не заполнит все.
*/
func main() {
	fmt.Print("Размер первой корзины: ")
	var bucket01 int
	fmt.Scan(&bucket01)

	fmt.Print("Размер второй корзины: ")
	var bucket02 int
	fmt.Scan(&bucket02)

	fmt.Print("Размер третьей корзины: ")
	var bucket03 int
	fmt.Scan(&bucket03)

	// max >= mid >= min
	max := bucket01
	mid := bucket02
	min := bucket03

	// max is max now
	if min > max && min > mid {
		min, max = max, min
	} else if mid > max {
		mid, max = max, mid
	}

	if min > mid {
		mid, min = min, mid
	}

	//fmt.Println(max, mid, min)
	for {
		//fmt.Println(max, mid, min)
		if max == 0 {
			fmt.Println("Все корзины заполнены")
			break
		}
		max--
		if mid <= 0 {
			continue
		}
		mid--

		if min <= 0 {
			continue
		}
		min--
	}
}
