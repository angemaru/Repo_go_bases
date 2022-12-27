package dia_3

import (
	"fmt"
)

func main3() {

	person1 := Person{1, "Juana", "03/02/2003"}
	employe1 := Employe{1, "Tech leader", person1}
	employe1.PrintEmployee()
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employe struct {
	ID       int
	Position string
	Person
}

func (emp Employe) PrintEmployee() {
	fmt.Println("ID: " + fmt.Sprint(emp.ID))

}
