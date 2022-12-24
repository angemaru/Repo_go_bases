package eje_bases_2
import "fmt"

func Prestamo(edad int, empleado bool, años int, sueldo float32){

if edad<=22 {
	fmt.Println("No se le puede conceder el préstamo porque no es mayor a 22 años")
}else if empleado == false {
	fmt.Println("No se le puede conceder el préstamo porque no se encuentra empleado")
}else if años == 0 {
	fmt.Println("No se le puede conceder el préstamo porque no lleva más de un año trabajando")
}else{
	if sueldo > 100000{
	fmt.Println("Cumples con todos los requerimientos, su solicitud a préstamo es aceptado. Además, no se te cobrarán intereses")
	}else{
		fmt.Println("Cumples con todos los requerimientos, su solicitud a préstamo es aceptado")
		}
}
	


}