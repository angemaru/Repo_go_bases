package main

import (
	"fmt"
	"log"
)

func _C2main() {
	fmt.Println("Comienza el programa")
	function()

	fmt.Println("Todo salió bien")

}

func function() {
	var apiSlice []string
	defer func() {

		//fmt.Println("Entré antes que el panic!")
		err := recover()
		if err != nil {
			//fmt.Println("Ocurrió un error")
			log.Println("Ocurrió un error", err)
		}
	}()
	_ = apiSlice[0]
}
