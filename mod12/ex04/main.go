package main

import (
	"bufio"
	"fmt"
	"io"
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
	filename := "log.txt"
	fmt.Println("Введите строку или exit для выхода из программы!")

	file, e := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(e)
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

		if _, err := io.WriteString(file, logString); err != nil {
			checkErr(err)
		}
	}

	file1, e1 := os.Open(filename)
	checkErr(e1)
	defer file1.Close()

	stat, e2 := file1.Stat()
	checkErr(e2)
	buf := make([]byte, stat.Size())

	if _, e3 := io.ReadFull(file1, buf); e3 != nil {
		checkErr(e3)
	}

	fmt.Printf("%s", buf)
}
