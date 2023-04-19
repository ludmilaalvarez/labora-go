package main

import (
	"fmt"
	"sort"
	"strings"
)

type Estudiante struct {
	nombre string
	nota   float64
	codigo string
}

func main() {
	menu()
}

func menu() {
	var estudiantes = []Estudiante{}
	op := 0

	for op != 4 {

		fmt.Printf("Menu:\n")
		fmt.Printf("1) Crear estudiantes.\n")
		fmt.Printf("2) Ordenar la lista.\n")
		fmt.Printf("3) Buscar un estudiate.\n")
		fmt.Printf("4) Salir del programa.\n")
		fmt.Scan(&op)
		switch op {
		case 4:
			if len(estudiantes) > 0 {
				defer func() {
					promedio(estudiantes)
					NotaMinMax(estudiantes)
				}()
			}
		case 1:
			estudiantes = append(estudiantes, completarDatos())
		case 2:
			estudiantes = ordenar(estudiantes)
		case 3:
			buscar(estudiantes)
		}

	}
}
func completarDatos() Estudiante {

	var (
		nombre string
		nota   float64
		codigo string
	)

	fmt.Printf("Ingresa el nombre del estudiante: ")
	fmt.Scan(&nombre)

	fmt.Printf("Ingresa la nota del estudiante: ")
	fmt.Scan(&nota)

	fmt.Printf("Ingresa el codigo del estudiante: ")
	fmt.Scan(&codigo)

	return crearEstudiante(nombre, nota, codigo)
}

func crearEstudiante(nombre string, nota float64, codigo string) Estudiante {
	estudiante := Estudiante{nombre: nombre, nota: nota, codigo: codigo}
	return estudiante
}
func ordenar(estudiantes []Estudiante) []Estudiante {
	{
		var opcion string
		var ascendente bool
		estudiantesOrdenados := []Estudiante{}
		fmt.Printf("Desea ordenar por nombre, nota, o codigo?\n")
		fmt.Scan(&opcion)
		fmt.Printf("Desea ordenar en orden ascendente? True o false \n")
		fmt.Scan(&ascendente)
		estudiantesOrdenados = ordenarEstudiantes(estudiantes, opcion, ascendente)
		fmt.Println(estudiantesOrdenados)
		return estudiantesOrdenados
	}

}
func ordenarEstudiantes(estudiantes []Estudiante, criterio string, metodo bool) []Estudiante {
	estudiantesOrdenados := make([]Estudiante, len(estudiantes))
	copy(estudiantesOrdenados, estudiantes)

	comparar := func(a, b bool) bool {
		if metodo {
			return a
		}
		return b
	}

	switch criterio {
	case "nombre":
		sort.Slice(estudiantesOrdenados, func(i, j int) bool {
			return comparar(strings.ToLower(estudiantesOrdenados[i].nombre) < strings.ToLower(estudiantesOrdenados[j].nombre),
				strings.ToLower(estudiantesOrdenados[i].nombre) > strings.ToLower(estudiantesOrdenados[j].nombre))
		})
	case "nota":
		sort.Slice(estudiantesOrdenados, func(i, j int) bool {
			return comparar(estudiantesOrdenados[i].nota < estudiantesOrdenados[j].nota,
				estudiantesOrdenados[i].nota > estudiantesOrdenados[j].nota)
		})
	case "codigo":
		sort.Slice(estudiantesOrdenados, func(i, j int) bool {
			return comparar(strings.ToLower(estudiantesOrdenados[i].codigo) < strings.ToLower(estudiantesOrdenados[j].codigo),
				strings.ToLower(estudiantesOrdenados[i].codigo) > strings.ToLower(estudiantesOrdenados[j].codigo))
		})
	default:
		fmt.Println("Criterio de ordenamiento no válido.")
	}

	return estudiantesOrdenados
}

func buscar(estudiantes []Estudiante) {
	var codigo string
	var encontrado bool

	fmt.Printf("Ingrese el codigo del estudiante que desea buscar: ")
	fmt.Scan(&codigo)

	for i := 0; i < len(estudiantes); i++ {
		if estudiantes[i].codigo == codigo {
			fmt.Printf("Se encontro el estudiante con Nombre: %s, nota: %.2f, codigo: %s\n", estudiantes[i].nombre, estudiantes[i].nota, estudiantes[i].codigo)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Printf("No existe nigun estudiante con el codigo %s\n", codigo)
	}
}

func promedio(estudiantes []Estudiante) {
	var (
		suma     float64
		promedio float64
	)
	for i := 0; i < len(estudiantes); i++ {
		suma = estudiantes[i].nota + suma
	}
	promedio = suma / float64(len(estudiantes))
	fmt.Printf("El promedio de las notas es %.2f \n", promedio)
}

func NotaMinMax(estudiantes []Estudiante) {

	sort.Slice(estudiantes, func(i, j int) bool {
		return estudiantes[i].nota < estudiantes[j].nota
	})

	fmt.Printf("El estudiante con mayor nota es:\n\t Nombre: %s, Nota: %.2f, Codigo: %s\n",
		estudiantes[len(estudiantes)-1].nombre,
		estudiantes[len(estudiantes)-1].nota,
		estudiantes[len(estudiantes)-1].codigo)
	fmt.Printf("El estudiante con menor nota es:\n\t Nombre: %s, Nota: %.2f, Codigo: %s\n\n",
		estudiantes[0].nombre,
		estudiantes[0].nota,
		estudiantes[0].codigo)
}
