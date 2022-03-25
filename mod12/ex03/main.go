package main

import (
	"fmt"
	"os"
)

func checkErr(e error, msg string) {
	if e != nil {
		fmt.Println(msg, e)
		os.Exit(255)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func fileinfo(filename string) bool {
	fi, err := os.Stat(filename)
	checkErr(err, "Невозможно получить информацию по файлу ")
	fmt.Printf("Права доступа к файлу %v\n", fi.Mode())
	return true
}

func openFileAndChmod(filename string) (fp *os.File) {
	fp, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		if os.IsPermission(err) {
			err = os.Chmod(filename, 0644)
			checkErr(err, "Невозможно сменить права доступа на файл ")

			fp2, err := os.OpenFile(filename,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			checkErr(err, "Невозможно открыть файл ")
			return fp2

		}
	}

	return fp

}

func main() {
	// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
	// Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.

	//
	filename := "text.log"
	file := openFileAndChmod(filename)
	defer file.Close()

	// отобразить права доступа к файлу
	fmt.Println("Открытие файла", filename)
	_ = fileinfo(filename)

	// сменитьь права на чтение только
	var err error
	err = os.Chmod(filename, 0400)
	checkErr(err, "Невозможно сменить права на файл")

	// отобразить права доступа к файлу
	fmt.Println("Сменили права на файл", filename)
	_ = fileinfo(filename)

	// попытаемся запись в файл с readonly
	// смешно так как открытие было после смены прав, то смена в рамках текущего процесса не работает
	// надо переоткрывать
	/*
		fmt.Println("Пытаемся записать строку в файл", filename)
		if _, err = file.WriteString("test\n"); err != nil {
			checkErr(err, "Невозможно записать строку в файл ")
		}
	*/

	// открываем заново
	fileR, errR := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(errR, "Невозможно открыть файл")
	if _, err = fileR.WriteString("test\n"); err != nil {
		checkErr(err, "Невозможно записать строку в файл ")
	}
	defer fileR.Close()

}
