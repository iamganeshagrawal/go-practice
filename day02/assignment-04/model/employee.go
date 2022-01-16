package model

import "fmt"

type Employee struct {
	Id    string
	Name  string
	Basic float64
	HRA   float64
	DA    float64
	Tax   float64
}

func NewEmployee(id string, name string, basic float64, hra float64, da float64, tax float64) *Employee {
	return &Employee{Id: id, Name: name, Basic: basic, HRA: hra, DA: da, Tax: tax}
}

func (e Employee) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Basic: %.2f, HRA: %.2f, DA: %.2f, Tax: %.2f%%, Salary: %.2f", e.Id, e.Name, e.Basic, e.HRA, e.DA, e.Tax, e.Salary())
}

func (e *Employee) Salary() float64 {
	total := e.Basic + e.HRA + e.DA
	return total - (total * e.Tax / 100)
}
