package main

/*
Assignment:2
    Write an interactive calculator
    Menu
        1. Add
        2. Subtract
        3. Multiply
        4. Divide
        5. Exit

    If the userChoice is 1 to 4
        Get the two numbers from the user
        Perform the operation
        Print the result
        Display the menu again
    If the userChoice is 5
        Exit the program
    If the userChoice is not 1 to 5
        Display "Invalid choice"
        Display the menu again
*/

import "fmt"

func main() {
	for {
		fmt.Println("Menu::")
		fmt.Println("1.\tAdd")
		fmt.Println("2.\tSubtract")
		fmt.Println("3.\tMultiply")
		fmt.Println("4.\tDivide")
		fmt.Println("5.\tExit")
		fmt.Println("Enter the choice :")
		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}

		var a, b, result int
		fmt.Println("Enter the values:")
		fmt.Scanf("%d %d\n", &a, &b)
		switch choice {
		case 1:
			result = a + b
		case 2:
			result = a - b
		case 3:
			result = a * b
		case 4:
			result = a / b
		}

		fmt.Printf("Result:\t%d\n", result)
	}
}
