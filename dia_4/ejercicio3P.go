package main

import (
	"errors"
	"fmt"
	"log"
)

var cli1 = Cliente{[]string{"archivo1", "archivo2", "archivo3"}, "Julián", 1000202, 3123993943, "Cra 1 e No 3-22 Casa 12"}
var cli2 = Cliente{[]string{"archivo12", "archivo22", "archivo32"}, "Ana", 1000202, 3123993943, "Cra 1 e No 3-22 Casa 12"}
var cli3 = Cliente{[]string{"archivo13", "archivo23", "archivo33"}, "María", 1000202, 3123993943, "Cra 1 e No 3-22 Casa 12"}
var Clientes = []Cliente{cli1, cli2, cli3}

func main() {
	var cli_a_registrar = Cliente{[]string{"Doc1.txt"}, "María", 1000202, 3123993943, "Cra 1 e No 3-22 Casa 12"}

	//tarea 1
	err1 := clienteRepetido(100020)
	if err1 != nil {
		deferFunction(err1)
	} else {
		fmt.Println("inserción correcta")
	}

	//tarea 2
	correcta, err2 := clienteSinCeros(cli_a_registrar)
	if err2 != nil {
		deferFunction(err2)
	} else {
		fmt.Println(correcta)
	}

	//tarea 3
	if err1 != nil && err2 != nil {
		deferFunction(errors.New("se detectaron varios errores en tiempo de ejecución"))
		fmt.Println(errors.New("fin de la ejecución"))
	}

}

type Cliente struct {
	Legajo    []string
	Nombre    string
	DNI       int
	Telefono  int
	Domicilio string
}

func clienteSinCeros(cliente Cliente) (bool, error) {
	if len(cliente.Legajo) < 1 || cliente.Nombre == "" || cliente.DNI == 0 || cliente.Telefono == 0 || cliente.Domicilio == "" {
		return false, errors.New("Error: algún campo está vacío")
	}
	return true, nil
}

func clienteRepetido(DNI int) error {
	for _, cliente := range Clientes {
		if cliente.DNI == DNI {
			return errors.New("Error: el cliente ya existe")

		}
	}
	return nil
}

func deferFunction(equivocacion error) {
	defer func() {
		fmt.Println("Ejecución finalizada")
		err := recover()
		if err != nil {
			log.Println("Ocurrió un error: ", err)
		}
	}()
	panic(equivocacion)
}
