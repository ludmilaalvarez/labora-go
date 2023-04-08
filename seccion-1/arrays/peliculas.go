package main

import "fmt"

func main() {
	peliculas := [10]string{"La era del hielo", "Shrek", "Arma mortal", "El padrino", "Lluvia de hambuguesa", "El gato con botas", "Tiempo violento", "Fragmentado", "Cenicienta", "La sirenita"}

	fmt.Println("Peliculas:")
	for x := 0; x < len(peliculas); x++ {
		fmt.Println(peliculas[x])
	}

	fmt.Println("Segunda Pelicula de la matriz: ", peliculas[1])
	fmt.Println("La longitud de la matriz: ", len(peliculas))

}
