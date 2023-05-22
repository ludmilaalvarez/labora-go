package main

import (
	"fmt"
)

func main() {
	var palabra string
	fmt.Printf("Ingrese la palabra que desea inverir:")
	fmt.Scan(&palabra)

	nuevapalabra := invertir(palabra)
	fmt.Println(nuevapalabra)

}

func invertir(palabra string) string {
	var b string
	for _, letra := range palabra {
		b = string(letra) + b
	}
	return b
}
