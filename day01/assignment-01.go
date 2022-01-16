package main

/*
Assignment:1
    Write a program that will print the prime numbers between 3 to 100
*/

import (
	"fmt"
	"math"
)

func main() {
	start := 3
	end := 100

	for num := start; num <= end; num++ {
		isPrime := true
		sqRoot := int(math.Sqrt(float64(num)))

		for i := 2; i <= sqRoot; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("%d, ", num)
		}
	}
}
