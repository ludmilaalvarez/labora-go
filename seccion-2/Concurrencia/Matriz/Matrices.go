package main

import (
	"fmt"
	"sync"
)

func main() {
	canal := make(chan []int64)
	var (
		longitud     int
		matrizfinal  [][]int64
		primerVector []int64
		wg           sync.WaitGroup
	)

	fmt.Print("Ingrese longitud de las matrices: ")
	fmt.Scan(&longitud)

	matrizUno := make([][]int64, longitud)
	for i := range matrizUno {
		matrizUno[i] = make([]int64, longitud)
	}
	matrizDos := make([][]int64, longitud)
	for i := range matrizDos {
		matrizDos[i] = make([]int64, longitud)
	}

	fmt.Print("Ingrese los datos de la primera matriz\n")
	matrizUno = llenarMatriz(matrizUno, longitud)
	fmt.Print("Ingrese los datos de la segunda matriz\n")
	matrizDos = llenarMatriz(matrizDos, longitud)
	fmt.Println("\nMatriz uno: ", matrizUno)
	fmt.Println("Matriz dos: ", matrizDos)

	wg.Add(longitud)

	for i := 0; i < longitud; i++ {
		vector := matrizUno[i]
		go multiplicarMatriz(matrizDos, vector, canal, &wg)
		primerVector = <-canal
		matrizfinal = append(matrizfinal, primerVector)
	}

	wg.Wait()
	fmt.Println("Resultado de la multiplicacion de matrices: ", matrizfinal)
}

func llenarMatriz(matriz [][]int64, longitud int) [][]int64 {

	for i := 0; i < longitud; i++ {
		for j := 0; j < longitud; j++ {
			fmt.Printf("Ingrese el valor da la Matriz en la posicion %d:%d: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	return matriz
}

func multiplicarMatriz(matriz [][]int64, vector []int64, canal chan []int64, wg *sync.WaitGroup) {
	var nuevovector []int64

	for i := 0; i < len(vector); i++ {
		var suma int64 = 0
		for j := 0; j < len(vector); j++ {
			suma = (matriz[j][i] * vector[j]) + suma
		}
		nuevovector = append(nuevovector, int64(suma))
	}
	canal <- nuevovector
	defer wg.Done()
}
