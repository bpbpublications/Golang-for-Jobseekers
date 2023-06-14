package fake

import "fmt"

type ExampleStruct struct {
	Item1 string
	Item2 string
	Item3 string
}

func (e ExampleStruct) LogLine() {
	fmt.Printf("%+v\n", e)
}
