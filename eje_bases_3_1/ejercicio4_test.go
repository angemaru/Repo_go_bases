package eje_bases_3_1

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestEst(t *testing.T) {

	//Se inicializan los datos a usar en el test (input/output)
	resultadoEsperado1 := 2.0
	resultadoEsperado2 := 5.0
	resultadoEsperado3 := 1.0

	// Se ejecuta el test
	resultado1 := Ejercicio4("avg", 2, 2, 2)
	resultado2 := Ejercicio4("max", 5, 2, 2)
	resultado3 := Ejercicio4("min", 1, 2, 2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")
	assert.Equal(t, resultadoEsperado2, resultado2, "deben ser iguales")
	assert.Equal(t, resultadoEsperado3, resultado3, "deben ser iguales")

}
