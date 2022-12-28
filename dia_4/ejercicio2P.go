package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func _2pmain() {

	file, err := os.Open("customers.txt")

	if err != nil {
		deferF()
	} else {
		b1 := make([]byte, 100)
		_, err := file.Read(b1)
		check(err)
		s := strings.Split(string(b1), ",")
		fmt.Println(s)
		//fmt.Printf("%d bytes: %s\n", n1, string(b1))
	}

}

func check(e error) {
	if e != nil {
		deferF()
	}
}

func deferF() {
	defer func() {
		fmt.Println("Ejecución finalizada")
		err := recover()
		if err != nil {
			log.Println("Ocurrió un error: ", err)
		}
	}()
	panic("el archivo indicado no fue encontrado o está dañado")
}

/*
func _2pmain() {

	file, err := os.Open("customers.txt")

	if err != nil {
		defer func() {
			fmt.Println("Ejecución finalizada")
			err := recover()
			if err != nil {
				log.Println("Ocurrió un error: ", err)
			}
		}()
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())

			//fmt.Println(scanner.Bytes())
		}
	}

}
*/
