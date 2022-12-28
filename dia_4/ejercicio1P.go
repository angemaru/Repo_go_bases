package main

import (
	"fmt"
	"log"
	"os"
)

func _p1main() {

	_, err := os.Open("customers.txt")
	if err != nil {
		defer func() {
			fmt.Println("Ejecución finalizada")
			err := recover()
			if err != nil {
				log.Println("Ocurrió un error: ", err)
			}
		}()
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}
