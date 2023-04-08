package main

import "fmt"

func main() {
	var (
		primerEquipo      = "Boca"
		segundoEquipo     = "River"
		tercerEquipo      = "Velez"
		cuartoEquipo      = "Flamengo"
		quintoEquipo      = "Real Madrid"
		titulosBoca       = 74
		titulosRiver      = 69
		titulosVelez      = 16
		titulosFlamengo   = 21
		titulosRealMadrid = 97
	)

	fmt.Printf("El equipo %s ha ganado %d \n", primerEquipo, titulosBoca)
	fmt.Printf("El equipo %s ha ganado %d \n", segundoEquipo, titulosRiver)
	fmt.Printf("El equipo %s ha ganado %d \n", tercerEquipo, titulosVelez)
	fmt.Printf("El equipo %s ha ganado %d \n", cuartoEquipo, titulosFlamengo)
	fmt.Printf("El equipo %s ha ganado %d", quintoEquipo, titulosRealMadrid)
}
