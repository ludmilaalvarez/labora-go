package main

import "fmt"

func main() {
	var num int

	fmt.Printf("Ingrese un valor del 1 al 7: ")
	fmt.Scan(&num)

	menu(&num)

}

func menu(dia *int) {
	switch *dia {
	case 1:
		fmt.Printf("Domingo")
	case 2:
		fmt.Printf("Lunes")
	case 3:
		fmt.Printf("Martes")
	case 4:
		fmt.Printf("Miercoles")
	case 5:
		fmt.Printf("Jueves")
	case 6:
		fmt.Printf("Viernes")
	case 7:
		fmt.Printf("Sabado")
	default:
		fmt.Printf("El valor ingresado no es valido!")
	}
}
