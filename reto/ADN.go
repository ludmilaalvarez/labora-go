package main

import (
	"fmt"
	"strings"
)

type ADN struct {
	cadena1 string
	cadena2 string
}

func main() {
	var secuencia ADN

	fmt.Print("Ingrese la primer secuencia de ADN: ")
	fmt.Scan(&secuencia.cadena1)
	fmt.Print("Ingrese la segunda secuencia de ADN: ")
	fmt.Scan(&secuencia.cadena2)
	fmt.Print("Las secuencia de ADN son iguales:")
	fmt.Println(compararADN(secuencia))
}

func compararADN(x ADN) bool {
	resultado := x.cadena2 + x.cadena2
	return strings.Contains(resultado, x.cadena1)
}
