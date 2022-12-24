package eje_bases_3_1

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Ejercicio4(operacion string, notas ...float64) float64 {

	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	} else {
		println(minFunc)
	}

	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	} else {
		println(averageFunc)
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	} else {
		println(maxFunc)
	}

	switch operacion {
	case "min":
		return minFunc(notas...)
	case "max":
		return maxFunc(notas...)
	case "avg":
		return averageFunc(notas...)
	default:
		return 0
	}

	/*

		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		println(minValue)
		println(averageValue)
		println(maxValue)*/

}

func operation(ope string) (func(values ...float64) float64, error) {
	fmt.Println(ope)
	switch ope {
	case "minimum":
		return opMinVal, nil
	case "average":
		return opAvgVal, nil
	case "maximum":
		return opMaxVal, nil
	default:
		return nil, errors.New("Error")
	}

}

func opMinVal(values ...float64) (value float64) {
	value = values[0]
	for _, number := range values {
		if number < value {
			value = number
		}
	}
	return

}

func opAvgVal(values ...float64) (value float64) {
	sum := 0.0
	for _, number := range values {
		sum += number
	}
	return sum / float64(len(values))
}

func opMaxVal(values ...float64) (value float64) {
	value = values[0]
	for _, number := range values {
		if number > value {
			value = number
		}
	}
	return
}
