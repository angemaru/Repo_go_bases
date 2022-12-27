package dia_3

import "fmt"

func _2main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*fl := i.(float64)
	fmt.Println(fl)*/

}
