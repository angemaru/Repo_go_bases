package main

import (
	"io"
	"log"
	"os"
)

func _v2main() {

	/*
		var s = "Este texto es copiado"

		//Leer s
		r := strings.NewReader(s)

		//copiar r
		b, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	*/

	var s = "Este texto es copado hasta la ultima linea"

	//Vamos a escribir en la terminar con el valor de s
	_, err := io.WriteString(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}
}
