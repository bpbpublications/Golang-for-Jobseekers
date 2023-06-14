package main

import (
	"fmt"

	"github.com/example/tester/sample"
)

func main() {
	fmt.Println("Program Begin")

	e := sample.ExposedStruct{
		ExposedSample: "sample",
		// below is a private portion of the struct
		// unexposedSample: "sample",
	}
	fmt.Println(e)
}
