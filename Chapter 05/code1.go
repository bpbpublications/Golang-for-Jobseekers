package main

import "fmt"

func attemptRecover() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func examplePanicFunc() {
	defer attemptRecover()
	fmt.Println("inside examplePanicFunc")
	panic("sudden panic")
}

func main() {
	examplePanicFunc()
	fmt.Println("final")
}
