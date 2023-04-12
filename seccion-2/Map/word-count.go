package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	var contador map[string]int
	contador = make(map[string]int)
	for _, word := range strings.Fields(str) {
		contador[word]++
	}
	return contador
}

func main() {

	resultado := WordCount("I ate a donut. Then I ate another donut.")
	fmt.Println(resultado)
}
