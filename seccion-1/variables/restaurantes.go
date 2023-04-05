package main

import "fmt"

func main() {
	var (
		primerRestaurante   = "Chucrut"
		segundoRestaurante  = "Quesito"
		tercerRestaurante   = "La tabla"
		calificacionChucrut = 2
		calificacionQuesito = 4.9
		calificacionLaTabla = 4.2
	)

	fmt.Printf("El restaurante %s tiene una calificacion de %d \n", primerRestaurante, calificacionChucrut)
	fmt.Printf("El restaurante %s tiene una calificacion de %.2f \n", segundoRestaurante, calificacionQuesito)
	fmt.Printf("El restaurante %s tiene una calificacion de %.2f", tercerRestaurante, calificacionLaTabla)
}
