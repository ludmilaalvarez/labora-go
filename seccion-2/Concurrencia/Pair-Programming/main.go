package main

import (
	"fmt"
	"sync"
)

func main() {
	var opcion int64

	canalsms := make(chan string, 100)
	canalemail := make(chan string, 100)
	canalpush := make(chan string, 100)

	for {
		fmt.Print("Opciones:\n\t",
			"1.Ingresar mensajes\n\t",
			"2.Procesar mensajes\n\t",
			"3.Salir\n")
		fmt.Scan(&opcion)

		switch opcion {
		case 3:
			return
		case 1:
			ingresarMensaje(canalsms, canalemail, canalpush)
		case 2:
			procesarMensajes(canalsms, canalemail, canalpush)
		default:
			return
		}
	}

}

func procesarMensajes(canalsms, canalemail, canalpush chan string) {
	var tipo int64
	var wg sync.WaitGroup

	continuar := 1

	fmt.Print("Ingrese el tipo de Mensaje que desea procesar:\n\t",
		"1.SMS\n\t",
		"2.Email\n\t",
		"3.Push\n\t",
		"4.Todos")
	fmt.Scan(&tipo)

	switch tipo {
	case 1:
		for continuar == 1 {
			wg.Add(1)
			go procesarSMS(canalsms, &wg)
			wg.Wait()
			fmt.Print("\nSi desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	case 2:
		for continuar == 1 {
			wg.Add(1)
			go procesarEmail(canalemail, &wg)
			wg.Wait()
			fmt.Print("Si desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	case 3:
		for continuar == 1 {
			wg.Add(1)
			go procesarPush(canalpush, &wg)
			wg.Wait()
			fmt.Print("Si desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	case 4:
		for continuar == 1 {
			procesarTodos(canalsms, canalemail, canalpush)
			fmt.Print("Si desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	}

}

func procesarSMS(canalsms <-chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	select {
	case mensaje, ok := <-canalsms:
		if ok {
			fmt.Printf("SMS: %s\n", mensaje)
		} else {
			fmt.Println("Canal SMS cerrado!")
		}
	default:
		fmt.Println("El canal de SMS no tiene ningun mensaje")
	}

}

func procesarEmail(canalemail <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case mensaje, ok := <-canalemail:
		if ok {
			fmt.Printf("Email: %s\n", mensaje)
		} else {
			fmt.Println("Canal Email cerrado!")
		}
	default:
		fmt.Println("El canal de Email no tiene ningun mensaje")
	}
}

func procesarPush(canalpush <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case mensaje, ok := <-canalpush:
		if ok {
			fmt.Printf("Push: %s\n", mensaje)
		} else {
			fmt.Println("Canal Push cerrado!")
		}
	default:
		fmt.Println("El canal de Push no tiene ningun mensaje")
	}
}

func procesarTodos(canalsms, canalemail, canalpush chan string) {
	var wg sync.WaitGroup
	wg.Add(3)
	go procesarSMS(canalsms, &wg)
	go procesarEmail(canalemail, &wg)
	go procesarPush(canalpush, &wg)
	wg.Wait()
}

func ingresarMensaje(canalsms chan<- string, canalemail chan<- string, canalpush chan<- string) {

	var (
		mensaje string
		tipo    int64
	)

	fmt.Print("Ingrese el mensaje: ")
	fmt.Scan(&mensaje)
	fmt.Print("Ingrese el tipo de mensaje:\n\t",
		"1.SMS\n\t",
		"2.Email\n\t",
		"3.Push\n")
	fmt.Scan(&tipo)

	switch tipo {
	case 1:
		fmt.Print("\nMensaje sms enviado...\n")
		canalsms <- mensaje

	case 2:
		fmt.Print("\nEmail enviado...\n")
		canalemail <- mensaje
	case 3:
		fmt.Print("\nNotificacion Push enviada...\n")
		canalpush <- mensaje
	}
}
