package main

import (
	"fmt"
)

func sumar(numeros []int64, canal chan<- int64) {
	//defer wg.Done()

	suma := int64(0)
	for _, num := range numeros {
		suma += num
	}
	canal <- suma
}

func main() {
	var numero int64
	fmt.Print("Ingrese el numero hasta el que desea sumar: ")
	fmt.Scan(&numero)
	datos := make([]int64, numero)
	for i := 0; i < len(datos); i++ {
		datos[i] = int64(i + 1)
	}
	canal := make(chan int64)
	tamanoFragmento := len(datos) / 2

	var suma1, suma2 int64

	go sumar(datos[:tamanoFragmento], canal)
	suma1 = <-canal
	go sumar(datos[tamanoFragmento:], canal)
	suma2 = <-canal

	sumaTotal := suma1 + suma2

	fmt.Printf("La suma total de los numeros del 1 al 100 es: %d", sumaTotal)

}
