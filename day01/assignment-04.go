package main

/*
Assignment:4
    Create increment(), decrement() functions that behaves as follows:

    increment() //=> 1
    increment() //=> 2
    increment() //=> 3
    increment() //=> 4

    decrement() //=> 3
    decrement() //=> 2
    decrement() //=> 1
    decrement() //=> 0
    decrement() //=> -1

    IMPORTANT NOTE: the outcome of increment() and decrement() functions should NOT BE able to be influenced from outside
*/

import "fmt"

type CounterFunc func() int

func main() {
	increment, decrement := getCounter()

	fmt.Println("calling increment(): ", increment())
	fmt.Println("calling increment(): ", increment())
	fmt.Println("calling increment(): ", increment())
	fmt.Println("calling increment(): ", increment())

	fmt.Println("calling decrement(): ", decrement())
	fmt.Println("calling decrement(): ", decrement())
	fmt.Println("calling decrement(): ", decrement())
	fmt.Println("calling decrement(): ", decrement())
	fmt.Println("calling decrement(): ", decrement())

}

func getCounter() (incrementCounter, decrementCounter CounterFunc) {
	count := 0

	incrementCounter = func() int {
		count += 1
		return count
	}
	decrementCounter = func() int {
		count -= 1
		return count
	}

	return incrementCounter, decrementCounter
}
