package model

import (
	"fmt"
	"strings"
)

type FullTimeEmployee struct {
	Employee
	Benefist []string
}

func NewFullTimeEmployee(id string, name string, basic float64, hra float64, da float64, tax float64, benefits []string) *FullTimeEmployee {
	return &FullTimeEmployee{
		Employee: Employee{
			Id:    id,
			Name:  name,
			Basic: basic,
			HRA:   hra,
			DA:    da,
			Tax:   tax,
		},
		Benefist: benefits,
	}
}

func (e *FullTimeEmployee) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Basic: %.2f, HRA: %.2f, DA: %.2f, Tax: %.2f%%, Salary: %.2f, Benefist: [%s]", e.Id, e.Name, e.Basic, e.HRA, e.DA, e.Tax, e.Salary(), strings.Join(e.Benefist, ", "))
}

func (e *FullTimeEmployee) Salary() float64 {
	salary := e.Employee.Salary()
	return salary + (salary * 0.20)
}
