package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataChan := make(chan int)
	evenChan := make(chan int)
	evenSumChan := make(chan int)
	oddChan := make(chan int)
	oddSumChan := make(chan int)

	readWG := new(sync.WaitGroup)
	readWG.Add(2)
	go source("data1.txt", dataChan, readWG)
	go source("data2.txt", dataChan, readWG)

	processWG := new(sync.WaitGroup)
	processWG.Add(4)
	go splitter(dataChan, evenChan, oddChan, processWG)
	go sum(evenChan, evenSumChan, processWG)
	go sum(oddChan, oddSumChan, processWG)
	go merger("result.txt", evenSumChan, oddSumChan, processWG)

	readWG.Wait()
	close(dataChan)
	processWG.Wait()
	fmt.Println("😄 Done")
}

func source(filename string, data chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		text := sc.Text()
		if num, err := strconv.ParseInt(text, 10, 0); err == nil {
			data <- int(num)
		}
	}
}

func splitter(dataChan <-chan int, evenChan, oddChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(evenChan)
	defer close(oddChan)
	for n := range dataChan {
		if n%2 == 0 {
			evenChan <- n
		} else {
			oddChan <- n
		}
	}
}

func sum(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for i := range in {
		sum += i
	}
	out <- sum
}

func merger(filename string, evenResultChan, oddResultChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for i := 0; i < 2; i++ {
		select {
		case result := <-evenResultChan:
			file.WriteString(fmt.Sprintf("Even total : %d\n", result))
		case result := <-oddResultChan:
			file.WriteString(fmt.Sprintf("Odd total : %d\n", result))
		}
	}
}
