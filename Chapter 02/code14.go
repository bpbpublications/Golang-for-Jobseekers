package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start main")
	wg.Add(2)
	go side()
	go side()
	fmt.Println("Return to main")
	wg.Wait()
	fmt.Println("End main")
}

func side() {
	fmt.Println("Start side process")
	time.Sleep(1)
	fmt.Println("End side process")
	wg.Done()
}
