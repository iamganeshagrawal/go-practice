package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Product struct {
	Id       int64
	Name     string
	Cost     float64
	Units    int64
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %.2f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discount float64) {
	p.Cost = p.Cost * (1 - discount)
}

type Products []Product

func (p Products) String() string {
	sb := new(strings.Builder)
	for idx, product := range p {
		sb.WriteString(product.String())
		if idx != len(p)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
func (p Products) TablePrint() {
	// sb := new(bytes.Buffer)
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "Id", "Name", "Cost", "Units", "Category")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")

	for _, product := range p {
		fmt.Fprintf(w, "\n %d\t%s\t%.2f\t%d\t%s\t", product.Id, product.Name, product.Cost, product.Units, product.Category)
	}
	// fmt.Printf(sb.String())
}

func main() {
	p := Product{
		Id:       4,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}

	fmt.Println(p)
	p.ApplyDiscount(0.2)
	fmt.Println(p)

	var productList Products
	productList = append(productList,
		Product{
			Id:       1,
			Name:     "Pen",
			Cost:     10,
			Units:    100,
			Category: "Stationary",
		},
		Product{
			Id:       2,
			Name:     "Pencil",
			Cost:     5,
			Units:    1000,
			Category: "Stationary",
		})

	fmt.Println(productList)
	productList.TablePrint()
}
