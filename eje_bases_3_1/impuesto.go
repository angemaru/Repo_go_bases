package eje_bases_3_1

import "fmt"

func Get_impuesto(salario float64) float64 {
	if salario > 150000 {
		fmt.Println(salario, "is negative")
		return salario * 0.27
	} else if salario > 50000 {
		fmt.Println(salario, "is middle")
		return salario * 0.17
	} else {
		fmt.Println(salario, "is positive")
		return 0.0
	}

}
