package eje_bases_3_1

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestSalario_P3(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	categoria_A := "A"
	minutos_mes_A := 300.0
	resultadoEsperado_A := 22500.0

	// Se inicializan los datos a usar en el test (input/output)
	categoria_B := "B"
	minutos_mes_B := 300.0
	resultadoEsperado_B := 9000.0

	// Se inicializan los datos a usar en el test (input/output)
	categoria_C := "C"
	minutos_mes_C := 300.0
	resultadoEsperado_C := 5000.0

	// Se ejecuta el test
	resultado_A := Get_salario(categoria_A, minutos_mes_A)
	resultado_B := Get_salario(categoria_B, minutos_mes_B)
	resultado_C := Get_salario(categoria_C, minutos_mes_C)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado_A, resultado_A, "deben ser iguales")
	assert.Equal(t, resultadoEsperado_B, resultado_B, "deben ser iguales")
	assert.Equal(t, resultadoEsperado_C, resultado_C, "deben ser iguales")

}
