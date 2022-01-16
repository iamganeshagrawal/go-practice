package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println("before swapping ", a, b)
	swap(&a, &b)
	fmt.Println("after swapping ", a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
