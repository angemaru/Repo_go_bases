package eje_bases_3_1

func Animal(animal string) (func(cantidad int) float64, string) {
	switch animal {
	case "perro":
		return perro, "El animal ingresado es correcto"
	case "gato":
		return gato, "El animal ingresado es correcto"
	case "hamster":
		return hamster, "El animal ingresado es correcto"
	case "tarantula":
		return tarantula, "El animal ingresado es correcto"
	default:
		return nil, "El animal ingresado es incorrecto"
	}

}

func perro(cantidad int) float64 {
	return float64(cantidad) * 10000
}

func gato(cantidad int) float64 {
	return float64(cantidad) * 5000
}
func hamster(cantidad int) float64 {
	return float64(cantidad) * 250
}
func tarantula(cantidad int) float64 {
	return float64(cantidad) * 150
}
