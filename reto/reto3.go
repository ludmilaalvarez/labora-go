package main

import "fmt"

func main() {
	var (
		sectores  int
		capacidad int
		personas  int
	)

	fmt.Printf("Ingrese la cantidad de sectores que desea: ")
	fmt.Scan(&sectores)
	var cantidades []int
	for i := 0; i < sectores; i++ {
		fmt.Printf("Ingrese la capacidad del sector %d: ", i+1)
		fmt.Scan(&capacidad)
		cantidades = append(cantidades, capacidad)
	}

	fmt.Printf("Ingrese la cantidad de personas que van a ingresar: ")
	fmt.Scan(&personas)

	distribucionTotal := distribucion(cantidades, personas)

	for i := 0; i < len(cantidades); i++ {
		fmt.Printf("En el sector %d hay un total de %d personas\n", i+1, distribucionTotal[i])
	}
}

func distribucion(cantidades []int, personas int) []int {
	var total []int

	for i := 0; i < len(cantidades); i++ {

		if personas > 0 {
			if cantidades[i] > personas {
				total = append(total, personas)
			} else if cantidades[i] <= personas {
				total = append(total, cantidades[i])
				personas -= cantidades[i]
			}
		}
	}
	return total
}
