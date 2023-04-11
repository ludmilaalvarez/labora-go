package main

import (
	"fmt"
)

type Persona struct {
	nombre   string
	edad     int64
	ciudad   string
	telefono int64
}

func main() {

	Juan := Persona{nombre: "Juan", edad: 30, ciudad: "Madrid", telefono: 5551234}
	Ana := Persona{nombre: "Ana", edad: 25, ciudad: "Barcelona", telefono: 5555678}

	fmt.Println("Persona 1: ", Juan)
	fmt.Println("Persona 2: ", Ana)

	cambiarCiudad(&Juan, "Cabo Frio")

	fmt.Println("Persona 1 con la ciudad actualizada: ", Juan)
	fmt.Println("Persona 2: ", Ana)

}

func cambiarCiudad(persona *Persona, ciudad string) {

	if persona.nombre == "Juan" && persona.ciudad != ciudad {
		persona.ciudad = ciudad
	}

}
