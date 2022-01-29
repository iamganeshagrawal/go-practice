package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = new(sync.WaitGroup)
	var ch = make(chan int)
	var result int
	wg.Add(1)
	go add(2, 6, wg, ch)
	result = <-ch
	wg.Wait()
	fmt.Println(result)
}

func add(a, b int, wg *sync.WaitGroup, ch chan int) {
	res := a + b
	ch <- res
	wg.Done()
}
