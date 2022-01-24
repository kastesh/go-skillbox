package main

import "fmt"

/*
В доме — 24 этажа. Лифт должен ходить вверх-вниз, пока не доставит всех пассажиров на первый этаж.
Три пассажира ждут на четвёртом, седьмом и десятом этажах.
При движении вверх лифт не должен останавливаться, при движении вниз — должен собирать всех,
но не более двух человек в лифте. При этом лифт каждый раз доезжает до самого верхнего этажа и
только после этого начинает движение вниз.
Напишите программу, которая доставит всех пассажиров на первый этаж.
*/
func main() {
	liftCapacity := 2

	userWaitFloor04 := 3
	userWaitFloor07 := 3
	userWaitFloor10 := 3

	floor := 11
	for {
		floor--

		if floor == 10 && userWaitFloor10 > 0 {
			fmt.Println("> Остановились на 10м этаже.")

			if liftCapacity > 0 {
				if userWaitFloor10 >= liftCapacity {
					userWaitFloor10 -= liftCapacity
					liftCapacity = 0
				} else {
					liftCapacity -= userWaitFloor10
					userWaitFloor10 = 0
				}
				fmt.Println("++ Забрали на 10ом этаже! Осталось ", userWaitFloor10, "человек.")
				fmt.Println("++ Мест в лифте осталось", liftCapacity)
			} else {
				fmt.Println("-- Места в лифте нет. Извините, подождите")
			}
		}

		if floor == 7 && userWaitFloor07 > 0 {
			fmt.Println("> Остановились на 7м этаже.")

			if liftCapacity > 0 {
				if userWaitFloor07 >= liftCapacity {
					userWaitFloor07 -= liftCapacity
					liftCapacity = 0
				} else {
					liftCapacity -= userWaitFloor07
					userWaitFloor07 = 0
				}
				fmt.Println("++ Забрали на 7ом этаже! Осталось ", userWaitFloor07, "человек.")
				fmt.Println("++ Мест в лифте осталось", liftCapacity)
			} else {
				fmt.Println("-- Места в лифте нет. Извините, подождите")
			}
		}

		if floor == 4 && userWaitFloor04 > 0 {
			fmt.Println("> Остановились на 4м этаже.")

			if liftCapacity > 0 {
				if userWaitFloor04 >= liftCapacity {
					userWaitFloor04 -= liftCapacity
					liftCapacity = 0
				} else {
					liftCapacity -= userWaitFloor04
					userWaitFloor04 = 0
				}
				fmt.Println("++ Забрали на 4ом этаже! Осталось ", userWaitFloor04, "человек.")
				fmt.Println("++ Мест в лифте осталось", liftCapacity)
			} else {
				fmt.Println("-- Места в лифте нет. Извините, подождите")
			}
		}

		if floor == 1 && liftCapacity < 2 {
			fmt.Println(">> Приехали на 1ый этаж. Из лифта вышли", liftCapacity, "пассажиров.")
			liftCapacity = 2
			floor = 20

			if userWaitFloor10 == 0 && userWaitFloor07 == 0 && userWaitFloor04 == 0 {
				fmt.Println(">> Мы всех довезли!")
				break
			}
		}
	}
}
