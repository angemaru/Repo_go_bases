package eje_bases_3_1

func Get_salario(categoria string, minutos_mes float64) float64 {

	numero_horas_mes := minutos_mes / 60
	switch categoria {
	case string('A'):
		return numero_horas_mes*3000.0 + (numero_horas_mes * 3000 * 0.5)
	case string('B'):
		return numero_horas_mes*1500.0 + (numero_horas_mes * 1500 * 0.2)
	case string('C'):
		return numero_horas_mes * 1000.0
	default:
		return 0
	}
}
