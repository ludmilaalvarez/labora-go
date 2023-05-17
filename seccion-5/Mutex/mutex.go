package main
 
import  ("fmt"
		"sync")


type NumeroAIncrementar struct{
	Numero int64
	Mutex sync.Mutex
}

func Incrementar(numero *NumeroAIncrementar ,wg *sync.WaitGroup ){
	defer wg.Done()
	(*numero).Mutex.Lock()
	numero.Numero++
	(*numero).Mutex.Unlock()
}


func main(){
	var wg sync.WaitGroup
	numero:= &NumeroAIncrementar{Numero:0}

	for i:=0; i<100; i++{
		wg.Add(2)

		go Incrementar(numero, &wg)
		go Incrementar(numero, &wg)
	}
	wg.Wait()
	fmt.Println("La suma total es: ", numero.Numero)
}