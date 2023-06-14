package main

import (
	"fmt"
	"time"
)

var exampleInt = 0

func exampleFunc1() {
	for x := 0; x < 5; x++ {
		exampleInt = exampleInt + x
		time.Sleep(2 * time.Second)
		fmt.Println(exampleInt)
	}
}

func exampleFunc2() {
	for x := 0; x < 5; x++ {
		exampleInt = exampleInt - x
		time.Sleep(1 * time.Second)
		fmt.Println(exampleInt)
	}
}

func main() {
	fmt.Println(exampleInt)
	go exampleFunc1()
	go exampleFunc2()

	fmt.Println("start sleep")
	time.Sleep(15 * time.Second)
	fmt.Println("end sleep")
}
