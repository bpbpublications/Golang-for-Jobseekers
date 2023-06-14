package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for j := 0; j < 3; j++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("j: %v\n", j)
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("i: %v\n", i)
	}
}
