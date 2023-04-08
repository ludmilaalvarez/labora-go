package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Printf("Ingresa el primer numero entero: ")
	fmt.Scan(&num1)

	fmt.Printf("Ingresa el segundo numero entero: ")
	fmt.Scan(&num2)

	calcular(&num1, &num2)
}

func calcular(n1 *int, n2 *int) {
	suma := *n1 + *n2
	resta := *n1 - *n2
	multiplicacion := *n1 * *n2
	var division float64 = float64(*n1) / float64(*n2)

	fmt.Printf("Suma: %d\n", suma)
	fmt.Printf("Resta: %d\n", resta)
	fmt.Printf("Multiplicacion: %d\n", multiplicacion)
	fmt.Printf("Division: %.2f", division)

}
