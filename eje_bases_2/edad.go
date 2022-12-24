package eje_bases_2
import (
	"fmt"
	"strconv")

func Get_info_empl (){
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

//Mostrar la edad de Benjamin
fmt.Println(employees["Benjamin"])

//Mostrar la cantidad de empleados mayores de 21
empl_may_21 := 0
for _,age := range employees {
	if age > 21 {
		empl_may_21++
	}
}
fmt.Println("La cantidad de empleados mayores a 21 es: "+strconv.Itoa(empl_may_21))
fmt.Println(employees)

//Agregando a Fede de 25 años
employees["Federico"]=25
fmt.Println(employees)


//Eliminando a Pedro de mapa
delete(employees,"Pedro")
fmt.Println(employees)
}