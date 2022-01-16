package main

import (
	"assignment/04/model"
	"fmt"
)

func main() {
	employee := model.NewEmployee("EMP01", "Ram", 10000.00, 1200.00, 600.00, 5.0)
	fmt.Println(employee)
	fullTimeEmployee := model.NewFullTimeEmployee("EMP02", "Shyam", 10000.00, 1200.00, 600.00, 5.0, []string{"Food Card", "2 Movie Tickets"})
	fmt.Println(fullTimeEmployee)
}
