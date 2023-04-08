package main

import "fmt"

type Planetas struct {
	nombre []string
	lunas  []int
}

func main() {
	/* 	mercurio := Planetas{"Mercurio", 0}
	   	venus := Planetas{"Venus", 0}
	   	tierra := Planetas{"Tierra", 1}
	   	marte := Planetas{"Marte", 2}
	   	jupiter := Planetas{"Jupiter", 62}
	   	saturno := Planetas{"Saturno", 63}
	   	urano := Planetas{"Urano", 27}
	   	neptuno := Planetas{"Neptuno", 13}
	*/

	nombres := []string{"Mercurio", "Venus", "Tierra", "Marte", "Jupiter", "Saturno", "Urano", "Neptuno"}
	lunas := []int{0, 0, 1, 2, 62, 63, 27, 13}
	planetas := Planetas{nombres, lunas}

	for i := 0; i < 8; i++ {
		if planetas.lunas[i] == 0 {
			fmt.Printf("El planeta %s no tiene lunas\n", planetas.nombre[i])
		} else {
			fmt.Printf("El planeta %s tiene %d lunas\n", planetas.nombre[i], planetas.lunas[i])
		}
	}
}
