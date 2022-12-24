package eje_bases_3_1

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func Test_perro(t *testing.T) {

	//Arrange
	animal := "perro"
	cantidad := 5
	resultado1 := 50000.0
	message_expected := "El animal ingresado es correcto"

	//Act
	funcionAnimal, mesage := Animal(animal)
	resultadoEsperado1 := funcionAnimal(cantidad)

	//Assert
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")
	assert.Equal(t, message_expected, mesage, "deben ser iguales")

}

func Test_loro(t *testing.T) {

	//Se inicializan los datos a usar en el test (input/output)
	animal := "loro"
	message_expected := "El animal ingresado es incorrecto"

	// Se ejecuta el test
	funcionAnimal, mesage := Animal(animal)

	// Se validan los resultados
	//assert.Equal(t, funcionAnimal, nil, "deben ser iguales")
	assert.Nil(t, funcionAnimal)
	assert.Equal(t, message_expected, mesage, "deben ser iguales")

}

func TestPerro(t *testing.T) {

	//Se inicializan los datos a usar en el test (input/output)
	cantidad := 5
	resultadoEsperado1 := 50000.0

	// Se ejecuta el test
	resultado1 := perro(cantidad)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestGato(t *testing.T) {

	//Se inicializan los datos a usar en el test (input/output)
	cantidad := 5
	resultadoEsperado1 := 25000.0

	// Se ejecuta el test
	resultado1 := gato(cantidad)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestHamster(t *testing.T) {

	//Se inicializan los datos a usar en el test (input/output)
	cantidad := 5
	resultadoEsperado1 := 1250.0

	// Se ejecuta el test
	resultado1 := hamster(cantidad)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestTarantula(t *testing.T) {

	//Se inicializan los datos a usar en el test (input/output)
	cantidad := 5
	resultadoEsperado1 := 750.0

	// Se ejecuta el test
	resultado1 := tarantula(cantidad)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}
