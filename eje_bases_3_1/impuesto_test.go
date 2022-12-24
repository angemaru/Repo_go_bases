package eje_bases_3_1

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestImpuesto(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	salario_1 := 40000.0
	salario_2 := 60000.0
	salario_3 := 160000.0
	impuesto_esperado_1 := 0.0
	impuesto_esperado_2 := 10200.0
	impuesto_esperado_3 := 43200.0

	// Se ejecuta el test
	resultado_A := Get_impuesto(salario_1)
	resultado_B := Get_impuesto(salario_2)
	resultado_C := Get_impuesto(salario_3)

	// Se validan los resultados
	assert.Equal(t, impuesto_esperado_1, resultado_A, "deben ser iguales")
	assert.Equal(t, impuesto_esperado_2, resultado_B, "deben ser iguales")
	assert.Equal(t, impuesto_esperado_3, resultado_C, "deben ser iguales")

}
