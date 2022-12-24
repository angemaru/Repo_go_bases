package eje_bases_2
import ("strconv"
"fmt")

func LetrasApp (letras string){
	fmt.Println("La cantidad de letras es: "+strconv.Itoa(len(letras)))
	fmt.Println("Las letras del array son: ")
	for _,letra := range letras{ //Acá usas teoría de runas
		fmt.Printf("%c %v \n",letra,letra) //Acá te lo devuelve como ascii
	}

	//Así tambien funciona:
	/*for _, letra := range letras {
		fmt.Println(string(letra))
	}*/


}

