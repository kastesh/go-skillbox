package main

import (
	"fmt"
	"strings"
)

/*
Необходимо сначала выводить пробелы, а затем — звёздочки.
Посмотрите, как количество пробелов и звёздочек в каждой строке связано
с введённым количеством строк и номером текущей строки. Внутри цикла по строкам, скорее всего,
понадобятся два цикла: один — для вывода пробелов, второй — для вывода звёздочек.
*/
func main() {

	fmt.Print("Высота елочки: ")
	var lenght int
	fmt.Scan(&lenght)

	//width := lenght*2 + 1
	space := lenght
	star := 1
	line := 0
	for line < lenght {
		//fmt.Println(star, space, line)
		fmt.Print(strings.Repeat(" ", space),
			strings.Repeat("*", star),
			strings.Repeat(" ", space), "\n")
		star += 2
		space--
		line++
	}
}
