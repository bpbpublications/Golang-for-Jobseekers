package main

import (
	"github.com/xxx/fake"
)

func main() {
	a := fake.ExampleStruct{Item1: "a"}
	PrintExample(a)
}

func PrintExample(e fake.ExampleStruct) {
	e.LogLine()
}
