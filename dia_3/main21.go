// package main
package dia_3

import "fmt"

func _main() {
	alumno := Alumno{"Ana", "Santana", 20, "02/04/2020"}
	alumno.detalle()

}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (alum Alumno) detalle() {
	fmt.Println("Nombre:" + alum.Nombre + " \nApellido: " + alum.Apellido + " \nDNI: " + fmt.Sprint(alum.DNI) + " \nFecha: " + alum.Fecha)
}
