package main

import (
	"errors"
	"fmt"
)

var _3ErrorSalarioMenor = errors.New("Error: el salario es menor a 10.000")

func _3main() {
	salary := 2000
	var err error

	if salary < 10000 {
		err = _3salaryMenor()
	} else {
		err = _3salaryMayor()
	}
	coincidence := errors.Is(err, _3ErrorSalarioMenor)
	fmt.Println(coincidence)

	if coincidence {
		fmt.Println(err)
	}

	fmt.Println("Debe pagar el impuesto")

}

func _3salaryMenor() error {
	return _3ErrorSalarioMenor
}

func _3salaryMayor() error {
	return errors.New("Error: el salario es mayor a 10.000")
}
