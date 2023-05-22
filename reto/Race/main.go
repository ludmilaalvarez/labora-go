package main

import (
	"fmt"
	"sync"
	"time"
)

var Mutex sync.Mutex

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	var saldo int = 0

	wg.Add(2)
	Mutex.Lock()
	
	go func() {
		defer Mutex.Unlock()
		ac := saldo
		time.Sleep(1 * time.Nanosecond) // imagenos que se hace alguna operación larga
		ac = ac + 100
		saldo = ac
		wg.Done()
		
	}()

	Mutex.Lock()
	
	go func() {
		defer Mutex.Unlock()
		ac := saldo
		time.Sleep(1 * time.Nanosecond) // imagenos que se hace alguna operación larga
		ac = ac + 100
		saldo = ac
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("saldo final", saldo)
}