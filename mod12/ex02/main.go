package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var filename string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Путь к файлу: ")
	filename, _ = reader.ReadString('\n')
	filename = strings.TrimSuffix(filename, "\n")

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 32)
	for {
		n, err := f.Read(buf)
		//fmt.Println(n)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if n > 0 {
			fmt.Printf("%s", buf[:n])
		}
	}
}
