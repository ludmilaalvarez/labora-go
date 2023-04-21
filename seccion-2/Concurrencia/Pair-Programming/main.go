package main

import (
	"fmt"
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
	continuar := 1
	fmt.Print("Ingrese el tipo de Mensaje que desea procesar:\n\t1.SMS\n\t2.Email\n\t3.Push\n\t4.Todos")
	fmt.Scan(&tipo)
	switch tipo {
	case 1:
		for continuar == 1 {
			go procesarSMS(canalsms)
			fmt.Print("\nSi desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	case 2:
		for continuar == 2 {
			go procesarEmail(canalemail)
			fmt.Print("Si desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	case 3:
		for continuar == 1 {
			go procesarPush(canalpush)
			fmt.Print("Si desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	case 4:
		for continuar == 1 {
			go procesarTodos(canalsms, canalemail, canalpush)
			fmt.Print("Si desea continuar procesando mensajes presione 1...")
			fmt.Scan(&continuar)
		}
	}

}

func procesarSMS(canalsms <-chan string) {
	var mensaje string
	mensaje = <-canalsms
	fmt.Printf("SMS: %s\n", mensaje)
}

func procesarEmail(canalemail <-chan string) {
	var mensaje string
	mensaje = <-canalemail
	fmt.Printf("Email: %s\n", mensaje)
}

func procesarPush(canalpush <-chan string) {
	var mensaje string
	mensaje = <-canalpush
	fmt.Printf("Notificacion Push: %s\n", mensaje)
}

func procesarTodos(canalsms, canalemail, canalpush chan string) {
	for i := 0; i < 5; i++ {
		go procesarSMS(canalsms)
		go procesarEmail(canalemail)
		go procesarPush(canalpush)
	}
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
