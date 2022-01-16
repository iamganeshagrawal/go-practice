package main

/*
Assignment:3
    Refactor the solution for Assignment:2 using functions
*/

import "fmt"

type OpFunc func(int, int) int

func main() {
	for {
		var choice int = printMenu()

		if choice == 5 {
			break
		}
		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		var op OpFunc = getOp(choice)
		doMagic(op)
	}
}

func doMagic(f OpFunc) {
	a, b := getNumbers()
	result := f(a, b)
	fmt.Println("Result:\t", result)
}

func printMenu() (choice int) {
	fmt.Println("---------Menu---------")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Printf("Enter the choice : ")
	fmt.Scanln(&choice)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func getOp(c int) OpFunc {
	switch c {
	case 1:
		return add
	case 2:
		return subtract
	case 3:
		return multiply
	case 4:
		return divide
	}
	return nil
}

func getNumbers() (a, b int) {
	fmt.Printf("Enter the numbers: ")
	fmt.Scanf("%d %d\n", &a, &b)
	return
}
