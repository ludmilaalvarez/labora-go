package main

import (
	"fmt"
	"sync"
	"time"
)


func sumarSecuencial(numeros []int64) int64{
	suma := int64(0)
	for _, num := range numeros {
		suma += num
	}
	return suma
}

func sumar(numeros []int64, resultado *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	suma := int64(0)
	for _, num := range numeros {
		suma += num
	}
	*resultado = suma
}

func main() {
	datos := make([]int64, 100000)
	for i := 0; i < len(datos); i++ {
		datos[i] = int64(i + 1)
	}

	tamanoFragmento := len(datos) / 2
	inicio := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)
	var suma1, suma2 int64
	go sumar(datos[:tamanoFragmento], &suma1, &wg)
	go sumar(datos[tamanoFragmento:], &suma2, &wg)

	wg.Wait()
	tiempoTranscurrido := time.Since(inicio)

	sumaTotal := suma1 + suma2

	fmt.Printf("La suma total es: %d\n", sumaTotal)

	fmt.Printf("El tiempo transcurrido es: %s\n", tiempoTranscurrido)
	
	inicio = time.Now()

	sumaTotal = sumarSecuencial(datos)

	tiempoTranscurrido = time.Since(inicio)
	fmt.Printf("Suma total secuencial: %d\n", sumaTotal)
	fmt.Printf("El tiempo transcurrido es: %s\n", tiempoTranscurrido)
}
