package eje_bases_3_1

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestPromedio(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	//var notas = []float64{3, 4.5, 2.3, 5, 4.8}

	resultadoEsperado := 3.50

	// Se ejecuta el test
	resultado := Promedio(2, 3, 4, 5)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
