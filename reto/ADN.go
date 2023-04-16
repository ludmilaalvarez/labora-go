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

	fmt.Println(compararADN(secuencia))
}

func compararADN(x ADN) bool {
	//	s := []string{x.cadena2, x.cadena2}
	//	resultado := strings.Join(s, "")
	resultado := x.cadena2 + x.cadena2
	return strings.Contains(resultado, x.cadena1)
}
