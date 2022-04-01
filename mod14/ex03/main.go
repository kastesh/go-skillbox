package main

import (
	"fmt"
	"os"
	"strconv"
)

var progName string

/*
Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его.
Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.
*/

func multipleNumbers(a, b int) (c int) {
	c = a * b
	return
}

func summNumbers(a, b int) (d int) {
	d = a + b
	return
}

func printHelp(exitCode int) {
	fmt.Println("Usage:", progName, "num1", "num2")
	fmt.Println("Exit!")
	os.Exit(exitCode)
}

func parseArgs() (int, int) {
	progName = os.Args[0] // program Name

	// if there is no cmd args
	if len(os.Args[1:]) == 0 {
		fmt.Println("Required options are missing!")
		printHelp(1)
	}
	if os.Args[1] == "help" {
		printHelp(0)
	}

	number1 := os.Args[1]
	number2 := os.Args[2]

	input1, err1 := strconv.Atoi(number1)
	input2, err2 := strconv.Atoi(number2)

	if err1 != nil || err2 != nil {
		printHelp(1)
	}

	return input1, input2
}

func main() {

	i1, i2 := parseArgs()
	fmt.Printf("%d x %d = %d\n", i1, i2, multipleNumbers(i1, i2))
	fmt.Printf("%d + %d = %d\n", i1, i2, summNumbers(i1, i2))

}
