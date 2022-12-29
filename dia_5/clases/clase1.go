package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		var numero string
		fmt.Scan(&numero)
		fmt.Println(numero)
	*/
	fmt.Println("Inicio")
	procesar()
	fmt.Println("Finalizando")

}

func procesar() {
	fmt.Println("Proceso iniciando")
	time.Sleep(1 * time.Second)
	fmt.Println("Proceso terminado")
}
