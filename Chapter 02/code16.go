package main

import (
	"errors"
	"fmt"
)

func main() {
	a, err := exampleFunc()
	b, err2 := exampleFunc2()
	c, err3 := exampleFunc3()
	fmt.Printf("%v - %v\n", a, err)
	fmt.Printf("%v - %v\n", b, err2)
	fmt.Printf("%v - %v\n", c, err3)

}

func exampleFunc() (int, error) {
	return 1, errors.New("This is an example error")
}

func exampleFunc2() (int, error) {
	return 1, fmt.Errorf("This is an example error")
}

func exampleFunc3() (int, error) {
	return 1, fmt.Errorf("This is an example error %v", "example")
}
