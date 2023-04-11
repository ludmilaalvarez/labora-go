package main

import "fmt"

func main() {
	var sangre string

	fmt.Printf("Ingrese el tipo de sangre ")
	fmt.Scan(&sangre)

	tipodesangre := clasificarSangre(sangre)
	fmt.Println(tipodesangre)

}

func clasificarSangre(tipoSangre string) string {
	switch tipoSangre {
	case "A+":
		return "Grupo sanguineo A, factor Rh positivo"
	case "A-":
		return "Grupo sanguineo A, factor Rh negativo"
	case "B+":
		return "Grupo sanguineo B, factor Rh positivo"
	case "B-":
		return "Grupo sanguineo B, factor Rh negativo"
	case "AB+":
		return "Grupo sanguineo AB, factor Rh positivo"
	case "AB-":
		return "Grupo sanguineo AB, factor Rh negativo"
	case "0+":
		return "Grupo sanguineo 0, factor Rh positivo"
	case "0-":
		return "Grupo sanguineo 0, factor Rh negativo"
	default:
		return "El valor ingresado no es valido!"
	}

}
