package main

import "fmt"

/*
Баллы ЕГЭ.
Введите результат первого экзамена:
91
Введите результат второго экзамена:
36
Введите результат третьего экзамена:
56
Сумма проходных баллов: 275
Количество набранных баллов: 183
Вы не поступили.
*/
func main() {
	passingPoints := 275
	fmt.Println("Баллы ЕГЭ.")
	fmt.Print("Введите результат первого экзамена: ")
	var points01 int
	fmt.Scan(&points01)

	fmt.Print("Введите результат второго экзамена: ")
	var points02 int
	fmt.Scan(&points02)

	fmt.Print("Введите результат третьего экзамена: ")
	var points03 int
	fmt.Scan(&points03)

	studentPoints := points01 + points02 + points03
	fmt.Println("Количество набранных баллов:", studentPoints)
	if studentPoints >= passingPoints {
		fmt.Println("Вы зачислены!")
	} else {
		fmt.Println("Вы не поступили!")
	}

}
