package main

import (
	"fmt"
)

func _1main() {
	readFile()
}

func readFile() {

	/*
		//Create file
		f, err := os.Create("nameFile.txt")
		if err != nil {
			log.Fatal(err)
		}

		//Write text inside the file
		_, err2 := f.WriteString("text that I want to write inside my file")
		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Print("Done!")*/

	createFile()

}

func createFile() {

	/*

		file := "./nameFile.txt"
		f, err := os.Stat(file) //retorna el nombre del archivo si existe
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Es un directorio: ", f.IsDir())
		fmt.Println("Nombre del archivo/directorio: ", f.Name())
		fmt.Println("Tamaño del archivo en Bytes: ", f.Size())
		fmt.Println("Fecha y hora del archivo: ", f.ModTime())
		fmt.Println("Permisos del archivo:", f.Mode())
	*/

	fmt.Println("El programa está comenzando")

	//os.Exit(0)

	fmt.Println("El programa está terminando")

}
