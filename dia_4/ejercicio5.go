package main

import (
	"errors"
	"fmt"
)

func _5main() {
	fmt.Println(monthlySalary(10, 2000))
}

func monthlySalary(horasMes float64, valorHora float64) (salarioCalculado float64, err error) {

	salarioCalculado = horasMes * valorHora
	if horasMes < 80 {
		return 0, errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if salarioCalculado >= 150000 {
		return (salarioCalculado - (salarioCalculado * 0.1)), nil
	}
	return salarioCalculado, nil

}
