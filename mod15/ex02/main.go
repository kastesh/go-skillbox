package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке по сравнению с исходным.
Напишите программу, демонстрирующую работу этого метода.
*/

func reversArr(a [10]int) (b [10]int) {
	for i := 0; i < 10; i++ {
		b[i] = a[9-i]
	}
	return
}

func getInputArray() (input [10]int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Comma separated numbers: ")
	var answer string
	answer, _ = reader.ReadString('\n')
	answer = strings.TrimSuffix(answer, "\n")

	answerS := strings.Split(answer, ",")
	for i, str := range answerS {
		strInt, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Value %v is not integer. Exit!", str)
			os.Exit(255)
		}
		if i < 10 {
			input[i] = strInt
		}
	}

	return
}

func main() {

	input := getInputArray()
	revers := reversArr(input)
	fmt.Println(revers)
}
