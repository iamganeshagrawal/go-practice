package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

type Product struct {
	Id   int64
	Name string
	Cost float64
}

type ProductList []Product
type ProductCallBackFunc func(Product) bool

func (pl ProductList) Include(product Product) bool {
	return pl.IndexOf(product) != -1
}

func (pl ProductList) IndexOf(product Product) int {
	for idx, p := range pl {
		if p == product {
			return idx
		}
	}
	return -1
}

func (pl ProductList) Filter(cb ProductCallBackFunc) ProductList {
	result := ProductList{}
	for _, p := range pl {
		if cb(p) {
			result = append(result, p)
		}
	}
	return result
}

func (pl ProductList) Any(cb ProductCallBackFunc) bool {
	for _, p := range pl {
		if cb(p) {
			return true
		}
	}
	return false
}

func (pl ProductList) All(cb ProductCallBackFunc) bool {
	for _, p := range pl {
		if !cb(p) {
			return false
		}
	}
	return true
}

func (pl ProductList) String() string {
	bb := new(bytes.Buffer)
	// minwidth, tabwidth, padding, padchar, flags
	w := tabwriter.NewWriter(bb, 10, 8, 0, '\t', tabwriter.AlignRight)

	fmt.Fprintf(w, " %s\t %s\t %s\t", "Id", "Name", "Cost")
	fmt.Fprintf(w, "\n %s\t %s\t %s\t", strings.Repeat("-", len("Id")), strings.Repeat("-", len("Name")), strings.Repeat("-", len("Cost")))

	for _, product := range pl {
		fmt.Fprintf(w, "\n %d\t %s\t %.2f\t", product.Id, product.Name, product.Cost)
	}
	w.Flush()
	return bb.String()
}

func main() {
	productList := ProductList{
		{
			Id:   1,
			Name: "Pencil",
			Cost: 5,
		},
		{
			Id:   2,
			Name: "Pen",
			Cost: 10,
		},
		{
			Id:   3,
			Name: "Eraser",
			Cost: 20,
		},
	}
	divider := strings.Repeat("-", 50)

	fmt.Println(productList)
	fmt.Println(divider)
	fmt.Println(productList.Include(Product{1, "Pencil", 5}))
	fmt.Println(divider)
	fmt.Println(productList.Filter(func(p Product) bool {
		return p.Cost > 10
	}))
	fmt.Println(divider)
}
