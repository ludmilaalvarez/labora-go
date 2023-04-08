package main

import "fmt"

func main() {
	var num int

	fmt.Printf("Ingrese el valor de x: ")
	fmt.Scan(&num)

	segmentarValorPorRangos(&num)

}

func segmentarValorPorRangos(x *int) {
	num := *x
	if num > 50 {
		num = num - 50
		if num > 50 {
			num = num - 50
			if num > 600 {
				num = num - 600
				if num > 800 {
					num = num - 800
					fmt.Printf("s1=50, s2=50, s3=600, s4=800, s5= %d", num)
				} else {
					fmt.Printf("s1=50, s2=50, s3=600, s4=%d, s5=0", num)
				}

			} else {
				fmt.Printf("s1=50, s2=50, s3=%d, s4=0, s5= 0", num)
			}
		} else {
			fmt.Printf("s1=50, s2=%d, s3=0, s4=0, s5= 0", num)
		}
	} else {
		fmt.Printf("s1=%d, s2=0, s3=0, s4=0, s5= 0", num)
	}

}
