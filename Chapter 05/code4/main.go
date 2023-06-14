package main

import (
	// "github.com/xxx/fake"
	"github.com/xxx/anotherfake"
)

func main() {
	a := anotherfake.ExampleStruct{}
	PrintExample(a)
}

type ExampleInterface interface {
	LogLine()
}

func PrintExample(e ExampleInterface) {
	e.LogLine()
}
