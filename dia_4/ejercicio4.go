package main

import (
	"errors"
	"fmt"
)

var (
	_4ErrorSalarioMenor = errors.New("Error: el mínimo posible es de 10.000 ")
	_4ErrorSalarioMayor = errors.New("El valor ingresado es mayor o igual a 10.000")
)

func _4main() {
	salary := 200
	var err error

	if salary < 10000 {
		err = _4salaryMenor(salary)
		//fmt.Println(err)
	} else {
		err = _4salaryMayor(salary)
	}
	coincidence := errors.Is(err, _4ErrorSalarioMenor)
	fmt.Println(coincidence)

	if coincidence {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar el impuesto")
	}

}

func _4salaryMenor(salary int) error {
	return fmt.Errorf("%w, y el salario ingresado es de: %v", _4ErrorSalarioMenor, salary)
}

func _4salaryMayor(salary int) error {
	return fmt.Errorf("%w, y el salario ingresado es de: %v", _4ErrorSalarioMenor, salary)
}
