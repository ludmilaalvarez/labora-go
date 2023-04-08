package main

import "fmt"

func main() {
	var a, b = 10, 20
	var ptrA *int

	ptrA = &a

	a, b = b, *ptrA

	fmt.Printf("A= %d, B= %d", a, b)

}
