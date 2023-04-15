package main

import (
	"fmt"
	"sort"
)

type Estudiante struct {
	nombre string
	nota   float64
	codigo string
}

type Nombre []Estudiante

func (a Nombre) Len() int           { return len(a) }
func (a Nombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nombre) Less(i, j int) bool { return a[i].nombre < a[j].nombre }

type Nota []Estudiante

func (a Nota) Len() int           { return len(a) }
func (a Nota) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nota) Less(i, j int) bool { return a[i].nota < a[j].nota }

type Codigo []Estudiante

func (a Codigo) Len() int           { return len(a) }
func (a Codigo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Codigo) Less(i, j int) bool { return a[i].codigo < a[j].codigo }

type NombreD []Estudiante

func (a NombreD) Len() int           { return len(a) }
func (a NombreD) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NombreD) Less(i, j int) bool { return a[i].nombre > a[j].nombre }

type NotaD []Estudiante

func (a NotaD) Len() int           { return len(a) }
func (a NotaD) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NotaD) Less(i, j int) bool { return a[i].nota > a[j].nota }

type CodigoD []Estudiante

func (a CodigoD) Len() int           { return len(a) }
func (a CodigoD) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CodigoD) Less(i, j int) bool { return a[i].codigo > a[j].codigo }

func main() {
	//var estudiante []

	menu()

}

func menu() {
	var estudiantes = []Estudiante{}
	op := 0

	//defer func notas(estudiantes)
	for op != 4 {

		fmt.Printf("Menu:\n")
		fmt.Printf("1) Crear estudiantes.\n")
		fmt.Printf("2) Ordenar la lista.\n")
		fmt.Printf("3) Buscar un estudiate.\n")
		fmt.Printf("4) Salir del programa.\n")
		fmt.Scan(&op)
		switch op {
		case 4:
			break
		case 1:
			estudiantes = append(estudiantes, completarDatos())
		case 2:
			estudiantes = ordenar(estudiantes)
		case 3:
			buscar(estudiantes)
		}

	}
	defer func(estudiantes []Estudiante) {
		var (
			suma     float64
			promedio float64
		)
		//		fmt.Println(len(estudiantes))
		for i := 0; i < len(estudiantes); i++ {
			suma = estudiantes[i].nota + suma
		}
		//		fmt.Println(suma)
		promedio = suma / float64(len(estudiantes))
		fmt.Printf("El promedio de las notas es %.2f ", promedio)
	}(estudiantes)

	defer func(estudiantes []Estudiante) {
		var (
			maxnota   float64
			minnota   float64
			alumnomax Estudiante
			alumnomin Estudiante
		)
		minnota = estudiantes[0].nota
		alumnomin = estudiantes[0]
		for i := 0; i < len(estudiantes); i++ {
			if estudiantes[i].nota > maxnota {
				maxnota = estudiantes[i].nota
				alumnomax = estudiantes[i]
			}
			if estudiantes[i].nota < minnota {
				minnota = estudiantes[i].nota
				alumnomin = estudiantes[i]
			}

		}
		fmt.Printf("El estudiante con mayor nota es:\n\t Nombre: %s, Nota: %.2f, Codigo: %s\n", alumnomax.nombre, alumnomax.nota, alumnomax.codigo)
		fmt.Printf("El estudiante con menor nota es:\n\t Nombre: %s, Nota: %.2f, Codigo: %s\n\n", alumnomin.nombre, alumnomin.nota, alumnomin.codigo)

	}(estudiantes)

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

	fmt.Printf("Ingresa el codigo del estudiante")
	fmt.Scan(&codigo)

	return crearEstudiante(nombre, nota, codigo)
}

func crearEstudiante(nombre string, nota float64, codigo string) Estudiante {
	estudiante := Estudiante{nombre: nombre, nota: nota, codigo: codigo}
	return estudiante
}

func ordenar(estudiantes []Estudiante) []Estudiante {

	var opcion int
	var orden int
	fmt.Printf("Desea ordenar por Nombre=1, Nota=2, Codigo=3? ")
	fmt.Scan(&opcion)
	fmt.Printf("Desea ordenar de forma ascendente=1 o descendente=2")
	fmt.Scan(&orden)

	if orden == 1 {
		if opcion == 1 {

			sort.Sort(Nombre(estudiantes))
			fmt.Println(estudiantes)
			return estudiantes

		} else if opcion == 2 {

			sort.Sort(Nota(estudiantes))
			fmt.Println(estudiantes)
			return estudiantes

		} else if opcion == 3 {

			sort.Sort(Codigo(estudiantes))
			fmt.Println(estudiantes)
			return estudiantes

		}
	} else if orden == 2 {
		if opcion == 1 {

			sort.Sort(NombreD(estudiantes))
			fmt.Println(estudiantes)
			return estudiantes

		} else if opcion == 2 {

			sort.Sort(NotaD(estudiantes))
			fmt.Println(estudiantes)
			return estudiantes

		} else if opcion == 3 {

			sort.Sort(CodigoD(estudiantes))
			fmt.Println(estudiantes)
			return estudiantes

		}
	}
	return []Estudiante{}
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

//func promedio(estudiantes []Estudiante)  {

//}
