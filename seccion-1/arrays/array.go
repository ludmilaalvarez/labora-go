package array

import "fmt"

func main() {
	numeros := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nombres := [5]string{"Gala", "Mauro", "Teo", "Roberto", "Jorge"}

	fmt.Println("Numeros:")
	for x := 0; x < len(numeros); x++ {
		fmt.Println(numeros[x])
	}
	fmt.Println("Nombres:")
	for x := 0; x < len(nombres); x++ {
		fmt.Println(nombres[x])
	}
}
