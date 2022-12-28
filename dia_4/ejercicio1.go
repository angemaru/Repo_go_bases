package main

import "fmt"

func _1main() {
	salary := 3000
	if salary < 150000 {
		e := &myError{
			msg: "Error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(e)
		return
	}
	fmt.Println("Debe pagar el impuesto")

}

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Ha ocurrido un error: %s", e.msg)

}
