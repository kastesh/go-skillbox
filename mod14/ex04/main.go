package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
Затем вызовите по очереди три функции, передавая результат из одной в другую.
*/

var internalInt1 int
var internalInt2 int
var internalInt3 int
var progName string

func printHelp(exitCode int) {
	fmt.Println("Usage:", progName, "num1")
	fmt.Println("Exit!")
	os.Exit(exitCode)
}

func initVars() {
	internalInt1 = 1
	internalInt2 = 2
	internalInt3 = 3
}

func parseArgs() int {
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

	input1, err1 := strconv.Atoi(number1)

	if err1 != nil {
		printHelp(1)
	}

	return input1
}

func function1(a int) (b int) {
	b = a + internalInt1
	return
}

func function2(a int) (b int) {
	b = a + internalInt2
	return
}

func function3(a int) (b int) {
	b = a + internalInt3
	return
}
func main() {
	initVars()

	i1 := parseArgs()

	res1 := function1(i1)
	res2 := function2(res1)
	res3 := function3(res2)
	fmt.Println("function1:", res1)
	fmt.Println("function2:", res2)
	fmt.Println("function3:", res3)
}
