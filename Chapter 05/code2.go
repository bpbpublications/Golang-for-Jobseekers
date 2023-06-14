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
	internal1()
}

func internal1() {
	fmt.Println("inside internal 1")
	internal2()
}

func internal2() {
	fmt.Println("inside internal 2")
	internal3()
}

func internal3() {
	fmt.Println("inside internal 3")
	a := []string{}
	fmt.Println(a[12])
}

func main() {
	examplePanicFunc()
	fmt.Println("final")
}
