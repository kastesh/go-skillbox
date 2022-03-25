package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 2020-02-10 15:00:00 продам гараж.
	var answer string
	fmt.Println("Введите строку или exit для выхода из программы!")
	file, err := os.Create("log.txt")
	checkErr(err)

	defer file.Close()

	var id int8
	for strings.Compare(answer, "exit") != 0 {
		id++
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Строка: ")
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")

		// Mon Jan 2 15:04:05 MST 2006
		currentTime := time.Now()
		dateString := currentTime.Format("2006-01-02 15:04:05")
		logString := fmt.Sprintf("%03d: %20s: %s\n",
			id, dateString, answer)

		if _, err := file.WriteString(logString); err != nil {
			log.Println(err)
		}
	}
}
