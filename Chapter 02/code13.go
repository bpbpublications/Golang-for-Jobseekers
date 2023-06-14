package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start main")
	go side()
	fmt.Println("Return to main")
	time.Sleep(5 * time.Second)
	fmt.Println("End main")
}

func side() {
	fmt.Println("Start side process")
	time.Sleep(1 * time.Second)
	fmt.Println("End side process")
}
