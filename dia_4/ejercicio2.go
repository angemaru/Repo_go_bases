package main

import (
	"errors"
	"fmt"
)

var errorSalarioMenor = errors.New("Error: el salario es menor a 10.000")

func _2main() {
	salary := 30000
	err := &myError1{salary}
	coincidence := errors.Is(err, errorSalarioMenor)
	fmt.Println(coincidence)

	if coincidence {
		fmt.Println(err)
	}

	fmt.Println("Debe pagar el impuesto")

}

type myError1 struct {
	value int
}

func (e *myError1) Error() string {
	if e.value < 10000 {
		return "Error: el salario es menor a 10.000"
	}
	return "Error: el salario es mayor a 10.000"

}
