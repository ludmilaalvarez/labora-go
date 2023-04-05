package main

import (
	"fmt"
)

func main() {
	var (
		primerPlaneta  = "Mercurio"
		segundoPlaneta = "Venus"
		tercerPlaneta  = "Tierra"
		cuartoPlaneta  = "Marte"
		quintoPlaneta  = "Jupiter"
		sextoPlaneta   = "Saturno"
		septimoPlaneta = "Urano"
		octavoPlaneta  = "Neptuno"
		lunaMercurio   = "Ninguna"
		lunaVenus      = "Ninguna"
		lunaTierra     = 1
		lunaMarte      = 2
		lunaJupiter    = 63
		lunaSaturno    = 62
		lunaUrano      = 27
		lunaNeptuno    = 13
	)

	fmt.Printf("%s tiene %s lunas \n", primerPlaneta, lunaMercurio)
	fmt.Printf("%s tiene %s lunas \n", segundoPlaneta, lunaVenus)
	fmt.Printf("%s tiene %d lunas \n", tercerPlaneta, lunaTierra)
	fmt.Printf("%s tiene %d lunas \n", cuartoPlaneta, lunaMarte)
	fmt.Printf("%s tiene %d lunas \n", quintoPlaneta, lunaJupiter)
	fmt.Printf("%s tiene %d lunas \n", sextoPlaneta, lunaSaturno)
	fmt.Printf("%s tiene %d lunas \n", septimoPlaneta, lunaUrano)
	fmt.Printf("%s tiene %d lunas ", octavoPlaneta, lunaNeptuno)
}
