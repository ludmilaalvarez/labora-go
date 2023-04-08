package main

import "fmt"

func main() {
	var segundos int

	fmt.Printf("Ingrese la cantidad de segundos: ")
	fmt.Scan(&segundos)

	calcularPeriodoTiempo(&segundos)

}

func calcularPeriodoTiempo(s *int) {
	var (
		dia      int
		hora     int
		minutos  int
		segundos int
		resto    int
	)

	seg := *s

	if seg > 60 {
		if seg > 3600 {
			if seg > 86400 {
				dia = seg / 86400
				resto = seg % 86400
				hora = resto / 3600
				resto = resto % 3600
				minutos = resto / 60
				resto = resto % 60
				segundos = resto
				fmt.Printf("%d segundos son %d dia, %d hora, %d minutos con %d segundos", seg, dia, hora, minutos, segundos)
			} else {
				hora = seg / 3600
				resto = seg % 3600
				minutos = resto / 60
				resto = resto % 60
				segundos = resto
				fmt.Printf("%d segundos son  %d hora, %d minutos con %d segundos", seg, hora, minutos, segundos)

			}
		} else {
			minutos = seg / 60
			resto = seg % 60
			segundos = resto
			fmt.Printf("%d segundos son %d minutos con %d segundos", seg, minutos, segundos)

		}
	}
}
