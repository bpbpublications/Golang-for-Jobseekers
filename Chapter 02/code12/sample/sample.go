package sample

import "fmt"

type unexposedStruct struct {
	Sample string
}

type ExposedStruct struct {
	ExposedSample   string
	unexposedSample string
}

func unexposedFunc() {
	fmt.Println("unexposed")
}

func ExposedFunc() {
	fmt.Println("exposed")
}

func NewUnexposedStruct() unexposedStruct {
	return unexposedStruct{
		Sample: "Sample",
	}
}
