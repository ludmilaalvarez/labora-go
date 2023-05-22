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
	var num Numeros

	fmt.Println("Ingrese el numero del cual desea saber sus divisores: ")
	fmt.Scan(&num.numero)

	num.calcularDivisores()

	fmt.Printf("Los divisores de %d son: %d \n", num.numero, num.numerosDivisores)
	fmt.Printf("La suma total de estos es: %d \n", num.suma)
	
	if num.suma== num.numero{
		fmt.Println("Es un numero perfecto!")
	}else {
		fmt.Println("No es un numero perfecto!")
	}


}

