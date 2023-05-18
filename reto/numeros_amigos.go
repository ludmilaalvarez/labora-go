package main 

import "fmt"

type Numeros struct{
	numero int
	numerosDivisores []int
	suma int
}

func(num *Numeros) calcularDivisores(){
	for i:=1; i<(*num).numero; i++{
		if ((*num).numero%i)==0{
			(*num).numerosDivisores= append((*num).numerosDivisores, i)
			(*num).suma+=i
		}
	}
}


func main(){
	var numUno Numeros
	var numDos Numeros

	fmt.Println("Ingrese el numero primer numero: ")
	fmt.Scan(&numUno.numero)
	fmt.Println("Ingrese el numero segundo numero: ")
	fmt.Scan(&numDos.numero)

	numUno.calcularDivisores()
	numDos.calcularDivisores()


	fmt.Printf("Los divisores de %d son: %d \n", numUno.numero, numUno.numerosDivisores)
	fmt.Printf("Los divisores de %d son: %d \n", numDos.numero, numDos.numerosDivisores)
	
	
	if (numUno.numero== numDos.suma) && (numDos.numero == numUno.suma){
		fmt.Println("Los numeros son amigos!")
	}else {
		fmt.Println("No son amigos")
	}


}

