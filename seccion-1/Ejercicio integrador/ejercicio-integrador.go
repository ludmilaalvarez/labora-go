package main

import (
	"fmt"
	"sort"
	"strings"
)

type Persona struct {
	nombre string
	edad   int64
	altura int64
	peso   int64
}

func main() {
	menu()
}

func menu() {
	var personas = []Persona{}
	op := 0

	for op != 5 {
		fmt.Printf("Menu:\n")
		fmt.Printf("1) Ingresar personas.\n")
		fmt.Printf("2) Ordenar la lista.\n")
		fmt.Printf("3) Buscar una persona.\n")
		fmt.Printf("4) Ver el IMC\n")
		fmt.Printf("5) Salir del programa.\n")
		fmt.Scan(&op)

		switch op {
		case 1:
			personas = completarDatos()
		case 2:
			personas = parametrosOrden(personas)
		case 3:
			buscarPersona(personas)
		case 4:
			masaCorporal(personas)

		}
	}

}

func completarDatos() []Persona {
	var (
		nombre string
		edad   int64
		altura int64
		peso   int64
	)
	personas := []Persona{}

	for i := 0; i < 5; i++ {

		fmt.Printf("Ingresa el nombre de la persona %d: ", i+1)
		fmt.Scan(&nombre)
		for validarString(nombre) {
			fmt.Printf("Ingresa el nombre de la persona %d: ", i+1)
			fmt.Scan(&nombre)
		}
		fmt.Printf("Ingresa la edad de la persona %d: ", i+1)
		fmt.Scan(&edad)
		for validarNumeros(edad, "edad") {
			fmt.Printf("Ingresa la edad de la persona %d: ", i+1)
			fmt.Scan(&edad)
		}
		fmt.Printf("Ingresa la altura de la persona %d: ", i+1)
		fmt.Scan(&altura)
		for validarNumeros(altura, "altura") {
			fmt.Printf("Ingresa la altura de la persona %d: ", i+1)
			fmt.Scan(&altura)
		}
		fmt.Printf("Ingresa el peso de la persona %d: ", i+1)
		fmt.Scan(&peso)
		for validarNumeros(peso, "peso") {
			fmt.Printf("Ingresa el peso de la persona %d: ", i+1)
			fmt.Scan(&peso)
		}

		personas = append(personas, crearPersona(nombre, edad, altura, peso))
	}
	fmt.Println("Registro de personas desordenado: \n", personas)

	return personas

}

func crearPersona(nombre string, edad int64, altura int64, peso int64) Persona {
	persona := Persona{nombre: nombre, edad: edad, altura: altura, peso: peso}
	return persona

}

func parametrosOrden(personas []Persona) []Persona {
	var opcion string
	personasOrdenadas := []Persona{}
	fmt.Printf("Desea ordenar por nombre, edad, altura o peso?\n")
	fmt.Scan(&opcion)
	for validarString(opcion) {
		fmt.Printf("Desea ordenar por nombre, edad, altura o peso?\n")
		fmt.Scan(&opcion)

	}
	personasOrdenadas = ordenarPersona(personas, opcion)
	fmt.Println(personasOrdenadas)
	return personasOrdenadas

}

func ordenarPersona(personas []Persona, criterio string) []Persona {
	personasOrdenadas := make([]Persona, len(personas))
	copy(personasOrdenadas, personas)

	switch criterio {
	case "nombre":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return strings.ToLower(personasOrdenadas[i].nombre) < strings.ToLower(personasOrdenadas[j].nombre)
		})
	case "edad":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return personasOrdenadas[i].edad < personasOrdenadas[j].edad
		})
	case "altura":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return personasOrdenadas[i].altura < personasOrdenadas[j].altura
		})
	case "peso":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return personasOrdenadas[i].peso < personasOrdenadas[j].peso
		})
	default:
		fmt.Println("Criterio no valido")
	}

	return personasOrdenadas

}

func validarString(nombre string) bool {

	if len(nombre) == 0 {
		fmt.Println("El nombre no puede estar vacio, vuelve a ingresar el valor")
		return true
	} else {
		return false
	}
}

func validarNumeros(numero int64, valor string) bool {
	if (valor == "edad") || (valor == "altura") || (valor == "peso") {
		if (numero < 0) || (numero == 0) {
			fmt.Print("El dato ingresado tiene que ser mayor a 0, vuelva a ingresar el valor\n")
			return true
		} else {
			return false
		}
	}
	return false
}

func buscarPersona(personas []Persona) {
	var nombr string
	var encontrado bool

	fmt.Printf("Ingrese el nombre de la persona que desea buscar: ")
	fmt.Scan(&nombr)

	for i := 0; i < len(personas); i++ {
		if personas[i].nombre == nombr {
			fmt.Printf("Se encontro la persona con Nombre: %s, Edad: %d, Altura: %d, Peso: %d\n", personas[i].nombre, personas[i].edad, personas[i].altura, personas[i].peso)
			encontrado = true
		}
	}

	if !encontrado {
		fmt.Printf("No existe niguna persona con el nombre %s\n", nombr)
	}
}

func masaCorporal(personas []Persona) {
	//	personasIMC:= make([]Persona, len(personas))
	//copy(personas, personas)
	for i := 0; i < 5; i++ {
		alturaTemp := float64(float64(personas[i].altura) / float64(100))
		//		fmt.Println(alturaTemp)
		imc := (float64(personas[i].peso) / (alturaTemp * alturaTemp))
		//		fmt.Println(imc)
		//personasIMC= personas[i].append("IMC" imc)
		if imc < 18.5 {
			fmt.Printf("Persona %d:\n \tNombre: %s, Edad: %d, Altura:%d, Peso:%d, \n\t IMC: %.2f, Bajo peso \n", i+1, personas[i].nombre, personas[i].edad, personas[i].altura, personas[i].peso, imc)
		}
		if (imc > 18.5) && (imc < 24.9) {
			fmt.Printf("Persona %d:\n \tNombre: %s, Edad: %d, Altura:%d, Peso:%d, \n\t IMC: %.2f, Peso normal \n", i+1, personas[i].nombre, personas[i].edad, personas[i].altura, personas[i].peso, imc)
		}
		if (imc > 25) && (imc < 29.9) {
			fmt.Printf("Persona %d:\n \tNombre: %s, Edad: %d, Altura:%d, Peso:%d, \n\t IMC: %.2f, Sobrepeso \n", i+1, personas[i].nombre, personas[i].edad, personas[i].altura, personas[i].peso, imc)
		}
		if imc > 30 {
			fmt.Printf("Persona %d:\n \tNombre: %s, Edad: %d, Altura:%d, Peso:%d, \n\t IMC: %.2f, Obesidad \n", i+1, personas[i].nombre, personas[i].edad, personas[i].altura, personas[i].peso, imc)
		}
	}
}
