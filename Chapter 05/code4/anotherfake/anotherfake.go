package anotherfake

import "fmt"

type ExampleStruct struct{}

func (e ExampleStruct) LogLine() {
	fmt.Printf("This is from anotherfake package :: %+v\n", e)
}
