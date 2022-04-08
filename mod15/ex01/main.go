package main

/*
Разработайте программу, позволяющую ввести 10 целых чисел,
а затем вывести из них количество чётных и нечётных чисел.
Для ввода и подсчёта используйте разные циклы.
*/
import "fmt"

func getEvenAndOdd(arr [10]int) (even int, odd int) {
	for _, a := range arr {
		if a%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return
}

func main() {

	var input [10]int
	for i, _ := range input {
		fmt.Printf("Input %v int: ", i+1)
		fmt.Scan(&input[i])
	}

	even, odd := getEvenAndOdd(input)
	fmt.Println("There are", even, "even numbers.")
	fmt.Println("There are", odd, "odd numbers.")
}
