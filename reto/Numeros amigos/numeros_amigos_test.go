package main 

import "testing"

type prueba struct{
	arg1 int
	divisores []int
	suma int
}

var numDivisoresTest=[]prueba{
	prueba{6, {3,2,1}, 6},
	prueba{3, {1}, 3},
	prueba{9, {3,1}, 4},
}

func TestcalcularDivisores(t *testing.T){

	for _, test:= range prueba{
		var item= Numeros{test.arg1, test.divisores, test.suma}
		if item.suma != test.suma{
			t.Errorf("Output %f not equal to expected %f", item.suma, test.suma)
		}
	}
}

