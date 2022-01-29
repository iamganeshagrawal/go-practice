package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)

	go print("Hello", ch1, ch2, wg)
	go print("World", ch2, ch1, wg)
	ch1 <- true

	wg.Wait()
	fmt.Println("👋 bye from main")
}

func print(s string, in, out chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
		out <- true
	}
	wg.Done()
}
