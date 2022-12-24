package eje_bases_2
import "fmt"

func Get_mes(number int){
	 meses := map[int]string{
		1:"Enero",
		2:"Febrero",
		3:"Marzo",
		4:"Abril",
		5:"Mayo",
		6:"Junio",
		7:"Julio",
		8:"Agosto",
		9:"Septiembre",
		10:"Octubre",
		11: "Noviembre",
		12:"Diciembre",
	}
	fmt.Println("El mes seleccionado es: "+meses[number])

}