package dia_3

import "fmt"

type Producto interface {
	Costo() float64
}
type Grande struct {
	Precio float64
}
type Pequeño struct {
	Precio float64
}
type Mediano struct {
	Precio float64
}

func (m Mediano) Costo() float64 {
	return m.Precio * 1.03
}

func (m Pequeño) Costo() float64 {
	return m.Precio + 20
}

func (m Grande) Costo() float64 {
	return 2500 + m.Precio*1.06
}
func main() {

	fmt.Println(factory("Grande", 3500).Costo())

}

func factory(prod string, precio float64) Producto {

	switch prod {
	case "Grande":
		return Grande{precio}
	case "Pequeño":
		return Pequeño{precio}
	case "Mediano":
		return Mediano{precio}
	default:
		return nil
	}

}
