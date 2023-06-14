package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	workers = 3
)

func processItem(input chan int, output chan int, wg *sync.WaitGroup) {
	for {
		in := <-input
		fmt.Printf("Working on input: %v\n", in)
		time.Sleep(2 * time.Second)
		output <- in + 1
		wg.Done()
	}
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup
	wg.Add(len(items))
	out := make(chan int, len(items))
	in := make(chan int)

	for i := 0; i < workers; i++ {
		go processItem(in, out, &wg)
	}

	for _, val := range items {
		in <- val
	}

	wg.Wait()

	total := 0
	for j := 0; j < len(items); j++ {
		total = total + <-out
	}
	fmt.Printf("Total sum is : %v\n", total)
}
