package eje_bases_3_1

func Promedio(notas ...float64) float64 {
	var suma_notas float64
	for _, nota := range notas {
		if nota >= 0 {
			suma_notas += nota
		}
	}
	return suma_notas / float64(len(notas))
}
